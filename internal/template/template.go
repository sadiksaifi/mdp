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
    </article>%s
</body>
</html>`

const liveReloadScript = `
    <script>
        (function() {
            var ws = new WebSocket('ws://localhost:%d/ws');
            ws.onmessage = function(event) {
                if (event.data === 'reload') {
                    location.reload();
                }
            };
            ws.onclose = function() {
                console.log('Live reload disconnected. Attempting to reconnect...');
                setTimeout(function() {
                    location.reload();
                }, 1000);
            };
        })();
    </script>`

// Generate creates a complete HTML document with the given title and content.
func Generate(title, content string) string {
	return fmt.Sprintf(htmlTemplate, title, githubMarkdownCSS, content, "")
}

// GenerateWithLiveReload creates an HTML document with live reload support.
func GenerateWithLiveReload(title, content string, port int) string {
	script := fmt.Sprintf(liveReloadScript, port)
	return fmt.Sprintf(htmlTemplate, title, githubMarkdownCSS, content, script)
}
