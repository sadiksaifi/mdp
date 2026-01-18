package server

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/websocket"

	"mdp/internal/browser"
	"mdp/internal/converter"
	"mdp/internal/filetree"
	"mdp/internal/linkrewriter"
	"mdp/internal/template"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for local development
	},
}

// Server handles live reload of markdown files.
type Server struct {
	port      int
	files     []string
	baseDir   string
	conv      *converter.Converter
	watcher   *fsnotify.Watcher
	clients   map[*websocket.Conn]bool
	clientsMu sync.RWMutex
	htmlCache string
	cacheMu   sync.RWMutex
}

// New creates a new live reload server.
func New(port int, files []string) (*Server, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, fmt.Errorf("failed to create file watcher: %w", err)
	}

	s := &Server{
		port:    port,
		files:   files,
		baseDir: findCommonBase(files),
		conv:    converter.New(),
		watcher: watcher,
		clients: make(map[*websocket.Conn]bool),
	}

	return s, nil
}

// Start starts the live reload server.
func (s *Server) Start() error {
	// Generate initial HTML
	if err := s.regenerateHTML(); err != nil {
		return err
	}

	// Watch all markdown files
	for _, file := range s.files {
		if err := s.watcher.Add(file); err != nil {
			return fmt.Errorf("failed to watch %s: %w", file, err)
		}
	}

	// Also watch directories for new files
	dirs := make(map[string]bool)
	for _, file := range s.files {
		dir := filepath.Dir(file)
		dirs[dir] = true
	}
	for dir := range dirs {
		if err := s.watcher.Add(dir); err != nil {
			log.Printf("Warning: could not watch directory %s: %v", dir, err)
		}
	}

	// Start file watcher goroutine
	go s.watchFiles()

	// Setup HTTP handlers using a new ServeMux to avoid conflicts
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.handleIndex)
	mux.HandleFunc("/ws", s.handleWebSocket)

	// Try to find an available port
	listener, err := s.findAvailablePort()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("http://localhost:%d", s.port)

	fmt.Printf("Starting live reload server at %s\n", url)
	fmt.Printf("Watching %d file(s) for changes\n", len(s.files))
	fmt.Println("Press Ctrl+C to stop")

	// Open browser automatically
	go func() {
		if err := browser.Open(url); err != nil {
			log.Printf("Warning: could not open browser: %v", err)
		}
	}()

	return http.Serve(listener, mux)
}

// findAvailablePort tries to bind to the configured port, incrementing if occupied.
func (s *Server) findAvailablePort() (net.Listener, error) {
	maxPort := 65535
	startPort := s.port

	for port := startPort; port <= maxPort; port++ {
		addr := fmt.Sprintf(":%d", port)
		listener, err := net.Listen("tcp", addr)
		if err == nil {
			s.port = port // Update the port to the one we actually bound to
			// Regenerate HTML with the correct port for WebSocket connection
			if port != startPort {
				if err := s.regenerateHTML(); err != nil {
					listener.Close()
					return nil, err
				}
			}
			return listener, nil
		}
	}

	return nil, fmt.Errorf("could not find available port starting from %d", startPort)
}

// Stop closes the server and cleans up resources.
func (s *Server) Stop() error {
	s.clientsMu.Lock()
	for client := range s.clients {
		client.Close()
	}
	s.clientsMu.Unlock()

	return s.watcher.Close()
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	s.cacheMu.RLock()
	html := s.htmlCache
	s.cacheMu.RUnlock()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html))
}

func (s *Server) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade failed: %v", err)
		return
	}

	s.clientsMu.Lock()
	s.clients[conn] = true
	s.clientsMu.Unlock()

	// Keep connection alive and handle disconnect
	for {
		if _, _, err := conn.ReadMessage(); err != nil {
			s.clientsMu.Lock()
			delete(s.clients, conn)
			s.clientsMu.Unlock()
			conn.Close()
			break
		}
	}
}

func (s *Server) watchFiles() {
	for {
		select {
		case event, ok := <-s.watcher.Events:
			if !ok {
				return
			}

			// Only react to write and create events for .md files
			if event.Op&(fsnotify.Write|fsnotify.Create) != 0 {
				if strings.HasSuffix(strings.ToLower(event.Name), ".md") {
					log.Printf("File changed: %s", event.Name)
					if err := s.regenerateHTML(); err != nil {
						log.Printf("Error regenerating HTML: %v", err)
						continue
					}
					s.notifyClients()
				}
			}

		case err, ok := <-s.watcher.Errors:
			if !ok {
				return
			}
			log.Printf("Watcher error: %v", err)
		}
	}
}

func (s *Server) regenerateHTML() error {
	if len(s.files) == 1 {
		return s.regenerateSingleFile()
	}
	return s.regenerateMultiFile()
}

func (s *Server) regenerateSingleFile() error {
	filePath := s.files[0]

	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	htmlContent, err := s.conv.Convert(content)
	if err != nil {
		return fmt.Errorf("error converting markdown: %w", err)
	}

	filename := filepath.Base(filePath)
	title := strings.TrimSuffix(filename, filepath.Ext(filename))

	html := template.GenerateWithLiveReload(title, htmlContent, s.port)

	s.cacheMu.Lock()
	s.htmlCache = html
	s.cacheMu.Unlock()

	return nil
}

func (s *Server) regenerateMultiFile() error {
	var entries []filetree.FileEntry

	for _, path := range s.files {
		content, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("error reading %s: %w", path, err)
		}

		htmlContent, err := s.conv.Convert(content)
		if err != nil {
			return fmt.Errorf("error converting %s: %w", path, err)
		}

		relPath := strings.TrimPrefix(path, s.baseDir)
		relPath = strings.TrimPrefix(relPath, string(filepath.Separator))

		entries = append(entries, filetree.FileEntry{
			ID:      sanitizeID(relPath),
			Path:    path,
			Name:    strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)),
			RelPath: relPath,
			Content: htmlContent,
		})
	}

	// Rewrite relative .md links to fragment identifiers
	rewriter := linkrewriter.New(entries)
	for i := range entries {
		entries[i].Content = rewriter.RewriteLinks(entries[i].Content, entries[i].RelPath)
	}

	tree := filetree.BuildTree(entries)
	title := s.generateTitle()
	html := template.GenerateMultiWithLiveReload(title, tree, entries, s.port)

	s.cacheMu.Lock()
	s.htmlCache = html
	s.cacheMu.Unlock()

	return nil
}

func (s *Server) notifyClients() {
	s.clientsMu.RLock()
	defer s.clientsMu.RUnlock()

	message := []byte("reload")
	for client := range s.clients {
		if err := client.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Printf("Error sending to client: %v", err)
		}
	}
}

func (s *Server) generateTitle() string {
	if s.baseDir != "" {
		return filepath.Base(s.baseDir) + " - Markdown Preview"
	}
	return fmt.Sprintf("%d Files - Markdown Preview", len(s.files))
}

// sanitizeID converts a path to a valid HTML id attribute.
func sanitizeID(path string) string {
	id := strings.ReplaceAll(path, "/", "-")
	id = strings.ReplaceAll(id, "\\", "-")
	id = strings.ReplaceAll(id, ".", "-")
	id = strings.ReplaceAll(id, " ", "-")
	id = strings.ToLower(id)
	id = strings.TrimLeft(id, "-")
	return id
}

// findCommonBase finds the common directory prefix of all paths.
func findCommonBase(paths []string) string {
	if len(paths) == 0 {
		return ""
	}

	if len(paths) == 1 {
		return filepath.Dir(paths[0])
	}

	first := filepath.Dir(paths[0])
	parts := strings.Split(first, string(filepath.Separator))

	for _, path := range paths[1:] {
		dir := filepath.Dir(path)
		dirParts := strings.Split(dir, string(filepath.Separator))

		minLen := len(parts)
		if len(dirParts) < minLen {
			minLen = len(dirParts)
		}

		commonLen := 0
		for i := 0; i < minLen; i++ {
			if parts[i] == dirParts[i] {
				commonLen = i + 1
			} else {
				break
			}
		}
		parts = parts[:commonLen]
	}

	return strings.Join(parts, string(filepath.Separator))
}
