package template

import (
	_ "embed"
	"fmt"
)

//go:embed github-markdown.min.css
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

// Generate creates a complete HTML document with the given title and content.
func Generate(title, content string) string {
	return fmt.Sprintf(htmlTemplate, title, githubMarkdownCSS, content)
}
