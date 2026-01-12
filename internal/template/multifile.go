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
    <div class="floating-buttons">
        <button class="sidebar-open-btn" aria-label="Open sidebar" title="Open sidebar">
            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M4 6a2 2 0 0 1 2 -2h12a2 2 0 0 1 2 2v12a2 2 0 0 1 -2 2h-12a2 2 0 0 1 -2 -2l0 -12" /><path d="M15 4l0 16" /></svg>
        </button>
        <span class="topbar-title">Preview</span>
        <button class="search-open-btn" aria-label="Search files" title="Search files">
            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M10 10m-7 0a7 7 0 1 0 14 0a7 7 0 1 0 -14 0" /><path d="M21 21l-6 -6" /></svg>
        </button>
    </div>

    <div class="sidebar-overlay"></div>

    <aside class="sidebar" role="navigation" aria-label="File navigation">
        <div class="sidebar-header">
            <h2>Files</h2>
            <div class="sidebar-header-buttons">
                <button class="sidebar-search-btn" aria-label="Search files" title="Search files">
                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M10 10m-7 0a7 7 0 1 0 14 0a7 7 0 1 0 -14 0" /><path d="M21 21l-6 -6" /></svg>
                </button>
                <button class="sidebar-close-btn" aria-label="Close sidebar" title="Close sidebar">
                    <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M4 6a2 2 0 0 1 2 -2h12a2 2 0 0 1 2 2v12a2 2 0 0 1 -2 2h-12a2 2 0 0 1 -2 -2l0 -12" /><path d="M9 4l0 16" /></svg>
                </button>
            </div>
        </div>
        <nav class="file-tree">
            %s
        </nav>
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
    justify-content: space-between;
    align-items: center;
    padding: 16px;
    border-bottom: 1px solid var(--sidebar-border);
    background: var(--sidebar-bg);
    z-index: 10;
    flex-shrink: 0;
}

.sidebar-header h2 {
    margin: 0;
    font-size: 14px;
    font-weight: 600;
    color: var(--fg-muted);
    text-transform: uppercase;
    letter-spacing: 0.5px;
}

.sidebar-header-buttons {
    display: flex;
    gap: 4px;
}

.sidebar-close-btn,
.sidebar-search-btn {
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

.sidebar-close-btn:hover,
.sidebar-search-btn:hover {
    color: var(--fg-color);
    background: var(--sidebar-hover);
}

.floating-buttons {
    position: fixed;
    top: 16px;
    left: 16px;
    z-index: 100;
    display: none;
    flex-direction: row;
    gap: 8px;
}

.sidebar.collapsed ~ .floating-buttons,
body:has(.sidebar.collapsed) .floating-buttons {
    display: flex;
}

.sidebar-open-btn,
.search-open-btn {
    display: flex;
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

.sidebar-open-btn:hover,
.search-open-btn:hover {
    color: var(--fg-color);
    background: var(--sidebar-hover);
}

.topbar-title {
    display: none;
}

.file-tree {
    padding: 8px 0;
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
    z-index: 150;
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

    .floating-buttons {
        display: flex;
        position: fixed;
        top: 0;
        left: 0;
        right: 0;
        height: 58px;
        background: var(--sidebar-bg);
        border-bottom: 1px solid var(--sidebar-border);
        padding: 0 12px;
        justify-content: space-between;
        align-items: center;
        gap: 0;
    }

    .floating-buttons .sidebar-open-btn,
    .floating-buttons .search-open-btn {
        border: none;
        background: transparent;
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
`

const sidebarJS = `
(function() {
    'use strict';

    var sidebar = document.querySelector('.sidebar');
    var openBtn = document.querySelector('.sidebar-open-btn');
    var closeBtn = document.querySelector('.sidebar-close-btn');
    var searchBtn = document.querySelector('.search-open-btn');
    var sidebarSearchBtn = document.querySelector('.sidebar-search-btn');
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
            document.body.style.overflow = 'hidden';
        } else {
            sidebar.classList.remove('collapsed');
            localStorage.setItem('mdp-sidebar-collapsed', 'false');
        }
    }

    function closeSidebar() {
        if (isMobile()) {
            sidebar.classList.remove('open');
            overlay.classList.remove('active');
            document.body.style.overflow = '';
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

    // Restore collapsed state from localStorage (desktop only)
    if (window.innerWidth > 768 && localStorage.getItem('mdp-sidebar-collapsed') === 'true') {
        sidebar.classList.add('collapsed');
    }

    // Set tooltip with keyboard shortcut based on platform
    var isMac = navigator.platform.toUpperCase().indexOf('MAC') >= 0;
    var shortcut = isMac ? '⌘B' : 'Ctrl+B';
    var searchShortcut = isMac ? '⌘K' : 'Ctrl+K';
    openBtn.title = 'Open sidebar (' + shortcut + ')';
    closeBtn.title = 'Close sidebar (' + shortcut + ')';
    searchBtn.title = 'Search files (' + searchShortcut + ')';
    sidebarSearchBtn.title = 'Search files (' + searchShortcut + ')';

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
    sidebarSearchBtn.addEventListener('click', openSearchPalette);

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
	for _, f := range files {
		buf.WriteString(fmt.Sprintf(
			`<section id="%s" class="content-section"><article class="markdown-body">%s</article></section>`,
			html.EscapeString(f.ID),
			f.Content,
		))
	}
	return buf.String()
}
