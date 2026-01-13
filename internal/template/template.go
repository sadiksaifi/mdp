package template

import (
	_ "embed"
	"fmt"
)

//go:embed github-markdown.min.css
var githubMarkdownCSS string

//go:embed chroma.css
var chromaCSS string

const commentsCSS = `
:root {
    --panel-section-height: 56px;
}

/* Comment Button */
.comment-btn {
    position: absolute;
    display: none;
    align-items: center;
    gap: 6px;
    padding: 8px 12px;
    background: #0969da;
    color: #ffffff;
    border: none;
    border-radius: 6px;
    font-size: 13px;
    font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", "Noto Sans", Helvetica, Arial, sans-serif;
    font-weight: 500;
    cursor: pointer;
    z-index: 500;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
    transition: opacity 0.15s ease, transform 0.15s ease, background 0.15s ease;
}

.comment-btn.visible {
    display: flex;
}

.comment-btn:hover {
    filter: brightness(1.1);
}

.comment-btn:active {
    transform: scale(0.95);
}

.comment-btn svg {
    width: 14px;
    height: 14px;
}

.comment-btn kbd {
    padding: 2px 5px;
    background: rgba(255, 255, 255, 0.2);
    border-radius: 3px;
    font-family: inherit;
    font-size: 11px;
    font-weight: 600;
}

/* Comment Highlights */
.comment-highlight {
    background-color: rgba(255, 212, 59, 0.4);
    cursor: pointer;
    border-radius: 2px;
    padding: 1px 0;
    transition: background-color 0.15s ease;
}

.comment-highlight:hover,
.comment-highlight.active {
    background-color: rgba(255, 212, 59, 0.7);
}

/* Comments Panel */
.comments-panel {
    position: fixed;
    right: 0;
    top: 0;
    bottom: 0;
    width: 320px;
    background: #f6f8fa;
    border-left: 1px solid #d1d9e0;
    display: flex;
    flex-direction: column;
    z-index: 200;
    transform: translateX(100%);
    transition: transform 0.3s ease;
    font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", "Noto Sans", Helvetica, Arial, sans-serif;
}

.comments-panel.open {
    transform: translateX(0);
}

.comments-panel-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    height: var(--panel-section-height);
    padding: 0 16px;
    border-bottom: 1px solid #d1d9e0;
    flex-shrink: 0;
}

.comments-panel-header h2 {
    margin: 0;
    font-size: 14px;
    font-weight: 600;
    color: #59636e;
    text-transform: uppercase;
    letter-spacing: 0.5px;
}

.comments-panel-close-btn {
    display: flex;
    align-items: center;
    justify-content: center;
    background: none;
    border: none;
    cursor: pointer;
    color: #59636e;
    padding: 4px;
    border-radius: 4px;
    transition: all 0.2s ease;
}

.comments-panel-close-btn:hover {
    color: #1f2328;
    background: #e6e8eb;
}

/* Comments List */
.comments-list {
    flex: 1;
    overflow-y: auto;
    padding: 0;
}

.comments-list::-webkit-scrollbar {
    width: 8px;
}

.comments-list::-webkit-scrollbar-track {
    background: transparent;
}

.comments-list::-webkit-scrollbar-thumb {
    background-color: #d1d9e0;
    border-radius: 4px;
}

/* Comment Entry */
.comment-entry {
    padding: 12px 16px;
    border-bottom: 1px solid #d1d9e0;
    cursor: pointer;
    transition: background 0.15s ease;
    position: relative;
}

.comment-entry:hover {
    background: #e6e8eb;
}

.comment-entry.active {
    background: #ddf4ff;
}

.comment-quote {
    margin: 0 0 8px 0;
    padding: 8px 12px;
    background: #e6e8eb;
    border-left: 3px solid #0969da;
    border-radius: 0 4px 4px 0;
    font-size: 13px;
    color: #59636e;
    font-style: italic;
    word-break: break-word;
}

.comment-text {
    font-size: 14px;
    color: #1f2328;
    line-height: 1.5;
    word-break: break-word;
}

.comment-actions {
    position: absolute;
    top: 8px;
    right: 8px;
    display: none;
    align-items: center;
    gap: 4px;
}

.comment-entry:hover .comment-actions {
    display: flex;
}

.comment-copy-btn,
.comment-edit-btn,
.comment-delete-btn {
    display: flex;
    align-items: center;
    justify-content: center;
    background: none;
    border: none;
    cursor: pointer;
    color: #59636e;
    padding: 4px;
    border-radius: 4px;
    transition: all 0.15s ease;
}

.comment-copy-btn:hover,
.comment-edit-btn:hover {
    color: #0969da;
    background: rgba(9, 105, 218, 0.1);
}

.comment-copy-btn.copied {
    color: #1a7f37;
}

.comment-delete-btn:hover {
    color: #cf222e;
    background: rgba(207, 34, 46, 0.1);
}

.comment-copy-btn svg,
.comment-edit-btn svg,
.comment-delete-btn svg {
    width: 14px;
    height: 14px;
}

/* Comment Input Form */
.comment-input-form {
    padding: 16px;
    border-bottom: 1px solid #d1d9e0;
    background: #f6f8fa;
}

.comment-input-quote {
    margin: 0 0 12px 0;
    padding: 8px 12px;
    background: #e6e8eb;
    border-left: 3px solid #0969da;
    border-radius: 0 4px 4px 0;
    font-size: 13px;
    color: #59636e;
    font-style: italic;
    word-break: break-word;
}

.comment-input-textarea {
    width: 100%;
    padding: 10px 12px;
    border: 1px solid #d1d9e0;
    border-radius: 6px;
    background: #ffffff;
    color: #1f2328;
    font-family: inherit;
    font-size: 14px;
    resize: vertical;
    min-height: 80px;
    box-sizing: border-box;
}

.comment-input-textarea:focus {
    outline: none;
    border-color: #0969da;
    box-shadow: 0 0 0 3px rgba(9, 105, 218, 0.1);
}

.comment-input-actions {
    display: flex;
    justify-content: flex-end;
    gap: 8px;
    margin-top: 12px;
}

.comment-cancel-btn,
.comment-save-btn {
    padding: 6px 12px;
    border-radius: 6px;
    font-size: 13px;
    font-weight: 500;
    cursor: pointer;
    transition: all 0.15s ease;
    font-family: inherit;
}

.comment-cancel-btn {
    background: transparent;
    border: 1px solid #d1d9e0;
    color: #1f2328;
}

.comment-cancel-btn:hover {
    background: #e6e8eb;
}

.comment-save-btn {
    display: inline-flex;
    align-items: center;
    gap: 6px;
    background: #0969da;
    border: none;
    color: #ffffff;
}

.comment-save-btn:hover {
    filter: brightness(1.1);
}

/* Open Comments Button */
.open-comments-btn {
    position: fixed;
    top: 16px;
    right: 16px;
    display: flex;
    align-items: center;
    justify-content: center;
    width: 36px;
    height: 36px;
    padding: 0;
    background: #f6f8fa;
    border: 1px solid #d1d9e0;
    border-radius: 8px;
    color: #59636e;
    cursor: pointer;
    z-index: 100;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
    transition: all 0.2s ease;
}

.open-comments-btn:hover {
    color: #1f2328;
    background: #e6e8eb;
}

.open-comments-btn svg {
    width: 18px;
    height: 18px;
}

.comment-count {
    position: absolute;
    top: -6px;
    right: -6px;
    min-width: 18px;
    height: 18px;
    padding: 0 5px;
    background: #0969da;
    color: #ffffff;
    border-radius: 9px;
    font-size: 11px;
    font-weight: 600;
    display: none;
    align-items: center;
    justify-content: center;
}

.comment-count.visible {
    display: flex;
}

/* Comments Panel Footer */
.comments-panel-footer {
    display: flex;
    align-items: center;
    height: var(--panel-section-height);
    padding: 0 16px;
    border-top: 1px solid #d1d9e0;
    flex-shrink: 0;
}

/* Copy Comments Button (in footer) */
.copy-comments-btn {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 8px;
    width: 100%;
    padding: 10px 16px;
    background: #f6f8fa;
    border: 1px solid #d1d9e0;
    border-radius: 8px;
    color: #59636e;
    font-size: 13px;
    font-weight: 500;
    font-family: inherit;
    cursor: pointer;
    transition: all 0.2s ease;
}

.copy-comments-btn:hover {
    color: #1f2328;
    background: #e6e8eb;
}

.copy-comments-btn.copied {
    color: #1a7f37;
}

.copy-comments-btn svg {
    width: 16px;
    height: 16px;
}

/* Empty state for comments */
.comments-empty {
    padding: 40px 16px;
    text-align: center;
    color: #59636e;
    font-size: 14px;
}

.comments-empty svg {
    width: 48px;
    height: 48px;
    margin-bottom: 12px;
    opacity: 0.5;
}

.comments-empty p {
    margin: 0 0 8px 0;
}

.comments-empty kbd {
    display: inline-block;
    padding: 2px 6px;
    background: #e6e8eb;
    border-radius: 4px;
    font-family: inherit;
    font-size: 12px;
    font-weight: 500;
}

/* Keyboard Shortcuts Help Modal */
.shortcuts-modal-overlay {
    position: fixed;
    inset: 0;
    background: rgba(0, 0, 0, 0.5);
    z-index: 1000;
    opacity: 0;
    visibility: hidden;
    transition: opacity 0.15s ease, visibility 0.15s ease;
}

.shortcuts-modal-overlay.active {
    opacity: 1;
    visibility: visible;
}

.shortcuts-modal {
    position: fixed;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%) scale(0.95);
    width: 90%;
    max-width: 400px;
    max-height: 80vh;
    background: #ffffff;
    border-radius: 12px;
    box-shadow: 0 8px 30px rgba(0, 0, 0, 0.2);
    z-index: 1001;
    opacity: 0;
    visibility: hidden;
    transition: opacity 0.15s ease, visibility 0.15s ease, transform 0.15s ease;
    overflow: hidden;
}

.shortcuts-modal.active {
    opacity: 1;
    visibility: visible;
    transform: translate(-50%, -50%) scale(1);
}

.shortcuts-modal-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 16px 20px;
    border-bottom: 1px solid #d1d9e0;
}

.shortcuts-modal-header h2 {
    margin: 0;
    font-size: 16px;
    font-weight: 600;
    color: #1f2328;
}

.shortcuts-modal-close {
    display: flex;
    align-items: center;
    justify-content: center;
    background: none;
    border: none;
    cursor: pointer;
    color: #59636e;
    padding: 4px;
    border-radius: 4px;
    transition: all 0.15s ease;
}

.shortcuts-modal-close:hover {
    color: #1f2328;
    background: #e6e8eb;
}

.shortcuts-modal-content {
    padding: 16px 20px;
    overflow-y: auto;
    max-height: calc(80vh - 60px);
}

.shortcut-row {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 8px 0;
}

.shortcut-row:not(:last-child) {
    border-bottom: 1px solid #e6e8eb;
}

.shortcut-action {
    font-size: 14px;
    color: #1f2328;
}

.shortcut-keys {
    display: flex;
    gap: 4px;
}

.shortcut-keys kbd {
    display: inline-block;
    padding: 4px 8px;
    background: #f6f8fa;
    border: 1px solid #d1d9e0;
    border-radius: 6px;
    font-family: inherit;
    font-size: 12px;
    font-weight: 500;
    color: #1f2328;
}

/* Dark Mode */
@media (prefers-color-scheme: dark) {
    .comment-btn {
        background: #58a6ff;
    }

    .comment-highlight {
        background-color: rgba(255, 212, 59, 0.25);
    }

    .comment-highlight:hover,
    .comment-highlight.active {
        background-color: rgba(255, 212, 59, 0.45);
    }

    .comments-panel {
        background: #161b22;
        border-left-color: #3d444d;
    }

    .comments-panel-header {
        border-bottom-color: #3d444d;
    }

    .comments-panel-header h2 {
        color: #9198a1;
    }

    .comments-panel-close-btn {
        color: #9198a1;
    }

    .comments-panel-close-btn:hover {
        color: #e6edf3;
        background: #21262d;
    }

    .comments-list::-webkit-scrollbar-thumb {
        background-color: #3d444d;
    }

    .comment-entry {
        border-bottom-color: #3d444d;
    }

    .comment-entry:hover {
        background: #21262d;
    }

    .comment-entry.active {
        background: #388bfd26;
    }

    .comment-quote {
        background: #21262d;
        border-left-color: #58a6ff;
        color: #9198a1;
    }

    .comment-text {
        color: #e6edf3;
    }

    .comment-copy-btn,
    .comment-edit-btn,
    .comment-delete-btn {
        color: #9198a1;
    }

    .comment-copy-btn:hover,
    .comment-edit-btn:hover {
        color: #58a6ff;
        background: rgba(88, 166, 255, 0.1);
    }

    .comment-copy-btn.copied {
        color: #3fb950;
    }

    .comment-delete-btn:hover {
        color: #f85149;
        background: rgba(248, 81, 73, 0.1);
    }

    .comment-input-form {
        background: #161b22;
        border-bottom-color: #3d444d;
    }

    .comment-input-quote {
        background: #21262d;
        border-left-color: #58a6ff;
        color: #9198a1;
    }

    .comment-input-textarea {
        background: #0d1117;
        border-color: #3d444d;
        color: #e6edf3;
    }

    .comment-input-textarea:focus {
        border-color: #58a6ff;
        box-shadow: 0 0 0 3px rgba(88, 166, 255, 0.1);
    }

    .comment-cancel-btn {
        border-color: #3d444d;
        color: #e6edf3;
    }

    .comment-cancel-btn:hover {
        background: #21262d;
    }

    .comment-save-btn {
        background: #58a6ff;
    }

    .open-comments-btn {
        background: #161b22;
        border-color: #3d444d;
        color: #9198a1;
        box-shadow: 0 2px 8px rgba(0, 0, 0, 0.3);
    }

    .open-comments-btn:hover {
        color: #e6edf3;
        background: #21262d;
    }

    .comment-count {
        background: #58a6ff;
    }

    .comments-panel-footer {
        border-top-color: #3d444d;
    }

    .copy-comments-btn {
        background: #21262d;
        border-color: #3d444d;
        color: #9198a1;
    }

    .copy-comments-btn:hover {
        color: #e6edf3;
        background: #30363d;
    }

    .copy-comments-btn.copied {
        color: #3fb950;
    }

    .comments-empty {
        color: #9198a1;
    }

    .comments-empty kbd {
        background: #21262d;
    }

    .shortcuts-modal-overlay {
        background: rgba(0, 0, 0, 0.7);
    }

    .shortcuts-modal {
        background: #161b22;
        box-shadow: 0 8px 30px rgba(0, 0, 0, 0.4);
    }

    .shortcuts-modal-header {
        border-bottom-color: #3d444d;
    }

    .shortcuts-modal-header h2 {
        color: #e6edf3;
    }

    .shortcuts-modal-close {
        color: #9198a1;
    }

    .shortcuts-modal-close:hover {
        color: #e6edf3;
        background: #21262d;
    }

    .shortcut-row:not(:last-child) {
        border-bottom-color: #3d444d;
    }

    .shortcut-action {
        color: #e6edf3;
    }

    .shortcut-keys kbd {
        background: #21262d;
        border-color: #3d444d;
        color: #e6edf3;
    }
}

/* Mobile Responsive */
@media (max-width: 768px) {
    .comments-panel {
        width: 100%;
    }

    .comment-btn {
        font-size: 12px;
        padding: 6px 10px;
    }

    .open-comments-btn {
        top: auto;
        bottom: 16px;
    }
}
`

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
            left: 16px;
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
    </a>
    %s
    %s
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

const commentsJS = `
    <script>
        (function() {
            'use strict';

            // Icons
            var commentIcon = '<svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"></path></svg>';
            var closeIcon = '<svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="18" y1="6" x2="6" y2="18"></line><line x1="6" y1="6" x2="18" y2="18"></line></svg>';
            var deleteIcon = '<svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="3 6 5 6 21 6"></polyline><path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"></path></svg>';
            var copyIcon = '<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg>';
            var checkIcon = '<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 6 9 17 4 12"></polyline></svg>';

            // State
            var STORAGE_KEY = 'mdp-comments-' + window.location.pathname;
            var comments = [];
            var commentCounter = 0;
            var currentRange = null;
            var currentSelectedText = '';

            // DOM Elements
            var commentBtn = document.querySelector('.comment-btn');
            var commentsPanel = document.querySelector('.comments-panel');
            var commentsList = document.querySelector('.comments-list');
            var commentInputForm = document.querySelector('.comment-input-form');
            var commentInputQuote = document.querySelector('.comment-input-quote');
            var commentInputTextarea = document.querySelector('.comment-input-textarea');
            var copyCommentsBtn = document.querySelector('.copy-comments-btn');
            var openCommentsBtn = document.querySelector('.open-comments-btn');
            var commentCount = document.querySelector('.comment-count');
            var commentsPanelCloseBtn = document.querySelector('.comments-panel-close-btn');
            var commentCancelBtn = document.querySelector('.comment-cancel-btn');
            var commentSaveBtn = document.querySelector('.comment-save-btn');
            var commentsEmpty = document.querySelector('.comments-empty');
            var shortcutsOverlay = document.querySelector('.shortcuts-modal-overlay');
            var shortcutsModal = document.querySelector('.shortcuts-modal');
            var shortcutsCloseBtn = document.querySelector('.shortcuts-modal-close');

            // Utility functions
            function escapeHtml(text) {
                var div = document.createElement('div');
                div.textContent = text;
                return div.innerHTML;
            }

            function truncateText(text, maxLen) {
                if (text.length <= maxLen) return text;
                return text.substring(0, maxLen) + '...';
            }

            function generateCommentId() {
                return 'comment-' + (++commentCounter);
            }

            // SessionStorage functions
            function saveCommentsToStorage() {
                try {
                    sessionStorage.setItem(STORAGE_KEY, JSON.stringify(comments));
                } catch (e) {
                    // Storage might be full or disabled
                }
            }

            function loadCommentsFromStorage() {
                try {
                    var stored = sessionStorage.getItem(STORAGE_KEY);
                    if (stored) {
                        comments = JSON.parse(stored);
                        // Update counter to avoid ID collisions
                        comments.forEach(function(c) {
                            var num = parseInt(c.id.replace('comment-', ''), 10);
                            if (num >= commentCounter) commentCounter = num;
                        });
                    }
                } catch (e) {
                    comments = [];
                }
            }

            function findTextInDocument(searchText) {
                var markdownBody = document.querySelector('.markdown-body');
                if (!markdownBody) return null;

                var treeWalker = document.createTreeWalker(
                    markdownBody,
                    NodeFilter.SHOW_TEXT,
                    null
                );

                var textContent = '';
                var textNodes = [];
                var node;

                while (node = treeWalker.nextNode()) {
                    textNodes.push({
                        node: node,
                        start: textContent.length,
                        end: textContent.length + node.textContent.length
                    });
                    textContent += node.textContent;
                }

                var searchIndex = textContent.indexOf(searchText);
                if (searchIndex === -1) return null;

                var startNode = null, endNode = null;
                var startOffset = 0, endOffset = 0;
                var searchEnd = searchIndex + searchText.length;

                for (var i = 0; i < textNodes.length; i++) {
                    var tn = textNodes[i];
                    if (!startNode && searchIndex >= tn.start && searchIndex < tn.end) {
                        startNode = tn.node;
                        startOffset = searchIndex - tn.start;
                    }
                    if (searchEnd > tn.start && searchEnd <= tn.end) {
                        endNode = tn.node;
                        endOffset = searchEnd - tn.start;
                        break;
                    }
                }

                if (!startNode || !endNode) return null;

                var range = document.createRange();
                range.setStart(startNode, startOffset);
                range.setEnd(endNode, endOffset);
                return range;
            }

            function restoreHighlightsFromStorage() {
                comments.forEach(function(comment) {
                    var range = findTextInDocument(comment.selectedText);
                    if (range) {
                        highlightRange(range, comment.id);
                    }
                });
            }

            function renderStoredComments() {
                comments.forEach(function(comment) {
                    addCommentEntry(comment);
                });
            }

            // Selection handling
            function isValidSelection(selection) {
                if (!selection || selection.isCollapsed) return false;
                var range = selection.getRangeAt(0);
                var container = range.commonAncestorContainer;
                var markdownBody = document.querySelector('.markdown-body');
                if (!markdownBody || !markdownBody.contains(container)) return false;
                if (commentsPanel && commentsPanel.contains(container)) return false;
                return true;
            }

            function positionCommentButton(range) {
                var rect = range.getBoundingClientRect();
                commentBtn.style.top = (rect.top + window.scrollY - 40) + 'px';
                commentBtn.style.left = (rect.left + window.scrollX) + 'px';
            }

            function showCommentButton() {
                commentBtn.classList.add('visible');
            }

            function hideCommentButton() {
                commentBtn.classList.remove('visible');
            }

            // Highlight creation
            function highlightRange(range, commentId) {
                try {
                    var mark = document.createElement('mark');
                    mark.className = 'comment-highlight';
                    mark.dataset.commentId = commentId;
                    range.surroundContents(mark);
                    return [mark];
                } catch (e) {
                    return highlightComplexRange(range, commentId);
                }
            }

            function highlightComplexRange(range, commentId) {
                var marks = [];
                var treeWalker = document.createTreeWalker(
                    range.commonAncestorContainer,
                    NodeFilter.SHOW_TEXT,
                    null
                );

                var textNodes = [];
                var node;
                var inRange = false;

                while (node = treeWalker.nextNode()) {
                    if (node === range.startContainer) inRange = true;
                    if (inRange) textNodes.push(node);
                    if (node === range.endContainer) break;
                }

                for (var i = textNodes.length - 1; i >= 0; i--) {
                    var textNode = textNodes[i];
                    var nodeRange = document.createRange();

                    if (textNode === range.startContainer) {
                        nodeRange.setStart(textNode, range.startOffset);
                    } else {
                        nodeRange.setStart(textNode, 0);
                    }

                    if (textNode === range.endContainer) {
                        nodeRange.setEnd(textNode, range.endOffset);
                    } else {
                        nodeRange.setEnd(textNode, textNode.length);
                    }

                    if (!nodeRange.collapsed) {
                        try {
                            var mark = document.createElement('mark');
                            mark.className = 'comment-highlight';
                            mark.dataset.commentId = commentId;
                            nodeRange.surroundContents(mark);
                            marks.unshift(mark);
                        } catch (err) {
                            // Skip nodes that can't be wrapped
                        }
                    }
                }
                return marks;
            }

            // Comments panel
            function openCommentsPanel() {
                commentsPanel.classList.add('open');
            }

            function closeCommentsPanel() {
                commentsPanel.classList.remove('open');
                hideInputForm();
            }

            function showInputForm(selectedText) {
                commentInputQuote.textContent = truncateText(selectedText, 150);
                commentInputTextarea.value = '';
                commentInputForm.style.display = 'block';
                // Use double-rAF to ensure focus after panel animation starts
                requestAnimationFrame(function() {
                    requestAnimationFrame(function() {
                        commentInputTextarea.focus();
                    });
                });
            }

            function hideInputForm() {
                commentInputForm.style.display = 'none';
                commentInputTextarea.value = '';
            }

            var smallCopyIcon = '<svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg>';
            var smallCheckIcon = '<svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 6 9 17 4 12"></polyline></svg>';
            var editIcon = '<svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M17 3a2.85 2.83 0 1 1 4 4L7.5 20.5 2 22l1.5-5.5Z"></path><path d="m15 5 4 4"></path></svg>';

            function copyIndividualComment(comment, btn) {
                var quotedText = comment.selectedText.split('\n').map(function(line) {
                    return '> ' + line;
                }).join('\n');
                var markdown = quotedText + '\n\nComment: ' + comment.commentText;

                navigator.clipboard.writeText(markdown).then(function() {
                    btn.innerHTML = smallCheckIcon;
                    btn.classList.add('copied');
                    setTimeout(function() {
                        btn.innerHTML = smallCopyIcon;
                        btn.classList.remove('copied');
                    }, 2000);
                });
            }

            var editingCommentId = null;

            function editComment(commentId) {
                var comment = comments.find(function(c) { return c.id === commentId; });
                if (!comment) return;

                editingCommentId = commentId;
                openCommentsPanel();
                commentInputQuote.textContent = truncateText(comment.selectedText, 150);
                commentInputTextarea.value = comment.commentText;
                commentInputForm.style.display = 'block';
                requestAnimationFrame(function() {
                    requestAnimationFrame(function() {
                        commentInputTextarea.focus();
                        commentInputTextarea.select();
                    });
                });
            }

            function addCommentEntry(comment) {
                var entry = document.createElement('div');
                entry.className = 'comment-entry';
                entry.dataset.commentId = comment.id;
                entry.innerHTML =
                    '<div class="comment-quote">' + escapeHtml(truncateText(comment.selectedText, 150)) + '</div>' +
                    '<div class="comment-text">' + escapeHtml(comment.commentText) + '</div>' +
                    '<div class="comment-actions">' +
                        '<button class="comment-copy-btn" aria-label="Copy comment">' + smallCopyIcon + '</button>' +
                        '<button class="comment-edit-btn" aria-label="Edit comment">' + editIcon + '</button>' +
                        '<button class="comment-delete-btn" aria-label="Delete comment">' + deleteIcon + '</button>' +
                    '</div>';

                commentsList.appendChild(entry);

                entry.querySelector('.comment-copy-btn').addEventListener('click', function(e) {
                    e.stopPropagation();
                    copyIndividualComment(comment, this);
                });

                entry.querySelector('.comment-edit-btn').addEventListener('click', function(e) {
                    e.stopPropagation();
                    editComment(comment.id);
                });

                entry.querySelector('.comment-delete-btn').addEventListener('click', function(e) {
                    e.stopPropagation();
                    deleteComment(comment.id);
                });

                entry.addEventListener('click', function(e) {
                    if (!e.target.closest('.comment-actions')) {
                        scrollToHighlight(comment.id);
                    }
                });
            }

            function deleteComment(commentId) {
                comments = comments.filter(function(c) { return c.id !== commentId; });

                var marks = document.querySelectorAll('.comment-highlight[data-comment-id="' + commentId + '"]');
                marks.forEach(function(mark) {
                    var parent = mark.parentNode;
                    while (mark.firstChild) {
                        parent.insertBefore(mark.firstChild, mark);
                    }
                    parent.removeChild(mark);
                    parent.normalize();
                });

                var entry = commentsList.querySelector('.comment-entry[data-comment-id="' + commentId + '"]');
                if (entry) entry.remove();

                updateCommentsUI();
                saveCommentsToStorage();
            }

            function scrollToHighlight(commentId) {
                var mark = document.querySelector('.comment-highlight[data-comment-id="' + commentId + '"]');
                if (mark) {
                    mark.scrollIntoView({ behavior: 'smooth', block: 'center' });
                    mark.classList.add('active');
                    setTimeout(function() {
                        mark.classList.remove('active');
                    }, 2000);
                }
            }

            // Copy to Markdown
            function copyCommentsAsMarkdown() {
                if (comments.length === 0) return;

                var markdown = comments.map(function(c) {
                    var quotedText = c.selectedText.split('\n').map(function(line) {
                        return '> ' + line;
                    }).join('\n');
                    return quotedText + '\n\nComment: ' + c.commentText;
                }).join('\n\n---\n\n');

                navigator.clipboard.writeText(markdown).then(function() {
                    copyCommentsBtn.innerHTML = checkIcon + '<span>Copied!</span>';
                    copyCommentsBtn.classList.add('copied');
                    setTimeout(function() {
                        copyCommentsBtn.innerHTML = copyIcon + '<span>Copy Comments</span>';
                        copyCommentsBtn.classList.remove('copied');
                    }, 2000);
                });
            }

            function updateCommentsUI() {
                var count = comments.length;
                commentCount.textContent = count;
                if (count > 0) {
                    commentCount.classList.add('visible');
                    commentsEmpty.style.display = 'none';
                } else {
                    commentCount.classList.remove('visible');
                    commentsEmpty.style.display = 'block';
                }
            }

            // Shortcuts modal functions
            function openShortcutsModal() {
                shortcutsOverlay.classList.add('active');
                shortcutsModal.classList.add('active');
            }

            function closeShortcutsModal() {
                shortcutsOverlay.classList.remove('active');
                shortcutsModal.classList.remove('active');
            }

            // Event listeners
            document.addEventListener('mouseup', function(e) {
                setTimeout(function() {
                    var selection = window.getSelection();
                    if (isValidSelection(selection)) {
                        currentRange = selection.getRangeAt(0).cloneRange();
                        currentSelectedText = selection.toString().trim();
                        positionCommentButton(currentRange);
                        showCommentButton();
                    }
                }, 10);
            });

            document.addEventListener('mousedown', function(e) {
                if (!commentBtn.contains(e.target) && !commentsPanel.contains(e.target)) {
                    hideCommentButton();
                }
            });

            document.addEventListener('selectionchange', function() {
                var selection = window.getSelection();
                if (!isValidSelection(selection)) {
                    hideCommentButton();
                }
            });

            commentBtn.addEventListener('click', function() {
                if (!currentRange || !currentSelectedText) return;

                hideCommentButton();
                openCommentsPanel();
                showInputForm(currentSelectedText);

                window.getSelection().removeAllRanges();
            });

            function saveComment() {
                var commentText = commentInputTextarea.value.trim();
                if (!commentText) return;

                // Check if we're editing an existing comment
                if (editingCommentId) {
                    var comment = comments.find(function(c) { return c.id === editingCommentId; });
                    if (comment) {
                        comment.commentText = commentText;
                        // Update the entry in the DOM
                        var entry = commentsList.querySelector('.comment-entry[data-comment-id="' + editingCommentId + '"]');
                        if (entry) {
                            entry.querySelector('.comment-text').textContent = commentText;
                        }
                        saveCommentsToStorage();
                    }
                    editingCommentId = null;
                    hideInputForm();
                    return;
                }

                // Creating a new comment
                if (!currentRange) return;

                var commentId = generateCommentId();

                highlightRange(currentRange, commentId);

                var comment = {
                    id: commentId,
                    selectedText: currentSelectedText,
                    commentText: commentText
                };
                comments.push(comment);

                addCommentEntry(comment);
                hideInputForm();
                updateCommentsUI();
                saveCommentsToStorage();

                currentRange = null;
                currentSelectedText = '';
            }

            commentSaveBtn.addEventListener('click', saveComment);

            // Ctrl/Cmd + Enter to save comment
            commentInputTextarea.addEventListener('keydown', function(e) {
                if (e.key === 'Enter' && (e.metaKey || e.ctrlKey)) {
                    e.preventDefault();
                    saveComment();
                }
            });

            commentCancelBtn.addEventListener('click', function() {
                hideInputForm();
                editingCommentId = null;
                currentRange = null;
                currentSelectedText = '';
            });

            commentsPanelCloseBtn.addEventListener('click', closeCommentsPanel);
            copyCommentsBtn.addEventListener('click', copyCommentsAsMarkdown);
            openCommentsBtn.addEventListener('click', function() {
                if (commentsPanel.classList.contains('open')) {
                    closeCommentsPanel();
                } else {
                    openCommentsPanel();
                }
            });

            // Click on highlight to show comment
            document.addEventListener('click', function(e) {
                var highlight = e.target.closest('.comment-highlight');
                if (highlight) {
                    var commentId = highlight.dataset.commentId;
                    openCommentsPanel();
                    var entry = commentsList.querySelector('.comment-entry[data-comment-id="' + commentId + '"]');
                    if (entry) {
                        entry.scrollIntoView({ behavior: 'smooth', block: 'center' });
                        entry.classList.add('active');
                        setTimeout(function() {
                            entry.classList.remove('active');
                        }, 2000);
                    }
                }
            });

            // Shortcuts modal event listeners
            shortcutsCloseBtn.addEventListener('click', closeShortcutsModal);
            shortcutsOverlay.addEventListener('click', function(e) {
                if (e.target === shortcutsOverlay) {
                    closeShortcutsModal();
                }
            });

            // Keyboard shortcuts (use capture phase to intercept before browser)
            document.addEventListener('keydown', function(e) {
                // Don't trigger shortcuts when typing in an input/textarea
                var isTyping = e.target.tagName === 'INPUT' || e.target.tagName === 'TEXTAREA';

                // Escape to close modals/panels
                if (e.key === 'Escape') {
                    if (shortcutsModal.classList.contains('active')) {
                        closeShortcutsModal();
                        return;
                    }
                    if (commentsPanel.classList.contains('open')) {
                        closeCommentsPanel();
                        return;
                    }
                }

                // '?' (Shift+/) to show keyboard shortcuts
                if (e.key === '?' && !isTyping) {
                    e.preventDefault();
                    openShortcutsModal();
                    return;
                }

                // Cmd+/ (Mac) or Ctrl+/ (Win/Linux) to toggle comments panel
                if (e.key === '/' && (e.metaKey || e.ctrlKey)) {
                    e.preventDefault();
                    e.stopPropagation();
                    if (commentsPanel.classList.contains('open')) {
                        closeCommentsPanel();
                    } else {
                        openCommentsPanel();
                    }
                    return;
                }

                // 'C' key to trigger comment when selection is active
                if (e.key === 'c' && !e.metaKey && !e.ctrlKey && !e.altKey && !isTyping) {
                    // Only trigger if comment button is visible (has selection)
                    if (commentBtn.classList.contains('visible')) {
                        e.preventDefault();
                        commentBtn.click();
                    }
                }
            }, true);

            // Initialize - load comments from sessionStorage
            loadCommentsFromStorage();
            if (comments.length > 0) {
                restoreHighlightsFromStorage();
                renderStoredComments();
            }
            updateCommentsUI();
        })();
    </script>`

const commentsHTML = `
    <!-- Comment Button -->
    <button class="comment-btn" aria-label="Add comment (C)">
        <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"></path></svg>
        <span>Comment</span>
        <kbd>C</kbd>
    </button>

    <!-- Open Comments Button -->
    <button class="open-comments-btn" title="Toggle comments (⌘/)">
        <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"></path></svg>
        <span class="comment-count">0</span>
    </button>

    <!-- Comments Panel -->
    <aside class="comments-panel" role="complementary" aria-label="Comments">
        <div class="comments-panel-header">
            <h2>Comments</h2>
            <button class="comments-panel-close-btn" aria-label="Close comments">
                <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="18" y1="6" x2="6" y2="18"></line><line x1="6" y1="6" x2="18" y2="18"></line></svg>
            </button>
        </div>
        <div class="comments-list">
            <div class="comments-empty">
                <svg xmlns="http://www.w3.org/2000/svg" width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"></path></svg>
                <p>No comments yet</p>
                <p>Select text and press <kbd>C</kbd> to add a comment</p>
            </div>
            <div class="comment-input-form" style="display: none;">
                <div class="comment-input-quote"></div>
                <textarea class="comment-input-textarea" placeholder="Add your comment..." rows="3"></textarea>
                <div class="comment-input-actions">
                    <button class="comment-cancel-btn">Cancel</button>
                    <button class="comment-save-btn">Save <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="9 10 4 15 9 20"></polyline><path d="M20 4v7a4 4 0 0 1-4 4H4"></path></svg></button>
                </div>
            </div>
        </div>
        <div class="comments-panel-footer">
            <button class="copy-comments-btn" title="Copy all comments as Markdown">
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg>
                <span>Copy All</span>
            </button>
        </div>
    </aside>

    <!-- Keyboard Shortcuts Modal -->
    <div class="shortcuts-modal-overlay" role="dialog" aria-modal="true" aria-labelledby="shortcuts-modal-title">
        <div class="shortcuts-modal">
            <div class="shortcuts-modal-header">
                <h2 id="shortcuts-modal-title">Keyboard Shortcuts</h2>
                <button class="shortcuts-modal-close" aria-label="Close shortcuts">
                    <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="18" y1="6" x2="6" y2="18"></line><line x1="6" y1="6" x2="18" y2="18"></line></svg>
                </button>
            </div>
            <div class="shortcuts-modal-content">
                <div class="shortcut-row">
                    <span class="shortcut-action">Show keyboard shortcuts</span>
                    <span class="shortcut-keys"><kbd>?</kbd></span>
                </div>
                <div class="shortcut-row">
                    <span class="shortcut-action">Add comment to selection</span>
                    <span class="shortcut-keys"><kbd>C</kbd></span>
                </div>
                <div class="shortcut-row">
                    <span class="shortcut-action">Toggle comments panel</span>
                    <span class="shortcut-keys"><kbd>⌘</kbd><kbd>/</kbd></span>
                </div>
                <div class="shortcut-row">
                    <span class="shortcut-action">Close panel / modal</span>
                    <span class="shortcut-keys"><kbd>Esc</kbd></span>
                </div>
                <div class="shortcut-row">
                    <span class="shortcut-action">Save comment</span>
                    <span class="shortcut-keys"><kbd>⌘</kbd><kbd>↵</kbd></span>
                </div>
            </div>
        </div>
    </div>`

// Generate creates a complete HTML document with the given title and content.
func Generate(title, content string) string {
	scripts := copyButtonScript + commentsJS
	return fmt.Sprintf(htmlTemplate, title, githubMarkdownCSS, chromaCSS, commentsCSS, content, commentsHTML, scripts)
}

// GenerateWithLiveReload creates an HTML document with live reload support.
func GenerateWithLiveReload(title, content string, port int) string {
	scripts := copyButtonScript + commentsJS + fmt.Sprintf(liveReloadScript, port)
	return fmt.Sprintf(htmlTemplate, title, githubMarkdownCSS, chromaCSS, commentsCSS, content, commentsHTML, scripts)
}
