package template

import (
	_ "embed"
	"fmt"
)

//go:embed github-markdown.min.css
var githubMarkdownCSS string

//go:embed chroma.css
var chromaCSS string

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

const copyButtonScript = `
    <script>
        (function() {
            var copyIcon = '<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg>';
            var checkIcon = '<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 6 9 17 4 12"></polyline></svg>';

            document.querySelectorAll('.markdown-body pre').forEach(function(pre) {
                var wrapper = document.createElement('div');
                wrapper.className = 'code-block-wrapper';
                pre.parentNode.insertBefore(wrapper, pre);
                wrapper.appendChild(pre);

                var btn = document.createElement('button');
                btn.className = 'code-copy-btn';
                btn.innerHTML = copyIcon;
                btn.title = 'Copy code';
                wrapper.appendChild(btn);

                btn.addEventListener('click', function() {
                    var code = pre.querySelector('code');
                    var text = code ? code.textContent : pre.textContent;
                    navigator.clipboard.writeText(text).then(function() {
                        btn.innerHTML = checkIcon;
                        btn.classList.add('copied');
                        setTimeout(function() {
                            btn.innerHTML = copyIcon;
                            btn.classList.remove('copied');
                        }, 2000);
                    });
                });
            });
        })();
    </script>`

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
	return fmt.Sprintf(htmlTemplate, title, githubMarkdownCSS, chromaCSS, content, copyButtonScript)
}

// GenerateWithLiveReload creates an HTML document with live reload support.
func GenerateWithLiveReload(title, content string, port int) string {
	script := copyButtonScript + fmt.Sprintf(liveReloadScript, port)
	return fmt.Sprintf(htmlTemplate, title, githubMarkdownCSS, chromaCSS, content, script)
}
