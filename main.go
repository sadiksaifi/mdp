package main

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

//go:embed assets/github-markdown.min.css
var githubMarkdownCSS string

const htmlTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>%s</title>
    <style>
        %s
    </style>
    <style>
        body {
            box-sizing: border-box;
            min-width: 200px;
            max-width: 980px;
            margin: 0 auto;
            padding: 45px;
        }
        @media (prefers-color-scheme: dark) {
            body {
                background-color: #0d1117;
            }
        }
    </style>
</head>
<body>
    <article class="markdown-body">
        %s
    </article>
</body>
</html>`

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <markdown-file.md>\n", os.Args[0])
		os.Exit(1)
	}

	filePath := os.Args[1]

	// Validate .md extension
	if !strings.HasSuffix(strings.ToLower(filePath), ".md") {
		fmt.Fprintf(os.Stderr, "Error: File must have .md extension\n")
		os.Exit(1)
	}

	// Read markdown file
	markdownContent, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Initialize goldmark with GFM extension
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithUnsafe(), // Allow raw HTML in markdown
		),
	)

	// Convert markdown to HTML
	var htmlBuffer strings.Builder
	if err := md.Convert(markdownContent, &htmlBuffer); err != nil {
		fmt.Fprintf(os.Stderr, "Error converting markdown: %v\n", err)
		os.Exit(1)
	}

	// Get filename for title
	filename := filepath.Base(filePath)
	title := strings.TrimSuffix(filename, filepath.Ext(filename))

	// Generate full HTML document
	fullHTML := fmt.Sprintf(htmlTemplate, title, githubMarkdownCSS, htmlBuffer.String())

	// Create output file in /tmp
	outputFileName := fmt.Sprintf("mdpreview-%s.html", strings.ReplaceAll(filename, ".md", ""))
	outputPath := filepath.Join("/tmp", outputFileName)

	if err := os.WriteFile(outputPath, []byte(fullHTML), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing HTML file: %v\n", err)
		os.Exit(1)
	}

	// Open in default browser using macOS 'open' command
	cmd := exec.Command("open", outputPath)
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error opening browser: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Opened %s in browser\n", filePath)
}
