package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"mdp/internal/browser"
	"mdp/internal/converter"
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

	if len(args) != 1 {
		return fmt.Errorf("Usage: mdp <markdown-file.md>\nRun 'mdp --help' for more information")
	}

	filePath := args[0]

	// Validate .md extension
	if !strings.HasSuffix(strings.ToLower(filePath), ".md") {
		return fmt.Errorf("Error: File must have .md extension")
	}

	// Read markdown file
	markdownContent, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("Error reading file: %v", err)
	}

	// Convert markdown to HTML
	conv := converter.New()
	htmlContent, err := conv.Convert(markdownContent)
	if err != nil {
		return fmt.Errorf("Error converting markdown: %v", err)
	}

	// Get filename for title
	filename := filepath.Base(filePath)
	title := strings.TrimSuffix(filename, filepath.Ext(filename))

	// Generate full HTML document
	fullHTML := template.Generate(title, htmlContent)

	// Create output file in /tmp
	outputFileName := fmt.Sprintf("mdpreview-%s.html", strings.ReplaceAll(filename, ".md", ""))
	outputPath := filepath.Join("/tmp", outputFileName)

	if err := os.WriteFile(outputPath, []byte(fullHTML), 0644); err != nil {
		return fmt.Errorf("Error writing HTML file: %v", err)
	}

	// Open in default browser
	if err := browser.Open(outputPath); err != nil {
		return fmt.Errorf("Error opening browser: %v", err)
	}

	fmt.Printf("Opened %s in browser\n", filePath)
	return nil
}

func printUsage() {
	fmt.Println(`mdp - Markdown Previewer

Usage:
  mdp <markdown-file.md>    Preview markdown file in browser
  mdp -h, --help            Show this help message
  mdp -v, --version         Show version

Examples:
  mdp README.md             Preview README.md in browser
  mdp docs/guide.md         Preview guide.md in browser`)
}
