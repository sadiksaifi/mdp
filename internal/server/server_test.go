package server

import (
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func TestNew(t *testing.T) {
	// Create a temporary markdown file
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.md")
	if err := os.WriteFile(tmpFile, []byte("# Test"), 0644); err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	tests := []struct {
		name    string
		port    int
		files   []string
		wantErr bool
	}{
		{
			name:    "valid single file",
			port:    8080,
			files:   []string{tmpFile},
			wantErr: false,
		},
		{
			name:    "valid multiple files",
			port:    3000,
			files:   []string{tmpFile, tmpFile},
			wantErr: false,
		},
		{
			name:    "empty files",
			port:    8080,
			files:   []string{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv, err := New(tt.port, tt.files)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if srv == nil {
					t.Error("New() returned nil server")
					return
				}
				if srv.port != tt.port {
					t.Errorf("New() port = %v, want %v", srv.port, tt.port)
				}
				if len(srv.files) != len(tt.files) {
					t.Errorf("New() files count = %v, want %v", len(srv.files), len(tt.files))
				}
				srv.Stop()
			}
		})
	}
}

func TestSanitizeID(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "simple filename",
			input: "readme.md",
			want:  "readme-md",
		},
		{
			name:  "path with slashes",
			input: "docs/guide/intro.md",
			want:  "docs-guide-intro-md",
		},
		{
			name:  "path with backslashes",
			input: "docs\\guide\\intro.md",
			want:  "docs-guide-intro-md",
		},
		{
			name:  "path with spaces",
			input: "my docs/my file.md",
			want:  "my-docs-my-file-md",
		},
		{
			name:  "uppercase letters",
			input: "README.MD",
			want:  "readme-md",
		},
		{
			name:  "leading slashes",
			input: "/docs/readme.md",
			want:  "docs-readme-md",
		},
		{
			name:  "multiple leading slashes",
			input: "///docs/readme.md",
			want:  "docs-readme-md",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sanitizeID(tt.input)
			if got != tt.want {
				t.Errorf("sanitizeID(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestFindCommonBase(t *testing.T) {
	tests := []struct {
		name  string
		paths []string
		want  string
	}{
		{
			name:  "empty paths",
			paths: []string{},
			want:  "",
		},
		{
			name:  "single path",
			paths: []string{"/home/user/docs/readme.md"},
			want:  "/home/user/docs",
		},
		{
			name:  "same directory",
			paths: []string{"/home/user/docs/a.md", "/home/user/docs/b.md"},
			want:  "/home/user/docs",
		},
		{
			name:  "different subdirectories",
			paths: []string{"/home/user/docs/guide/a.md", "/home/user/docs/api/b.md"},
			want:  "/home/user/docs",
		},
		{
			name:  "completely different paths",
			paths: []string{"/home/user/a.md", "/var/log/b.md"},
			want:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findCommonBase(tt.paths)
			if got != tt.want {
				t.Errorf("findCommonBase(%v) = %q, want %q", tt.paths, got, tt.want)
			}
		})
	}
}

func TestServer_generateTitle(t *testing.T) {
	tests := []struct {
		name    string
		baseDir string
		files   []string
		want    string
	}{
		{
			name:    "with base directory",
			baseDir: "/home/user/docs",
			files:   []string{"a.md", "b.md"},
			want:    "docs - Markdown Preview",
		},
		{
			name:    "without base directory",
			baseDir: "",
			files:   []string{"a.md", "b.md", "c.md"},
			want:    "3 Files - Markdown Preview",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &Server{
				baseDir: tt.baseDir,
				files:   tt.files,
			}
			got := srv.generateTitle()
			if got != tt.want {
				t.Errorf("generateTitle() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestServer_handleIndex(t *testing.T) {
	// Create a temporary markdown file
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.md")
	if err := os.WriteFile(tmpFile, []byte("# Hello World\n\nThis is a test."), 0644); err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	srv, err := New(8080, []string{tmpFile})
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer srv.Stop()

	// Regenerate HTML to populate cache
	if err := srv.regenerateHTML(); err != nil {
		t.Fatalf("Failed to regenerate HTML: %v", err)
	}

	// Create a test request
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	// Call the handler
	srv.handleIndex(rec, req)

	// Check response
	if rec.Code != http.StatusOK {
		t.Errorf("handleIndex() status = %d, want %d", rec.Code, http.StatusOK)
	}

	contentType := rec.Header().Get("Content-Type")
	if contentType != "text/html; charset=utf-8" {
		t.Errorf("handleIndex() Content-Type = %q, want %q", contentType, "text/html; charset=utf-8")
	}

	body := rec.Body.String()
	if !strings.Contains(body, "Hello World") {
		t.Error("handleIndex() response body should contain 'Hello World'")
	}
	if !strings.Contains(body, "WebSocket") {
		t.Error("handleIndex() response body should contain WebSocket script for live reload")
	}
}

func TestServer_regenerateSingleFile(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.md")
	content := "# Test Header\n\nSome **bold** text."
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	srv, err := New(8080, []string{tmpFile})
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer srv.Stop()

	if err := srv.regenerateSingleFile(); err != nil {
		t.Errorf("regenerateSingleFile() error = %v", err)
	}

	srv.cacheMu.RLock()
	html := srv.htmlCache
	srv.cacheMu.RUnlock()

	if html == "" {
		t.Error("regenerateSingleFile() should populate htmlCache")
	}
	if !strings.Contains(html, "Test Header") {
		t.Error("regenerateSingleFile() HTML should contain header text")
	}
	if !strings.Contains(html, "<strong>bold</strong>") {
		t.Error("regenerateSingleFile() HTML should contain converted bold text")
	}
}

func TestServer_regenerateMultiFile(t *testing.T) {
	tmpDir := t.TempDir()

	// Create multiple markdown files
	file1 := filepath.Join(tmpDir, "readme.md")
	file2 := filepath.Join(tmpDir, "guide.md")

	if err := os.WriteFile(file1, []byte("# README\n\nIntro text."), 0644); err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	if err := os.WriteFile(file2, []byte("# Guide\n\nGuide content."), 0644); err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	srv, err := New(8080, []string{file1, file2})
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer srv.Stop()

	if err := srv.regenerateMultiFile(); err != nil {
		t.Errorf("regenerateMultiFile() error = %v", err)
	}

	srv.cacheMu.RLock()
	html := srv.htmlCache
	srv.cacheMu.RUnlock()

	if html == "" {
		t.Error("regenerateMultiFile() should populate htmlCache")
	}
	if !strings.Contains(html, "README") {
		t.Error("regenerateMultiFile() HTML should contain README content")
	}
	if !strings.Contains(html, "Guide") {
		t.Error("regenerateMultiFile() HTML should contain Guide content")
	}
	if !strings.Contains(html, "sidebar") {
		t.Error("regenerateMultiFile() HTML should contain sidebar")
	}
}

func TestServer_regenerateHTML_NonExistentFile(t *testing.T) {
	srv, err := New(8080, []string{"/nonexistent/file.md"})
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer srv.Stop()

	err = srv.regenerateHTML()
	if err == nil {
		t.Error("regenerateHTML() should return error for non-existent file")
	}
}

func TestServer_WebSocket(t *testing.T) {
	// Create a temporary markdown file
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.md")
	if err := os.WriteFile(tmpFile, []byte("# Test"), 0644); err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	srv, err := New(8080, []string{tmpFile})
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer srv.Stop()

	if err := srv.regenerateHTML(); err != nil {
		t.Fatalf("Failed to regenerate HTML: %v", err)
	}

	// Create test server
	testServer := httptest.NewServer(http.HandlerFunc(srv.handleWebSocket))
	defer testServer.Close()

	// Connect WebSocket
	wsURL := "ws" + strings.TrimPrefix(testServer.URL, "http") + "/ws"
	ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("Failed to connect WebSocket: %v", err)
	}
	defer ws.Close()

	// Wait a bit for the connection to be registered
	time.Sleep(50 * time.Millisecond)

	// Check that client is registered
	srv.clientsMu.RLock()
	clientCount := len(srv.clients)
	srv.clientsMu.RUnlock()

	if clientCount != 1 {
		t.Errorf("Expected 1 client, got %d", clientCount)
	}

	// Test notifyClients
	done := make(chan bool)
	go func() {
		_, message, err := ws.ReadMessage()
		if err != nil {
			t.Errorf("Failed to read message: %v", err)
			return
		}
		if string(message) != "reload" {
			t.Errorf("Expected 'reload' message, got %q", string(message))
		}
		done <- true
	}()

	srv.notifyClients()

	select {
	case <-done:
		// Success
	case <-time.After(time.Second):
		t.Error("Timeout waiting for reload message")
	}
}

func TestServer_Stop(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.md")
	if err := os.WriteFile(tmpFile, []byte("# Test"), 0644); err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	srv, err := New(8080, []string{tmpFile})
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	// Stop should not error
	if err := srv.Stop(); err != nil {
		t.Errorf("Stop() error = %v", err)
	}
}

func TestServer_findAvailablePort(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.md")
	if err := os.WriteFile(tmpFile, []byte("# Test"), 0644); err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	t.Run("finds available port", func(t *testing.T) {
		srv, err := New(0, []string{tmpFile}) // Port 0 lets OS assign
		if err != nil {
			t.Fatalf("Failed to create server: %v", err)
		}
		defer srv.Stop()

		// Use a high random port that's likely available
		srv.port = 19876

		if err := srv.regenerateHTML(); err != nil {
			t.Fatalf("Failed to regenerate HTML: %v", err)
		}

		listener, err := srv.findAvailablePort()
		if err != nil {
			t.Errorf("findAvailablePort() error = %v", err)
			return
		}
		defer listener.Close()

		if srv.port < 19876 {
			t.Errorf("findAvailablePort() port = %d, expected >= 19876", srv.port)
		}
	})

	t.Run("tries next port when occupied", func(t *testing.T) {
		srv, err := New(0, []string{tmpFile})
		if err != nil {
			t.Fatalf("Failed to create server: %v", err)
		}
		defer srv.Stop()

		// Occupy a port
		occupiedPort := 19877
		occupiedListener, err := net.Listen("tcp", ":19877")
		if err != nil {
			t.Fatalf("Failed to occupy port: %v", err)
		}
		defer occupiedListener.Close()

		// Try to find port starting from occupied one
		srv.port = occupiedPort

		if err := srv.regenerateHTML(); err != nil {
			t.Fatalf("Failed to regenerate HTML: %v", err)
		}

		listener, err := srv.findAvailablePort()
		if err != nil {
			t.Errorf("findAvailablePort() error = %v", err)
			return
		}
		defer listener.Close()

		// Should have found a port greater than the occupied one
		if srv.port <= occupiedPort {
			t.Errorf("findAvailablePort() should find port > %d, got %d", occupiedPort, srv.port)
		}
	})
}
