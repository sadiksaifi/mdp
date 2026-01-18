package template

import (
	"fmt"
	"html"
	"strings"

	"mdp/internal/filetree"
)

const multiFileTemplate = `<!DOCTYPE html>
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
</head>
<body>
    <script>if(history.scrollRestoration)history.scrollRestoration='manual';window.scrollTo(0,0);</script>

    <!-- Desktop Top Bar -->
    <header class="topbar">
        <div class="topbar-left">
            <button class="topbar-btn topbar-sidebar-btn" aria-label="Toggle sidebar" title="Toggle sidebar (⌘B)">
                <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M4 6a2 2 0 0 1 2 -2h12a2 2 0 0 1 2 2v12a2 2 0 0 1 -2 2h-12a2 2 0 0 1 -2 -2l0 -12" /><path d="M9 4l0 16" /></svg>
            </button>
        </div>
        <div class="topbar-center">
            <span class="topbar-brand">MARKDOWN PREVIEW</span>
        </div>
        <div class="topbar-right">
            <button class="topbar-btn topbar-search-btn" aria-label="Search files" title="Search files (⌘K)">
                <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M10 10m-7 0a7 7 0 1 0 14 0a7 7 0 1 0 -14 0" /><path d="M21 21l-6 -6" /></svg>
            </button>
            <div class="topbar-divider"></div>
            <button class="topbar-btn topbar-comment-btn" aria-label="Toggle comments" title="Toggle comments (⌘/)">
                <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"></path></svg>
                <span class="topbar-comment-count">0</span>
            </button>
            <div class="topbar-divider"></div>
            <button class="topbar-btn topbar-help-btn" aria-label="Keyboard shortcuts" title="Keyboard shortcuts (?)">
                <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><path d="M9.09 9a3 3 0 0 1 5.83 1c0 2-3 3-3 3"></path><line x1="12" y1="17" x2="12.01" y2="17"></line></svg>
            </button>
        </div>
    </header>

    <!-- Mobile Top Bar (floating buttons) -->
    <div class="floating-buttons">
        <button class="sidebar-open-btn" aria-label="Open sidebar" title="Open sidebar">
            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M4 6a2 2 0 0 1 2 -2h12a2 2 0 0 1 2 2v12a2 2 0 0 1 -2 2h-12a2 2 0 0 1 -2 -2l0 -12" /><path d="M15 4l0 16" /></svg>
        </button>
        <span class="topbar-title">MARKDOWN PREVIEW</span>
        <button class="search-open-btn" aria-label="Search files" title="Search files">
            <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M10 10m-7 0a7 7 0 1 0 14 0a7 7 0 1 0 -14 0" /><path d="M21 21l-6 -6" /></svg>
        </button>
    </div>

    <div class="sidebar-overlay"></div>

    <aside class="sidebar" role="navigation" aria-label="File navigation">
        <div class="sidebar-header">
            <h2>Files</h2>
        </div>
        <nav class="file-tree">
            %s
        </nav>
        <div class="sidebar-footer">
            <a href="https://github.com/sadiksaifi/mdp" target="_blank" rel="noopener noreferrer" title="Star on GitHub">
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="currentColor"><path d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"/></svg>
                <span>Star on GitHub</span>
            </a>
        </div>
    </aside>

    <div class="search-palette-overlay" id="search-overlay"></div>
    <div class="search-palette" role="dialog" aria-label="File search" id="search-palette">
        <input type="text" class="search-palette-input" id="search-input" placeholder="Search files..." autocomplete="off">
        <ul class="search-palette-results" id="search-results"></ul>
        <div class="search-palette-hint">
            <span><kbd>↑↓</kbd> Navigate</span>
            <span><kbd>Enter</kbd> Open</span>
            <span><kbd>Esc</kbd> Close</span>
        </div>
    </div>

    <main class="content">
        %s
    </main>

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
                    <span class="shortcut-action">Search files</span>
                    <span class="shortcut-keys"><kbd>⌘</kbd><kbd>K</kbd></span>
                </div>
                <div class="shortcut-row">
                    <span class="shortcut-action">Toggle sidebar</span>
                    <span class="shortcut-keys"><kbd>⌘</kbd><kbd>B</kbd></span>
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
    </div>

    <script>
        %s
    </script>
    %s
</body>
</html>`

const sidebarCSS = `
:root {
    --sidebar-width: 280px;
    --sidebar-bg: #f6f8fa;
    --sidebar-border: #d1d9e0;
    --sidebar-hover: #e6e8eb;
    --sidebar-active: #0969da;
    --sidebar-active-bg: #ddf4ff;
    --content-padding: 45px;
    --transition-speed: 0.3s;
    --fg-color: #1f2328;
    --fg-muted: #59636e;
    --panel-section-height: 44px;
    --topbar-height: 56px;
}

@media (prefers-color-scheme: dark) {
    :root {
        --sidebar-bg: #161b22;
        --sidebar-border: #3d444d;
        --sidebar-hover: #21262d;
        --sidebar-active: #58a6ff;
        --sidebar-active-bg: #388bfd26;
        --fg-color: #e6edf3;
        --fg-muted: #9198a1;
    }
    body {
        background-color: #0d1117;
    }
}

* {
    box-sizing: border-box;
}

html {
    scroll-behavior: auto;
}

:target {
    scroll-margin-top: 100vh;
}

body {
    margin: 0;
    padding: 0;
    max-width: none;
    display: flex;
    min-height: 100vh;
    color: var(--fg-color);
    font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", "Noto Sans", Helvetica, Arial, sans-serif;
}

/* Desktop Top Bar */
.topbar {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    height: var(--topbar-height);
    background: var(--sidebar-bg);
    border-bottom: 1px solid var(--sidebar-border);
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0 4px 0 16px;
    z-index: 250;
}

.topbar-left,
.topbar-right {
    display: flex;
    align-items: center;
    gap: 4px;
}

.topbar-center {
    position: absolute;
    left: 50%;
    transform: translateX(-50%);
    display: flex;
    align-items: center;
    gap: 8px;
}

.topbar-brand {
    font-size: 16px;
    font-weight: 700;
    color: var(--fg-color);
    letter-spacing: 1px;
}

.topbar-brand-sub {
    font-size: 11px;
    font-weight: 500;
    color: var(--fg-muted);
    text-transform: uppercase;
    letter-spacing: 0.5px;
}

.topbar-btn {
    display: flex;
    align-items: center;
    justify-content: center;
    background: transparent;
    border: none;
    border-radius: 6px;
    padding: 8px;
    cursor: pointer;
    color: var(--fg-muted);
    transition: all 0.2s ease;
}

.topbar-btn:hover {
    color: var(--fg-color);
    background: var(--sidebar-hover);
}

.topbar-btn svg {
    width: 20px;
    height: 20px;
}

.topbar-btn.active {
    color: var(--sidebar-active);
    background: var(--sidebar-active-bg);
}

.topbar-comment-btn {
    position: relative;
}

.topbar-comment-count {
    position: absolute;
    top: 2px;
    right: 2px;
    min-width: 16px;
    height: 16px;
    padding: 0 4px;
    background: var(--sidebar-active);
    color: #ffffff;
    border-radius: 8px;
    font-size: 10px;
    font-weight: 600;
    display: none;
    align-items: center;
    justify-content: center;
}

.topbar-comment-count.visible {
    display: flex;
}

.topbar-divider {
    width: 1px;
    height: 24px;
    background: var(--sidebar-border);
    margin: 0 4px;
}

.sidebar {
    position: fixed;
    left: 0;
    top: var(--topbar-height);
    bottom: 0;
    width: var(--sidebar-width);
    background: var(--sidebar-bg);
    border-right: 1px solid var(--sidebar-border);
    display: flex;
    flex-direction: column;
    overflow: hidden;
    z-index: 200;
    transition: transform var(--transition-speed) ease, opacity var(--transition-speed) ease;
}

.sidebar.collapsed {
    transform: translateX(-100%);
}

.sidebar-header {
    display: flex;
    align-items: center;
    height: var(--panel-section-height);
    padding: 0 16px;
    flex-shrink: 0;
}

.sidebar-header h2 {
    margin: 0;
    font-size: 12px;
    font-weight: 600;
    color: var(--fg-muted);
    text-transform: uppercase;
    letter-spacing: 0.5px;
}

.sidebar-footer {
    display: flex;
    align-items: center;
    justify-content: flex-start;
    height: var(--panel-section-height);
    padding: 0 16px;
    border-top: 1px solid var(--sidebar-border);
    flex-shrink: 0;
}

.sidebar-footer a {
    display: flex;
    align-items: center;
    gap: 8px;
    height: 100%;
    color: var(--fg-muted);
    text-decoration: none;
    font-size: 12px;
    transition: color 0.2s ease;
}

.sidebar-footer a:hover {
    color: var(--fg-color);
}

.sidebar-footer svg {
    flex-shrink: 0;
}

/* Mobile floating buttons (hidden on desktop) */
.floating-buttons {
    display: none;
}

.topbar-title {
    display: none;
}

.file-tree {
    padding: 0;
    flex: 1;
    overflow-y: auto;
    scrollbar-width: thin;
    scrollbar-color: var(--sidebar-border) transparent;
}

.file-tree::-webkit-scrollbar {
    width: 8px;
}

.file-tree::-webkit-scrollbar-track {
    background: transparent;
}

.file-tree::-webkit-scrollbar-thumb {
    background-color: var(--sidebar-border);
    border-radius: 4px;
}

.file-tree::-webkit-scrollbar-thumb:hover {
    background-color: var(--fg-muted);
}

.file-tree ul {
    list-style: none;
    margin: 0;
    padding: 0;
}

.file-tree li {
    margin: 0;
}

.file-tree a {
    display: block;
    padding: 8px 16px;
    color: var(--fg-color);
    text-decoration: none;
    font-size: 14px;
    border-left: 3px solid transparent;
    transition: all 0.15s ease;
}

.file-tree a:hover {
    background: var(--sidebar-hover);
}

.file-tree a.active {
    background: var(--sidebar-active-bg);
    border-left-color: var(--sidebar-active);
    color: var(--sidebar-active);
    font-weight: 500;
}

.file-tree .directory > span {
    display: flex;
    align-items: center;
    padding: 8px 16px;
    font-size: 13px;
    font-weight: 600;
    color: var(--fg-muted);
    cursor: pointer;
    user-select: none;
}

.file-tree .directory > span:hover {
    background: var(--sidebar-hover);
}

.file-tree .directory > span::before {
    content: '\25B6';
    display: inline-block;
    margin-right: 6px;
    font-size: 10px;
    transition: transform 0.2s ease;
}

.file-tree .directory.open > span::before {
    transform: rotate(90deg);
}

.file-tree .directory > ul {
    padding-left: 16px;
    display: none;
}

.file-tree .directory.open > ul {
    display: block;
}

.sidebar.collapsed ~ .content,
body:has(.sidebar.collapsed) .content {
    margin-left: auto;
    margin-right: auto;
}

.content {
    transition: margin-left var(--transition-speed) ease;
    flex: 1;
    margin-left: auto;
    margin-right: auto;
    padding: calc(var(--topbar-height) + var(--content-padding)) var(--content-padding) var(--content-padding);
    max-width: 980px;
    min-width: 0;
}

.content-section {
    display: none;
}

.content-section.active {
    display: block;
}

.sidebar-overlay {
    display: none;
    position: fixed;
    inset: 0;
    background: rgba(0, 0, 0, 0.5);
    z-index: 150;
    opacity: 0;
    transition: opacity var(--transition-speed) ease;
}

.sidebar-overlay.active {
    opacity: 1;
}

@media (max-width: 768px) {
    /* Hide desktop top bar on mobile */
    .topbar {
        display: none;
    }

    .sidebar {
        top: 0;
        transform: translateX(-100%);
        z-index: 300;
    }

    .sidebar.open {
        transform: translateX(0);
    }

    .sidebar.collapsed {
        transform: translateX(-100%);
    }

    /* Show mobile top bar */
    .floating-buttons {
        display: flex;
        position: fixed;
        top: 0;
        left: 0;
        right: 0;
        height: 58px;
        background: var(--sidebar-bg);
        border: none;
        border-bottom: 1px solid var(--sidebar-border);
        border-radius: 0;
        padding: 0 12px;
        justify-content: space-between;
        align-items: center;
        gap: 0;
        z-index: 250;
    }

    .floating-buttons .sidebar-open-btn,
    .floating-buttons .search-open-btn {
        display: flex;
        align-items: center;
        justify-content: center;
        border: none;
        background: transparent;
        padding: 8px;
        color: var(--fg-muted);
        cursor: pointer;
    }

    .sidebar-search-btn {
        display: none;
    }

    .topbar-title {
        display: block;
        font-size: 16px;
        font-weight: 600;
        color: var(--fg-color);
    }

    .content {
        margin-left: 0;
        padding: 74px 20px 20px;
    }

    .sidebar-overlay {
        display: block;
        pointer-events: none;
    }

    .sidebar-overlay.active {
        pointer-events: auto;
    }

    /* Show floating comments button on mobile */
    .open-comments-btn {
        display: flex;
        position: fixed;
        bottom: 16px;
        right: 16px;
        align-items: center;
        justify-content: center;
        width: 44px;
        height: 44px;
        padding: 0;
        background: var(--sidebar-bg);
        border: 1px solid var(--sidebar-border);
        border-radius: 22px;
        color: var(--fg-muted);
        cursor: pointer;
        z-index: 100;
        box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
    }

    .open-comments-btn svg {
        width: 20px;
        height: 20px;
    }

    .comment-count {
        position: absolute;
        top: -4px;
        right: -4px;
        min-width: 18px;
        height: 18px;
        padding: 0 5px;
        background: var(--sidebar-active);
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

    .comments-panel {
        top: 0;
        width: 100%;
    }

    .comment-btn {
        font-size: 12px;
        padding: 6px 10px;
    }
}

@media (min-width: 769px) and (max-width: 1200px) {
    :root {
        --sidebar-width: 240px;
        --content-padding: 30px;
    }
}

/* Search Palette */
.search-palette-overlay {
    position: fixed;
    inset: 0;
    background: rgba(0, 0, 0, 0.5);
    z-index: 300;
    opacity: 0;
    visibility: hidden;
    transition: opacity 0.15s ease, visibility 0.15s ease;
}

.search-palette-overlay.active {
    opacity: 1;
    visibility: visible;
}

.search-palette {
    position: fixed;
    top: 20%;
    left: 50%;
    transform: translateX(-50%) scale(0.95);
    width: 90%;
    max-width: 500px;
    background: var(--sidebar-bg);
    border: 1px solid var(--sidebar-border);
    border-radius: 8px;
    box-shadow: 0 8px 30px rgba(0, 0, 0, 0.3);
    z-index: 400;
    opacity: 0;
    visibility: hidden;
    transition: opacity 0.15s ease, visibility 0.15s ease, transform 0.15s ease;
}

.search-palette.active {
    opacity: 1;
    visibility: visible;
    transform: translateX(-50%) scale(1);
}

.search-palette-input {
    width: 100%;
    padding: 14px 16px;
    font-size: 16px;
    font-family: inherit;
    border: none;
    border-bottom: 1px solid var(--sidebar-border);
    border-radius: 8px 8px 0 0;
    background: transparent;
    color: var(--fg-color);
    outline: none;
    box-sizing: border-box;
}

.search-palette-input::placeholder {
    color: var(--fg-muted);
}

.search-palette-results {
    list-style: none;
    margin: 0;
    padding: 0;
    max-height: 300px;
    overflow-y: auto;
    scrollbar-width: thin;
    scrollbar-color: var(--sidebar-border) transparent;
}

.search-palette-results::-webkit-scrollbar {
    width: 8px;
}

.search-palette-results::-webkit-scrollbar-track {
    background: transparent;
}

.search-palette-results::-webkit-scrollbar-thumb {
    background-color: var(--sidebar-border);
    border-radius: 4px;
}

.search-palette-results::-webkit-scrollbar-thumb:hover {
    background-color: var(--fg-muted);
}

.search-palette-item {
    display: flex;
    flex-direction: column;
    padding: 10px 16px;
    cursor: pointer;
    border-left: 3px solid transparent;
    transition: background 0.1s ease;
}

.search-palette-item:hover {
    background: var(--sidebar-hover);
}

.search-palette-item.selected {
    background: var(--sidebar-active-bg);
    border-left-color: var(--sidebar-active);
}

.search-palette-item-name {
    font-size: 14px;
    font-weight: 500;
    color: var(--fg-color);
}

.search-palette-item.selected .search-palette-item-name {
    color: var(--sidebar-active);
}

.search-palette-item-path {
    font-size: 12px;
    color: var(--fg-muted);
    margin-top: 2px;
}

.search-palette-empty {
    padding: 20px 16px;
    text-align: center;
    color: var(--fg-muted);
    font-size: 14px;
}

.search-palette-hint {
    padding: 8px 16px;
    font-size: 11px;
    color: var(--fg-muted);
    border-top: 1px solid var(--sidebar-border);
    display: flex;
    gap: 12px;
    justify-content: center;
}

.search-palette-hint kbd {
    background: var(--sidebar-hover);
    padding: 2px 6px;
    border-radius: 3px;
    font-family: inherit;
    font-size: 11px;
}

/* Comment Button */
.comment-btn {
    position: absolute;
    display: none;
    align-items: center;
    gap: 6px;
    padding: 8px 12px;
    background: var(--sidebar-active);
    color: #ffffff;
    border: none;
    border-radius: 6px;
    font-size: 13px;
    font-family: inherit;
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
    top: var(--topbar-height);
    bottom: 0;
    width: var(--sidebar-width);
    background: var(--sidebar-bg);
    border-left: 1px solid var(--sidebar-border);
    display: flex;
    flex-direction: column;
    z-index: 200;
    transform: translateX(100%);
    transition: transform var(--transition-speed) ease;
}

.comments-panel.open {
    transform: translateX(0);
}

.comments-panel-header {
    display: flex;
    align-items: center;
    height: var(--panel-section-height);
    padding: 0 16px;
    flex-shrink: 0;
}

.comments-panel-header h2 {
    margin: 0;
    font-size: 12px;
    font-weight: 600;
    color: var(--fg-muted);
    text-transform: uppercase;
    letter-spacing: 0.5px;
}

/* Comments List */
.comments-list {
    flex: 1;
    overflow-y: auto;
    padding: 0;
    scrollbar-width: thin;
    scrollbar-color: var(--sidebar-border) transparent;
}

.comments-list::-webkit-scrollbar {
    width: 8px;
}

.comments-list::-webkit-scrollbar-track {
    background: transparent;
}

.comments-list::-webkit-scrollbar-thumb {
    background-color: var(--sidebar-border);
    border-radius: 4px;
}

/* Comment Entry */
.comment-entry {
    padding: 12px 16px;
    border-bottom: 1px solid var(--sidebar-border);
    cursor: pointer;
    transition: background 0.15s ease;
    position: relative;
}

.comment-entry:hover {
    background: var(--sidebar-hover);
}

.comment-entry.active {
    background: var(--sidebar-active-bg);
}

.comment-quote {
    margin: 0 0 8px 0;
    padding: 8px 12px;
    background: var(--sidebar-hover);
    border-left: 3px solid var(--sidebar-active);
    border-radius: 0 4px 4px 0;
    font-size: 13px;
    color: var(--fg-muted);
    font-style: italic;
    word-break: break-word;
}

.comment-text {
    font-size: 14px;
    color: var(--fg-color);
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
    color: var(--fg-muted);
    padding: 4px;
    border-radius: 4px;
    transition: all 0.15s ease;
}

.comment-copy-btn:hover,
.comment-edit-btn:hover {
    color: var(--sidebar-active);
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
    border-bottom: 1px solid var(--sidebar-border);
    background: var(--sidebar-bg);
}

.comment-input-quote {
    margin: 0 0 12px 0;
    padding: 8px 12px;
    background: var(--sidebar-hover);
    border-left: 3px solid var(--sidebar-active);
    border-radius: 0 4px 4px 0;
    font-size: 13px;
    color: var(--fg-muted);
    font-style: italic;
    word-break: break-word;
}

.comment-input-textarea {
    width: 100%;
    padding: 10px 12px;
    border: 1px solid var(--sidebar-border);
    border-radius: 6px;
    background: var(--bgColor-default, #ffffff);
    color: var(--fg-color);
    font-family: inherit;
    font-size: 14px;
    resize: vertical;
    min-height: 80px;
    box-sizing: border-box;
}

.comment-input-textarea:focus {
    outline: none;
    border-color: var(--sidebar-active);
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
    border: 1px solid var(--sidebar-border);
    color: var(--fg-color);
}

.comment-cancel-btn:hover {
    background: var(--sidebar-hover);
}

.comment-save-btn {
    display: inline-flex;
    align-items: center;
    gap: 6px;
    background: var(--sidebar-active);
    border: none;
    color: #ffffff;
}

.comment-save-btn:hover {
    filter: brightness(1.1);
}

/* Open Comments Button - hidden on desktop (shown in topbar), shown on mobile */
.open-comments-btn {
    display: none;
}

/* Comments Panel Footer */
.comments-panel-footer {
    display: flex;
    align-items: center;
    justify-content: center;
    height: var(--panel-section-height);
    padding: 0 16px;
    border-top: 1px solid var(--sidebar-border);
    flex-shrink: 0;
}

/* Copy Comments Button (in footer) */
.copy-comments-btn {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 6px;
    height: 36px;
    padding: 0 12px;
    background: transparent;
    border: none;
    border-radius: 6px;
    color: var(--fg-muted);
    font-size: 12px;
    font-weight: 500;
    font-family: inherit;
    cursor: pointer;
    transition: all 0.15s ease;
}

.copy-comments-btn:hover {
    color: var(--accent-color);
    background: var(--accent-bg);
}

.copy-comments-btn.copied {
    color: #1a7f37;
    background: rgba(26, 127, 55, 0.08);
}

.copy-comments-btn svg {
    width: 16px;
    height: 16px;
}

/* Empty state for comments */
.comments-empty {
    padding: 40px 16px;
    text-align: center;
    color: var(--fg-muted);
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
    background: var(--sidebar-hover);
    border-radius: 4px;
    font-family: inherit;
    font-size: 12px;
    font-weight: 500;
}

/* Keyboard Shortcuts Modal */
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
    background: var(--sidebar-bg);
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
    border-bottom: 1px solid var(--sidebar-border);
}

.shortcuts-modal-header h2 {
    margin: 0;
    font-size: 16px;
    font-weight: 600;
    color: var(--fg-color);
}

.shortcuts-modal-close {
    display: flex;
    align-items: center;
    justify-content: center;
    background: none;
    border: none;
    cursor: pointer;
    color: var(--fg-muted);
    padding: 4px;
    border-radius: 4px;
    transition: all 0.15s ease;
}

.shortcuts-modal-close:hover {
    color: var(--fg-color);
    background: var(--sidebar-hover);
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
    border-bottom: 1px solid var(--sidebar-hover);
}

.shortcut-action {
    font-size: 14px;
    color: var(--fg-color);
}

.shortcut-keys {
    display: flex;
    gap: 4px;
}

.shortcut-keys kbd {
    display: inline-block;
    padding: 4px 8px;
    background: var(--sidebar-hover);
    border: 1px solid var(--sidebar-border);
    border-radius: 6px;
    font-family: inherit;
    font-size: 12px;
    font-weight: 500;
    color: var(--fg-color);
}

@media (prefers-color-scheme: dark) {
    .comment-highlight {
        background-color: rgba(255, 212, 59, 0.25);
    }

    .comment-highlight:hover,
    .comment-highlight.active {
        background-color: rgba(255, 212, 59, 0.45);
    }

    .comment-input-textarea {
        background: var(--bgColor-default, #0d1117);
    }

    .comment-input-textarea:focus {
        box-shadow: 0 0 0 3px rgba(88, 166, 255, 0.1);
    }

    .comment-copy-btn:hover,
    .comment-edit-btn:hover {
        background: rgba(88, 166, 255, 0.1);
    }

    .comment-copy-btn.copied {
        color: #3fb950;
    }

    .comment-delete-btn:hover {
        color: #f85149;
        background: rgba(248, 81, 73, 0.1);
    }

    .open-comments-btn {
        box-shadow: 0 2px 8px rgba(0, 0, 0, 0.3);
    }

    .copy-comments-btn.copied {
        color: #3fb950;
        background: rgba(63, 185, 80, 0.1);
    }

    .shortcuts-modal-overlay {
        background: rgba(0, 0, 0, 0.7);
    }

    .shortcuts-modal {
        box-shadow: 0 8px 30px rgba(0, 0, 0, 0.4);
    }
}

`

const sidebarJS = `
(function() {
    'use strict';

    var sidebar = document.querySelector('.sidebar');
    var openBtn = document.querySelector('.sidebar-open-btn');
    var searchBtn = document.querySelector('.search-open-btn');
    var overlay = document.querySelector('.sidebar-overlay');
    var fileLinks = document.querySelectorAll('.file-tree a[data-file]');
    var contentSections = document.querySelectorAll('.content-section');
    var directories = document.querySelectorAll('.file-tree .directory > span');

    // Desktop topbar elements
    var topbarSidebarBtn = document.querySelector('.topbar-sidebar-btn');
    var topbarSearchBtn = document.querySelector('.topbar-search-btn');
    var topbarCommentBtn = document.querySelector('.topbar-comment-btn');
    var topbarCommentCount = document.querySelector('.topbar-comment-count');
    var topbarHelpBtn = document.querySelector('.topbar-help-btn');

    function isMobile() {
        return window.innerWidth <= 768;
    }

    var isInitialLoad = true;

    function showFile(fileId) {
        for (var i = 0; i < contentSections.length; i++) {
            if (contentSections[i].id === fileId) {
                contentSections[i].classList.add('active');
            } else {
                contentSections[i].classList.remove('active');
            }
        }

        for (var i = 0; i < fileLinks.length; i++) {
            if (fileLinks[i].dataset.file === fileId) {
                fileLinks[i].classList.add('active');
            } else {
                fileLinks[i].classList.remove('active');
            }
        }

        if (isMobile()) {
            closeSidebar();
        }

        history.replaceState(null, '', '#' + fileId);

        if (isInitialLoad) {
            isInitialLoad = false;
            requestAnimationFrame(function() {
                window.scrollTo(0, 0);
            });
        }
    }

    function openSidebar() {
        if (isMobile()) {
            sidebar.classList.add('open');
            overlay.classList.add('active');
            document.body.style.overflow = 'hidden';
        } else {
            sidebar.classList.remove('collapsed');
            topbarSidebarBtn.classList.add('active');
        }
    }

    function closeSidebar() {
        if (isMobile()) {
            sidebar.classList.remove('open');
            overlay.classList.remove('active');
            document.body.style.overflow = '';
        } else {
            sidebar.classList.add('collapsed');
            topbarSidebarBtn.classList.remove('active');
        }
    }

    function toggleSidebar() {
        if (isMobile()) {
            if (sidebar.classList.contains('open')) {
                closeSidebar();
            } else {
                openSidebar();
            }
        } else {
            if (sidebar.classList.contains('collapsed')) {
                openSidebar();
            } else {
                closeSidebar();
            }
        }
    }

    function toggleDirectory(e) {
        e.target.parentElement.classList.toggle('open');
    }

    for (var i = 0; i < fileLinks.length; i++) {
        fileLinks[i].addEventListener('click', function(e) {
            e.preventDefault();
            showFile(this.dataset.file);
        });
    }

    // Handle clicks on internal links within content (rewritten .md links)
    document.addEventListener('click', function(e) {
        var link = e.target.closest('a[href^="#"]');
        if (!link) return;

        // Skip sidebar links (they have data-file attribute)
        if (link.dataset.file) return;

        var fileId = link.getAttribute('href').slice(1);
        // Check if this ID matches a file section
        for (var i = 0; i < fileLinks.length; i++) {
            if (fileLinks[i].dataset.file === fileId) {
                e.preventDefault();
                showFile(fileId);
                return;
            }
        }
        // If not a file section, let the browser handle it (regular anchor)
    });

    openBtn.addEventListener('click', openSidebar);
    overlay.addEventListener('click', closeSidebar);

    // Desktop topbar event listeners
    topbarSidebarBtn.addEventListener('click', toggleSidebar);
    topbarSearchBtn.addEventListener('click', function() {
        openSearchPalette();
    });
    topbarHelpBtn.addEventListener('click', function() {
        openShortcutsModal();
    });

    for (var i = 0; i < directories.length; i++) {
        directories[i].addEventListener('click', toggleDirectory);
    }

    // Initialize sidebar state on desktop (open by default)
    if (!isMobile() && !sidebar.classList.contains('collapsed')) {
        topbarSidebarBtn.classList.add('active');
    }

    // Set tooltip with keyboard shortcut based on platform
    var isMac = navigator.platform.toUpperCase().indexOf('MAC') >= 0;
    var shortcut = isMac ? '⌘B' : 'Ctrl+B';
    var searchShortcut = isMac ? '⌘K' : 'Ctrl+K';
    openBtn.title = 'Open sidebar (' + shortcut + ')';
    searchBtn.title = 'Search files (' + searchShortcut + ')';

    // Initialize file display
    if (window.location.hash) {
        var fileId = window.location.hash.slice(1);
        var exists = false;
        for (var i = 0; i < fileLinks.length; i++) {
            if (fileLinks[i].dataset.file === fileId) {
                exists = true;
                break;
            }
        }
        if (exists) {
            showFile(fileId);
        } else if (fileLinks.length > 0) {
            showFile(fileLinks[0].dataset.file);
        }
    } else if (fileLinks.length > 0) {
        showFile(fileLinks[0].dataset.file);
    }

    // Search Palette
    var searchOverlay = document.getElementById('search-overlay');
    var searchPalette = document.getElementById('search-palette');
    var searchInput = document.getElementById('search-input');
    var searchResults = document.getElementById('search-results');
    var paletteOpen = false;
    var selectedIndex = 0;
    var filteredFiles = [];

    // Build file list from DOM
    var allFiles = [];
    for (var i = 0; i < fileLinks.length; i++) {
        var link = fileLinks[i];
        var name = link.textContent;
        var id = link.dataset.file;
        // Get path from parent directory structure
        var path = '';
        var li = link.parentElement;
        var parent = li.parentElement;
        while (parent && parent.classList.contains('directory') === false) {
            parent = parent.parentElement;
        }
        if (parent && parent.classList.contains('directory')) {
            var dirSpan = parent.querySelector(':scope > span');
            if (dirSpan) {
                path = dirSpan.textContent + '/';
            }
        }
        allFiles.push({ id: id, name: name, path: path + name });
    }

    function fuzzyMatch(query, text) {
        query = query.toLowerCase();
        text = text.toLowerCase();
        var qi = 0;
        for (var ti = 0; ti < text.length && qi < query.length; ti++) {
            if (text[ti] === query[qi]) qi++;
        }
        return qi === query.length;
    }

    function renderResults() {
        searchResults.innerHTML = '';
        if (filteredFiles.length === 0) {
            searchResults.innerHTML = '<li class="search-palette-empty">No files found</li>';
            return;
        }
        for (var i = 0; i < filteredFiles.length; i++) {
            var file = filteredFiles[i];
            var li = document.createElement('li');
            li.className = 'search-palette-item' + (i === selectedIndex ? ' selected' : '');
            li.dataset.index = i;
            li.innerHTML = '<span class="search-palette-item-name">' + file.name + '</span>' +
                           '<span class="search-palette-item-path">' + file.path + '</span>';
            searchResults.appendChild(li);
        }
    }

    function filterFiles(query) {
        if (!query) {
            filteredFiles = allFiles.slice();
        } else {
            filteredFiles = [];
            for (var i = 0; i < allFiles.length; i++) {
                if (fuzzyMatch(query, allFiles[i].path)) {
                    filteredFiles.push(allFiles[i]);
                }
            }
        }
        selectedIndex = 0;
        renderResults();
    }

    function openSearchPalette() {
        paletteOpen = true;
        searchOverlay.classList.add('active');
        searchPalette.classList.add('active');
        searchInput.value = '';
        filterFiles('');
        // Focus after element becomes visible
        // Use double-rAF to ensure browser has painted, then focus
        requestAnimationFrame(function() {
            requestAnimationFrame(function() {
                searchInput.focus();
            });
        });
    }

    function closeSearchPalette() {
        paletteOpen = false;
        searchOverlay.classList.remove('active');
        searchPalette.classList.remove('active');
        searchInput.value = '';
    }

    function selectCurrentFile() {
        if (filteredFiles.length > 0 && filteredFiles[selectedIndex]) {
            showFile(filteredFiles[selectedIndex].id);
            closeSearchPalette();
        }
    }

    function moveSelection(dir) {
        if (filteredFiles.length === 0) return;
        selectedIndex += dir;
        if (selectedIndex < 0) selectedIndex = filteredFiles.length - 1;
        if (selectedIndex >= filteredFiles.length) selectedIndex = 0;
        renderResults();
        // Scroll selected item into view
        var selected = searchResults.querySelector('.selected');
        if (selected) {
            selected.scrollIntoView({ block: 'nearest' });
        }
    }

    searchInput.addEventListener('input', function() {
        filterFiles(this.value);
    });

    searchResults.addEventListener('click', function(e) {
        var item = e.target.closest('.search-palette-item');
        if (item) {
            selectedIndex = parseInt(item.dataset.index, 10);
            selectCurrentFile();
        }
    });

    searchOverlay.addEventListener('click', closeSearchPalette);
    searchBtn.addEventListener('click', openSearchPalette);

    // Keyboard shortcuts
    document.addEventListener('keydown', function(e) {
        // Cmd+K or Ctrl+K to open search palette
        if (e.key === 'k' && (e.metaKey || e.ctrlKey)) {
            e.preventDefault();
            if (paletteOpen) {
                closeSearchPalette();
            } else {
                openSearchPalette();
            }
            return;
        }

        // When palette is open
        if (paletteOpen) {
            if (e.key === 'Escape') {
                e.preventDefault();
                closeSearchPalette();
                return;
            }
            if (e.key === 'Enter') {
                e.preventDefault();
                selectCurrentFile();
                return;
            }
            if (e.key === 'ArrowDown' || (e.key === 'n' && e.ctrlKey)) {
                e.preventDefault();
                moveSelection(1);
                return;
            }
            if (e.key === 'ArrowUp' || (e.key === 'p' && e.ctrlKey)) {
                e.preventDefault();
                moveSelection(-1);
                return;
            }
            return;
        }

        // Escape to close sidebar on mobile
        if (e.key === 'Escape') {
            if (isMobile() && sidebar.classList.contains('open')) {
                closeSidebar();
            }
        }
        // Cmd+B (Mac) or Ctrl+B (Windows/Linux) to toggle sidebar on desktop
        if (e.key === 'b' && (e.metaKey || e.ctrlKey)) {
            if (!isMobile()) {
                e.preventDefault();
                if (sidebar.classList.contains('collapsed')) {
                    openSidebar();
                } else {
                    closeSidebar();
                }
            }
        }
    });

    // Copy button for code blocks
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

    // ========== Comments Feature ==========
    var deleteIcon = '<svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="3 6 5 6 21 6"></polyline><path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"></path></svg>';
    var commentCopyIcon = '<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg>';
    var commentCheckIcon = '<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 6 9 17 4 12"></polyline></svg>';

    // Multi-file: Store comments per file
    var STORAGE_KEY = 'mdp-comments-multi';
    var commentsByFile = {};
    var commentCounter = 0;
    var currentRange = null;
    var currentSelectedText = '';

    // DOM Elements for comments
    var commentBtn = document.querySelector('.comment-btn');
    var commentsPanel = document.querySelector('.comments-panel');
    var commentsList = document.querySelector('.comments-list');
    var commentInputForm = document.querySelector('.comment-input-form');
    var commentInputQuote = document.querySelector('.comment-input-quote');
    var commentInputTextarea = document.querySelector('.comment-input-textarea');
    var copyCommentsBtn = document.querySelector('.copy-comments-btn');
    var openCommentsBtn = document.querySelector('.open-comments-btn');
    var commentCountEl = document.querySelector('.comment-count');
    var commentCancelBtn = document.querySelector('.comment-cancel-btn');
    var commentSaveBtn = document.querySelector('.comment-save-btn');
    var commentsEmpty = document.querySelector('.comments-empty');
    var shortcutsOverlay = document.querySelector('.shortcuts-modal-overlay');
    var shortcutsModal = document.querySelector('.shortcuts-modal');
    var shortcutsCloseBtn = document.querySelector('.shortcuts-modal-close');

    function getCurrentFileId() {
        var activeSection = document.querySelector('.content-section.active');
        return activeSection ? activeSection.id : null;
    }

    function getComments() {
        var fileId = getCurrentFileId();
        return fileId ? (commentsByFile[fileId] || []) : [];
    }

    function setComments(comments) {
        var fileId = getCurrentFileId();
        if (fileId) {
            commentsByFile[fileId] = comments;
            saveCommentsToStorage();
        }
    }

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

    // SessionStorage persistence functions
    function saveCommentsToStorage() {
        try {
            sessionStorage.setItem(STORAGE_KEY, JSON.stringify(commentsByFile));
        } catch (e) {
            // Silently fail if storage is full or unavailable
        }
    }

    function loadCommentsFromStorage() {
        try {
            var stored = sessionStorage.getItem(STORAGE_KEY);
            if (stored) {
                commentsByFile = JSON.parse(stored);
                // Update comment counter to avoid ID collisions
                var maxId = 0;
                Object.keys(commentsByFile).forEach(function(fileId) {
                    commentsByFile[fileId].forEach(function(c) {
                        var num = parseInt(c.id.replace('comment-', ''), 10);
                        if (num > maxId) maxId = num;
                    });
                });
                commentCounter = maxId;
            }
        } catch (e) {
            commentsByFile = {};
        }
    }

    function findTextInDocument(searchText, container) {
        if (!container) return null;
        var treeWalker = document.createTreeWalker(container, NodeFilter.SHOW_TEXT, null);
        var fullText = '';
        var textNodes = [];
        var node;

        while (node = treeWalker.nextNode()) {
            textNodes.push({ node: node, start: fullText.length });
            fullText += node.textContent;
        }

        var index = fullText.indexOf(searchText);
        if (index === -1) return null;

        var endIndex = index + searchText.length;
        var startNode = null, endNode = null;
        var startOffset = 0, endOffset = 0;

        for (var i = 0; i < textNodes.length; i++) {
            var info = textNodes[i];
            var nodeEnd = info.start + info.node.textContent.length;

            if (!startNode && index >= info.start && index < nodeEnd) {
                startNode = info.node;
                startOffset = index - info.start;
            }
            if (endIndex > info.start && endIndex <= nodeEnd) {
                endNode = info.node;
                endOffset = endIndex - info.start;
                break;
            }
        }

        if (!startNode || !endNode) return null;

        var range = document.createRange();
        range.setStart(startNode, startOffset);
        range.setEnd(endNode, endOffset);
        return range;
    }

    function restoreHighlightsForFile(fileId) {
        var fileComments = commentsByFile[fileId] || [];
        var container = document.querySelector('#' + fileId + ' .markdown-body');
        if (!container) return;

        fileComments.forEach(function(comment) {
            var range = findTextInDocument(comment.selectedText, container);
            if (range) {
                highlightRange(range, comment.id);
            }
        });
    }

    function isValidSelection(selection) {
        if (!selection || selection.isCollapsed) return false;
        var range = selection.getRangeAt(0);
        var container = range.commonAncestorContainer;
        // Must be within active content section
        var activeSection = document.querySelector('.content-section.active .markdown-body');
        if (!activeSection || !activeSection.contains(container)) return false;
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

    function openCommentsPanel() {
        commentsPanel.classList.add('open');
        topbarCommentBtn.classList.add('active');
    }

    function closeCommentsPanel() {
        commentsPanel.classList.remove('open');
        topbarCommentBtn.classList.remove('active');
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

    function renderCommentsForCurrentFile() {
        // Clear existing entries (except input form)
        var entries = commentsList.querySelectorAll('.comment-entry');
        entries.forEach(function(entry) { entry.remove(); });

        var comments = getComments();
        comments.forEach(function(comment) {
            addCommentEntryToDOM(comment);
        });

        updateCommentsUI();
    }

    var smallCopyIcon = '<svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg>';
    var smallCheckIcon = '<svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 6 9 17 4 12"></polyline></svg>';
    var editIcon = '<svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M17 3a2.85 2.83 0 1 1 4 4L7.5 20.5 2 22l1.5-5.5Z"></path><path d="m15 5 4 4"></path></svg>';
    var editingCommentId = null;

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

    function editComment(commentId) {
        var comments = getComments();
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

    function addCommentEntryToDOM(comment) {
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
        var comments = getComments();
        comments = comments.filter(function(c) { return c.id !== commentId; });
        setComments(comments);

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

    function copyCommentsAsMarkdown() {
        var comments = getComments();
        if (comments.length === 0) return;

        var markdown = comments.map(function(c) {
            var quotedText = c.selectedText.split('\n').map(function(line) {
                return '> ' + line;
            }).join('\n');
            return quotedText + '\n\nComment: ' + c.commentText;
        }).join('\n\n---\n\n');

        navigator.clipboard.writeText(markdown).then(function() {
            copyCommentsBtn.innerHTML = commentCheckIcon + '<span>Copied!</span>';
            copyCommentsBtn.classList.add('copied');
            setTimeout(function() {
                copyCommentsBtn.innerHTML = commentCopyIcon + '<span>Copy Comments</span>';
                copyCommentsBtn.classList.remove('copied');
            }, 2000);
        });
    }

    function updateCommentsUI() {
        var comments = getComments();
        var count = comments.length;
        // Update mobile comment count
        commentCountEl.textContent = count;
        // Update topbar comment count
        topbarCommentCount.textContent = count;
        if (count > 0) {
            commentCountEl.classList.add('visible');
            topbarCommentCount.classList.add('visible');
            commentsEmpty.style.display = 'none';
        } else {
            commentCountEl.classList.remove('visible');
            topbarCommentCount.classList.remove('visible');
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

    // Shortcuts modal event listeners
    shortcutsCloseBtn.addEventListener('click', closeShortcutsModal);
    shortcutsOverlay.addEventListener('click', function(e) {
        if (e.target === shortcutsOverlay) {
            closeShortcutsModal();
        }
    });

    // Override showFile to update comments display and restore highlights
    var originalShowFile = showFile;
    var restoredFiles = {};
    showFile = function(fileId) {
        originalShowFile(fileId);
        // Restore highlights for this file if not already done
        if (!restoredFiles[fileId] && commentsByFile[fileId] && commentsByFile[fileId].length > 0) {
            restoreHighlightsForFile(fileId);
            restoredFiles[fileId] = true;
        }
        // Re-render comments for the new file
        renderCommentsForCurrentFile();
    };

    // Event listeners for comments
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
            var comments = getComments();
            var comment = comments.find(function(c) { return c.id === editingCommentId; });
            if (comment) {
                comment.commentText = commentText;
                setComments(comments);
                // Update the entry in the DOM
                var entry = commentsList.querySelector('.comment-entry[data-comment-id="' + editingCommentId + '"]');
                if (entry) {
                    entry.querySelector('.comment-text').textContent = commentText;
                }
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

        var comments = getComments();
        comments.push(comment);
        setComments(comments);

        addCommentEntryToDOM(comment);
        hideInputForm();
        updateCommentsUI();

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

    copyCommentsBtn.addEventListener('click', copyCommentsAsMarkdown);
    openCommentsBtn.addEventListener('click', function() {
        if (commentsPanel.classList.contains('open')) {
            closeCommentsPanel();
        } else {
            openCommentsPanel();
        }
    });

    // Desktop topbar comment button
    topbarCommentBtn.addEventListener('click', function() {
        if (commentsPanel.classList.contains('open')) {
            closeCommentsPanel();
            topbarCommentBtn.classList.remove('active');
        } else {
            openCommentsPanel();
            topbarCommentBtn.classList.add('active');
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

    // Keyboard shortcuts for comments (use capture phase to intercept before browser)
    document.addEventListener('keydown', function(e) {
        // Don't trigger shortcuts when typing in an input/textarea
        var isTyping = e.target.tagName === 'INPUT' || e.target.tagName === 'TEXTAREA';

        // Escape to close modals/panels (prioritize shortcuts modal > comments panel)
        if (e.key === 'Escape') {
            if (shortcutsModal.classList.contains('active')) {
                closeShortcutsModal();
                e.preventDefault();
                return;
            }
            if (commentsPanel.classList.contains('open')) {
                closeCommentsPanel();
                e.preventDefault();
                return;
            }
        }

        // '?' (Shift+/) to show keyboard shortcuts
        if (e.key === '?' && !isTyping && !paletteOpen) {
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
            // Don't trigger if search palette is open
            if (paletteOpen) return;
            // Only trigger if comment button is visible (has selection)
            if (commentBtn.classList.contains('visible')) {
                e.preventDefault();
                commentBtn.click();
            }
        }
    }, true);

    // Initialize comments - load from sessionStorage
    loadCommentsFromStorage();

    // Restore highlights for initial file
    var initialFileId = getCurrentFileId();
    if (initialFileId && commentsByFile[initialFileId] && commentsByFile[initialFileId].length > 0) {
        restoreHighlightsForFile(initialFileId);
        restoredFiles[initialFileId] = true;
        renderCommentsForCurrentFile();
    }

    updateCommentsUI();
})();
`

const multiFileLiveReloadScript = `
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

// GenerateMulti creates an HTML document with sidebar navigation for multiple files.
func GenerateMulti(title string, tree *filetree.TreeNode, files []filetree.FileEntry) string {
	sidebarHTML := generateSidebarHTML(tree)
	contentHTML := generateContentSections(files)

	return fmt.Sprintf(multiFileTemplate,
		html.EscapeString(title),
		githubMarkdownCSS,
		chromaCSS,
		sidebarCSS,
		sidebarHTML,
		contentHTML,
		sidebarJS,
		"", // No live reload script
	)
}

// GenerateMultiWithLiveReload creates an HTML document with sidebar navigation and live reload.
func GenerateMultiWithLiveReload(title string, tree *filetree.TreeNode, files []filetree.FileEntry, port int) string {
	sidebarHTML := generateSidebarHTML(tree)
	contentHTML := generateContentSections(files)
	liveReloadScript := fmt.Sprintf(multiFileLiveReloadScript, port)

	return fmt.Sprintf(multiFileTemplate,
		html.EscapeString(title),
		githubMarkdownCSS,
		chromaCSS,
		sidebarCSS,
		sidebarHTML,
		contentHTML,
		sidebarJS,
		liveReloadScript,
	)
}

// generateSidebarHTML creates the file tree HTML structure.
func generateSidebarHTML(tree *filetree.TreeNode) string {
	var buf strings.Builder
	buf.WriteString("<ul>")
	for _, child := range tree.Children {
		renderTreeNode(&buf, child)
	}
	buf.WriteString("</ul>")
	return buf.String()
}

// renderTreeNode recursively renders a tree node to HTML.
func renderTreeNode(buf *strings.Builder, node *filetree.TreeNode) {
	if node.IsDir {
		buf.WriteString(`<li class="directory open">`)
		buf.WriteString(fmt.Sprintf("<span>%s</span>", html.EscapeString(node.Name)))
		buf.WriteString("<ul>")
		for _, child := range node.Children {
			renderTreeNode(buf, child)
		}
		buf.WriteString("</ul></li>")
	} else if node.File != nil {
		buf.WriteString("<li>")
		buf.WriteString(fmt.Sprintf(
			`<a href="#%s" data-file="%s">%s</a>`,
			html.EscapeString(node.File.ID),
			html.EscapeString(node.File.ID),
			html.EscapeString(node.File.Name),
		))
		buf.WriteString("</li>")
	}
}

// generateContentSections creates the content divs for each file.
func generateContentSections(files []filetree.FileEntry) string {
	var buf strings.Builder
	for i, f := range files {
		class := "content-section"
		if i == 0 {
			class = "content-section active"
		}
		buf.WriteString(fmt.Sprintf(
			`<section id="%s" class="%s"><article class="markdown-body">%s</article></section>`,
			html.EscapeString(f.ID),
			class,
			f.Content,
		))
	}
	return buf.String()
}
