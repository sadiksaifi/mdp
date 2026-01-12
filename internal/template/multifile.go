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
</head>
<body>
    <button class="sidebar-open-btn" aria-label="Open sidebar" title="Open sidebar">
		<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="icon icon-tabler icons-tabler-outline icon-tabler-layout-sidebar-right"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M4 6a2 2 0 0 1 2 -2h12a2 2 0 0 1 2 2v12a2 2 0 0 1 -2 2h-12a2 2 0 0 1 -2 -2l0 -12" /><path d="M15 4l0 16" /></svg>
    </button>

    <div class="sidebar-overlay"></div>

    <aside class="sidebar" role="navigation" aria-label="File navigation">
        <div class="sidebar-header">
            <h2>Files</h2>
            <button class="sidebar-close-btn" aria-label="Close sidebar" title="Close sidebar">
				<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="icon icon-tabler icons-tabler-outline icon-tabler-layout-sidebar"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M4 6a2 2 0 0 1 2 -2h12a2 2 0 0 1 2 2v12a2 2 0 0 1 -2 2h-12a2 2 0 0 1 -2 -2l0 -12" /><path d="M9 4l0 16" /></svg>
            </button>
        </div>
        <nav class="file-tree">
            %s
        </nav>
    </aside>

    <main class="content">
        %s
    </main>

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

body {
    margin: 0;
    padding: 0;
    max-width: none;
    display: flex;
    min-height: 100vh;
    color: var(--fg-color);
    font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", "Noto Sans", Helvetica, Arial, sans-serif;
}

.sidebar {
    position: fixed;
    left: 0;
    top: 0;
    bottom: 0;
    width: var(--sidebar-width);
    background: var(--sidebar-bg);
    border-right: 1px solid var(--sidebar-border);
    overflow-y: auto;
    z-index: 200;
    transition: transform var(--transition-speed) ease, opacity var(--transition-speed) ease;
}

.sidebar.collapsed {
    transform: translateX(-100%);
}

.sidebar-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 16px;
    border-bottom: 1px solid var(--sidebar-border);
    position: sticky;
    top: 0;
    background: var(--sidebar-bg);
    z-index: 10;
}

.sidebar-header h2 {
    margin: 0;
    font-size: 14px;
    font-weight: 600;
    color: var(--fg-muted);
    text-transform: uppercase;
    letter-spacing: 0.5px;
}

.sidebar-close-btn {
    display: flex;
    align-items: center;
    justify-content: center;
    background: none;
    border: none;
    cursor: pointer;
    color: var(--fg-muted);
    padding: 4px;
    border-radius: 4px;
    transition: all 0.2s ease;
}

.sidebar-close-btn:hover {
    color: var(--fg-color);
    background: var(--sidebar-hover);
}

.sidebar-open-btn {
    position: fixed;
    top: 16px;
    left: 16px;
    z-index: 100;
    display: none;
    align-items: center;
    justify-content: center;
    background: var(--sidebar-bg);
    border: 1px solid var(--sidebar-border);
    border-radius: 6px;
    padding: 8px;
    cursor: pointer;
    color: var(--fg-muted);
    transition: all 0.2s ease;
}

.sidebar-open-btn:hover {
    color: var(--fg-color);
    background: var(--sidebar-hover);
}

.sidebar.collapsed ~ .sidebar-open-btn,
body:has(.sidebar.collapsed) .sidebar-open-btn {
    display: flex;
}

.file-tree {
    padding: 8px 0;
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
    margin-left: var(--sidebar-width);
    margin-right: auto;
    padding: var(--content-padding);
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
    z-index: 50;
    opacity: 0;
    transition: opacity var(--transition-speed) ease;
}

.sidebar-overlay.active {
    opacity: 1;
}

@media (max-width: 768px) {
    .sidebar {
        transform: translateX(-100%);
    }

    .sidebar.open {
        transform: translateX(0);
    }

    .sidebar.collapsed {
        transform: translateX(-100%);
    }

    .sidebar-open-btn {
        display: flex;
    }

    .content {
        margin-left: 0;
        padding: 70px 20px 20px;
    }

    .sidebar-overlay {
        display: block;
        pointer-events: none;
    }

    .sidebar-overlay.active {
        pointer-events: auto;
    }
}

@media (min-width: 769px) and (max-width: 1200px) {
    :root {
        --sidebar-width: 240px;
        --content-padding: 30px;
    }
}
`

const sidebarJS = `
(function() {
    'use strict';

    var sidebar = document.querySelector('.sidebar');
    var openBtn = document.querySelector('.sidebar-open-btn');
    var closeBtn = document.querySelector('.sidebar-close-btn');
    var overlay = document.querySelector('.sidebar-overlay');
    var fileLinks = document.querySelectorAll('.file-tree a[data-file]');
    var contentSections = document.querySelectorAll('.content-section');
    var directories = document.querySelectorAll('.file-tree .directory > span');

    function isMobile() {
        return window.innerWidth <= 768;
    }

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
    }

    function openSidebar() {
        if (isMobile()) {
            sidebar.classList.add('open');
            overlay.classList.add('active');
        } else {
            sidebar.classList.remove('collapsed');
            localStorage.setItem('mdp-sidebar-collapsed', 'false');
        }
    }

    function closeSidebar() {
        if (isMobile()) {
            sidebar.classList.remove('open');
            overlay.classList.remove('active');
        } else {
            sidebar.classList.add('collapsed');
            localStorage.setItem('mdp-sidebar-collapsed', 'true');
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

    openBtn.addEventListener('click', openSidebar);
    closeBtn.addEventListener('click', closeSidebar);
    overlay.addEventListener('click', closeSidebar);

    for (var i = 0; i < directories.length; i++) {
        directories[i].addEventListener('click', toggleDirectory);
    }

    document.addEventListener('keydown', function(e) {
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

    // Restore collapsed state from localStorage (desktop only)
    if (window.innerWidth > 768 && localStorage.getItem('mdp-sidebar-collapsed') === 'true') {
        sidebar.classList.add('collapsed');
    }

    // Set tooltip with keyboard shortcut based on platform
    var isMac = navigator.platform.toUpperCase().indexOf('MAC') >= 0;
    var shortcut = isMac ? '⌘B' : 'Ctrl+B';
    openBtn.title = 'Open sidebar (' + shortcut + ')';
    closeBtn.title = 'Close sidebar (' + shortcut + ')';

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
	for _, f := range files {
		buf.WriteString(fmt.Sprintf(
			`<section id="%s" class="content-section"><article class="markdown-body">%s</article></section>`,
			html.EscapeString(f.ID),
			f.Content,
		))
	}
	return buf.String()
}
