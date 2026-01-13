package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	gitignore "github.com/sabhiram/go-gitignore"

	"mdp/internal/browser"
	"mdp/internal/converter"
	"mdp/internal/filetree"
	"mdp/internal/server"
	"mdp/internal/template"
)

var version = "dev"

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	// Handle help and version before flag parsing
	if len(args) == 1 {
		switch args[0] {
		case "-h", "--help":
			printUsage()
			return nil
		case "-v", "--version":
			fmt.Printf("mdp %s\n", version)
			return nil
		}
	}

	// Setup flags
	fs := flag.NewFlagSet("mdp", flag.ContinueOnError)
	fs.Usage = func() {} // Suppress default usage

	serveFlag := fs.Bool("serve", false, "Start live reload server")
	portFlag := fs.Int("port", 8080, "Port for live reload server (only with --serve)")
	outputFlag := fs.String("output", "", "Write HTML to file instead of opening browser")
	fs.StringVar(outputFlag, "O", "", "Write HTML to file instead of opening browser (shorthand)")

	// Parse flags
	if err := fs.Parse(args); err != nil {
		if err == flag.ErrHelp {
			printUsage()
			return nil
		}
		return err
	}

	// Get remaining arguments (file paths)
	fileArgs := fs.Args()

	if len(fileArgs) == 0 {
		return fmt.Errorf("Usage: mdp <markdown-file.md>\nRun 'mdp --help' for more information")
	}

	files, err := resolveFiles(fileArgs)
	if err != nil {
		return err
	}

	if len(files) == 0 {
		return fmt.Errorf("No markdown files found")
	}

	// Validate flag combinations
	if *outputFlag != "" && *serveFlag {
		return fmt.Errorf("cannot use --output with --serve")
	}

	// Serve mode with live reload
	if *serveFlag {
		return runServe(files, *portFlag)
	}

	// Static mode (original behavior)
	if len(files) == 1 {
		return runSingleFile(files[0], *outputFlag)
	}

	return runMultiFile(files, *outputFlag)
}

// runServe starts the live reload server.
func runServe(files []string, port int) error {
	srv, err := server.New(port, files)
	if err != nil {
		return fmt.Errorf("failed to create server: %w", err)
	}

	return srv.Start()
}

// resolveFiles expands directories and validates all paths.
func resolveFiles(args []string) ([]string, error) {
	var files []string
	for _, arg := range args {
		info, err := os.Stat(arg)
		if err != nil {
			return nil, fmt.Errorf("Error accessing %s: %v", arg, err)
		}

		if info.IsDir() {
			discovered, err := discoverMarkdownFiles(arg)
			if err != nil {
				return nil, err
			}
			files = append(files, discovered...)
		} else {
			if !strings.HasSuffix(strings.ToLower(arg), ".md") {
				return nil, fmt.Errorf("Error: File must have .md extension: %s", arg)
			}
			files = append(files, arg)
		}
	}
	return files, nil
}

// discoverMarkdownFiles walks a directory recursively to find all .md files,
// respecting .gitignore files at each level of the directory tree.
func discoverMarkdownFiles(dir string) ([]string, error) {
	var files []string

	// Map of directory path to its gitignore matcher
	ignoreMatchers := make(map[string]*gitignore.GitIgnore)

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip hidden directories (like .git)
		if d.IsDir() && strings.HasPrefix(d.Name(), ".") && path != dir {
			return filepath.SkipDir
		}

		// Check if this path is ignored by any applicable .gitignore
		if isIgnored(path, dir, ignoreMatchers) {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		if d.IsDir() {
			// Look for .gitignore in this directory
			gitignorePath := filepath.Join(path, ".gitignore")
			if _, statErr := os.Stat(gitignorePath); statErr == nil {
				if matcher, compileErr := gitignore.CompileIgnoreFile(gitignorePath); compileErr == nil {
					ignoreMatchers[path] = matcher
				}
			}
			return nil
		}

		// Check if it's a markdown file
		if strings.HasSuffix(strings.ToLower(d.Name()), ".md") {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("Error walking directory %s: %v", dir, err)
	}
	sort.Strings(files)
	return files, nil
}

// isIgnored checks if a path should be ignored based on all applicable .gitignore files.
func isIgnored(path, baseDir string, matchers map[string]*gitignore.GitIgnore) bool {
	for ignoreDir, matcher := range matchers {
		// Check if this gitignore applies (path is under ignoreDir)
		if !strings.HasPrefix(path, ignoreDir+string(filepath.Separator)) && path != ignoreDir {
			continue
		}

		// Get path relative to the gitignore's directory
		relPath, err := filepath.Rel(ignoreDir, path)
		if err != nil {
			continue
		}

		if matcher.MatchesPath(relPath) {
			return true
		}
	}
	return false
}

// runSingleFile handles single file preview (original behavior).
// If outputPath is provided, writes to that path instead of /tmp and skips browser.
func runSingleFile(filePath string, outputPath string) error {
	markdownContent, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("Error reading file: %v", err)
	}

	conv := converter.New()
	htmlContent, err := conv.Convert(markdownContent)
	if err != nil {
		return fmt.Errorf("Error converting markdown: %v", err)
	}

	filename := filepath.Base(filePath)
	title := strings.TrimSuffix(filename, filepath.Ext(filename))

	fullHTML := template.Generate(title, htmlContent)

	// Determine output path
	openBrowser := false
	if outputPath == "" {
		outputFileName := fmt.Sprintf("mdpreview-%s.html", strings.ReplaceAll(filename, ".md", ""))
		outputPath = filepath.Join("/tmp", outputFileName)
		openBrowser = true
	}

	if err := os.WriteFile(outputPath, []byte(fullHTML), 0644); err != nil {
		return fmt.Errorf("Error writing HTML file: %v", err)
	}

	if openBrowser {
		if err := browser.Open(outputPath); err != nil {
			return fmt.Errorf("Error opening browser: %v", err)
		}
		fmt.Printf("Opened %s in browser\n", filePath)
	} else {
		fmt.Printf("Wrote %s\n", outputPath)
	}
	return nil
}

// runMultiFile handles multiple files preview with sidebar.
// If outputPath is provided, writes to that path instead of /tmp and skips browser.
func runMultiFile(filePaths []string, outputPath string) error {
	conv := converter.New()

	baseDir := findCommonBase(filePaths)

	var entries []filetree.FileEntry
	for _, path := range filePaths {
		content, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("Error reading %s: %v", path, err)
		}

		htmlContent, err := conv.Convert(content)
		if err != nil {
			return fmt.Errorf("Error converting %s: %v", path, err)
		}

		relPath := strings.TrimPrefix(path, baseDir)
		relPath = strings.TrimPrefix(relPath, string(filepath.Separator))

		entries = append(entries, filetree.FileEntry{
			ID:      sanitizeID(relPath),
			Path:    path,
			Name:    strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)),
			RelPath: relPath,
			Content: htmlContent,
		})
	}

	tree := filetree.BuildTree(entries)

	title := generateTitle(baseDir, filePaths)
	fullHTML := template.GenerateMulti(title, tree, entries)

	// Determine output path
	openBrowser := false
	if outputPath == "" {
		outputPath = filepath.Join("/tmp", "mdpreview-multi.html")
		openBrowser = true
	}

	if err := os.WriteFile(outputPath, []byte(fullHTML), 0644); err != nil {
		return fmt.Errorf("Error writing HTML file: %v", err)
	}

	if openBrowser {
		if err := browser.Open(outputPath); err != nil {
			return fmt.Errorf("Error opening browser: %v", err)
		}
		fmt.Printf("Opened %d files in browser\n", len(filePaths))
	} else {
		fmt.Printf("Wrote %s\n", outputPath)
	}
	return nil
}

// sanitizeID converts a path to a valid HTML id attribute.
func sanitizeID(path string) string {
	id := strings.ReplaceAll(path, "/", "-")
	id = strings.ReplaceAll(id, "\\", "-")
	id = strings.ReplaceAll(id, ".", "-")
	id = strings.ReplaceAll(id, " ", "-")
	id = strings.ToLower(id)
	// Remove leading hyphens
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

	// Get directory of first file
	first := filepath.Dir(paths[0])
	parts := strings.Split(first, string(filepath.Separator))

	for _, path := range paths[1:] {
		dir := filepath.Dir(path)
		dirParts := strings.Split(dir, string(filepath.Separator))

		// Find common prefix
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

// generateTitle creates a title for the multi-file preview.
func generateTitle(baseDir string, paths []string) string {
	if baseDir != "" {
		return filepath.Base(baseDir) + " - Markdown Preview"
	}
	return fmt.Sprintf("%d Files - Markdown Preview", len(paths))
}

func printUsage() {
	fmt.Println(`mdp - Markdown Previewer

Usage:
  mdp <file.md>                Preview single markdown file
  mdp <file1.md> <file2.md>    Preview multiple files with sidebar
  mdp <directory>              Preview all .md files in directory
  mdp -h, --help               Show this help message
  mdp -v, --version            Show version

Options:
  -O, --output <file>          Write HTML to file instead of opening browser
  --serve                      Start live reload server instead of opening browser
  --port <port>                Port for live reload server (default: 8080)

Examples:
  mdp README.md                Preview single file
  mdp docs/                    Preview all markdown in docs/
  mdp README.md CHANGELOG.md   Preview multiple files with sidebar
  mdp -O site.html docs/       Convert docs to single HTML file
  mdp --serve README.md        Start live reload server for single file
  mdp --serve --port 3000 .    Live reload all markdown in current directory`)
}
