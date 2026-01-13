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
        .github-link {
            position: fixed;
            bottom: 16px;
            right: 16px;
            display: flex;
            align-items: center;
            gap: 8px;
            padding: 8px 12px;
            background: #21262d;
            border: 1px solid #30363d;
            border-radius: 8px;
            color: #8b949e;
            text-decoration: none;
            font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", "Noto Sans", Helvetica, Arial, sans-serif;
            font-size: 12px;
            transition: all 0.2s ease;
            z-index: 100;
        }
        .github-link:hover {
            color: #e6edf3;
            background: #30363d;
        }
        .github-link svg {
            flex-shrink: 0;
        }
        @media (prefers-color-scheme: light) {
            .github-link {
                background: #f6f8fa;
                border-color: #d0d7de;
                color: #656d76;
            }
            .github-link:hover {
                color: #1f2328;
                background: #eaeef2;
            }
        }
    </style>
</head>
<body>
    <article class="markdown-body">
        %s
    </article>
    <a href="https://github.com/sadiksaifi/mdp" class="github-link" target="_blank" rel="noopener noreferrer" title="Star on GitHub">
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="currentColor"><path d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"/></svg>
        <span>Made by Sadik Saifi</span>
    </a>%s
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
