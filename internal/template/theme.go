package template

// themeCSS overrides the prefers-color-scheme defaults when the user picks
// a theme explicitly. It covers both the custom chrome variables used by the
// single-file (panel-*) and multi-file (sidebar-*) templates and the GitHub
// markdown variables for the content body.
const themeCSS = `
html[data-theme="light"] {
    color-scheme: light;
    --fg-color: #1f2328;
    --fg-muted: #59636e;
    --panel-bg: #f6f8fa;
    --panel-border: #d1d9e0;
    --panel-hover: #e6e8eb;
    --accent-color: #0969da;
    --accent-bg: #ddf4ff;
    --sidebar-bg: #f6f8fa;
    --sidebar-border: #d1d9e0;
    --sidebar-hover: #e6e8eb;
    --sidebar-active: #0969da;
    --sidebar-active-bg: #ddf4ff;
}
html[data-theme="dark"] {
    color-scheme: dark;
    --fg-color: #e6edf3;
    --fg-muted: #9198a1;
    --panel-bg: #161b22;
    --panel-border: #3d444d;
    --panel-hover: #21262d;
    --accent-color: #58a6ff;
    --accent-bg: #388bfd26;
    --sidebar-bg: #161b22;
    --sidebar-border: #3d444d;
    --sidebar-hover: #21262d;
    --sidebar-active: #58a6ff;
    --sidebar-active-bg: #388bfd26;
}
html[data-theme="light"] body {
    background-color: #ffffff;
}
html[data-theme="dark"] body {
    background-color: #0d1117;
}
html[data-theme="light"] .markdown-body {
    color-scheme: light;--focus-outlineColor:#0969da;--fgColor-default:#1f2328;--fgColor-muted:#59636e;--fgColor-accent:#0969da;--fgColor-success:#1a7f37;--fgColor-attention:#9a6700;--fgColor-danger:#d1242f;--fgColor-done:#8250df;--bgColor-default:#ffffff;--bgColor-muted:#f6f8fa;--bgColor-neutral-muted:#818b981f;--bgColor-attention-muted:#fff8c5;--borderColor-default:#d1d9e0;--borderColor-muted:#d1d9e0b3;--borderColor-neutral-muted:#d1d9e0b3;--borderColor-accent-emphasis:#0969da;--borderColor-success-emphasis:#1a7f37;--borderColor-attention-emphasis:#9a6700;--borderColor-danger-emphasis:#cf222e;--borderColor-done-emphasis:#8250df;--color-prettylights-syntax-comment:#59636e;--color-prettylights-syntax-constant:#0550ae;--color-prettylights-syntax-constant-other-reference-link:#0a3069;--color-prettylights-syntax-entity:#6639ba;--color-prettylights-syntax-storage-modifier-import:#1f2328;--color-prettylights-syntax-entity-tag:#0550ae;--color-prettylights-syntax-keyword:#cf222e;--color-prettylights-syntax-string:#0a3069;--color-prettylights-syntax-variable:#953800;--color-prettylights-syntax-brackethighlighter-unmatched:#82071e;--color-prettylights-syntax-brackethighlighter-angle:#59636e;--color-prettylights-syntax-invalid-illegal-text:#f6f8fa;--color-prettylights-syntax-invalid-illegal-bg:#82071e;--color-prettylights-syntax-carriage-return-text:#f6f8fa;--color-prettylights-syntax-carriage-return-bg:#cf222e;--color-prettylights-syntax-string-regexp:#116329;--color-prettylights-syntax-markup-list:#3b2300;--color-prettylights-syntax-markup-heading:#0550ae;--color-prettylights-syntax-markup-italic:#1f2328;--color-prettylights-syntax-markup-bold:#1f2328;--color-prettylights-syntax-markup-deleted-text:#82071e;--color-prettylights-syntax-markup-deleted-bg:#ffebe9;--color-prettylights-syntax-markup-inserted-text:#116329;--color-prettylights-syntax-markup-inserted-bg:#dafbe1;--color-prettylights-syntax-markup-changed-text:#953800;--color-prettylights-syntax-markup-changed-bg:#ffd8b5;--color-prettylights-syntax-markup-ignored-text:#d1d9e0;--color-prettylights-syntax-markup-ignored-bg:#0550ae;--color-prettylights-syntax-meta-diff-range:#8250df;--color-prettylights-syntax-sublimelinter-gutter-mark:#818b98;
}
html[data-theme="dark"] .markdown-body {
    color-scheme: dark;--focus-outlineColor:#1f6feb;--fgColor-default:#f0f6fc;--fgColor-muted:#9198a1;--fgColor-accent:#4493f8;--fgColor-success:#3fb950;--fgColor-attention:#d29922;--fgColor-danger:#f85149;--fgColor-done:#ab7df8;--bgColor-default:#0d1117;--bgColor-muted:#151b23;--bgColor-neutral-muted:#656c7633;--bgColor-attention-muted:#bb800926;--borderColor-default:#3d444d;--borderColor-muted:#3d444db3;--borderColor-neutral-muted:#3d444db3;--borderColor-accent-emphasis:#1f6feb;--borderColor-success-emphasis:#238636;--borderColor-attention-emphasis:#9e6a03;--borderColor-danger-emphasis:#da3633;--borderColor-done-emphasis:#8957e5;--color-prettylights-syntax-comment:#9198a1;--color-prettylights-syntax-constant:#79c0ff;--color-prettylights-syntax-constant-other-reference-link:#a5d6ff;--color-prettylights-syntax-entity:#d2a8ff;--color-prettylights-syntax-storage-modifier-import:#f0f6fc;--color-prettylights-syntax-entity-tag:#7ee787;--color-prettylights-syntax-keyword:#ff7b72;--color-prettylights-syntax-string:#a5d6ff;--color-prettylights-syntax-variable:#ffa657;--color-prettylights-syntax-brackethighlighter-unmatched:#f85149;--color-prettylights-syntax-brackethighlighter-angle:#9198a1;--color-prettylights-syntax-invalid-illegal-text:#f0f6fc;--color-prettylights-syntax-invalid-illegal-bg:#8e1519;--color-prettylights-syntax-carriage-return-text:#f0f6fc;--color-prettylights-syntax-carriage-return-bg:#b62324;--color-prettylights-syntax-string-regexp:#7ee787;--color-prettylights-syntax-markup-list:#f2cc60;--color-prettylights-syntax-markup-heading:#1f6feb;--color-prettylights-syntax-markup-italic:#f0f6fc;--color-prettylights-syntax-markup-bold:#f0f6fc;--color-prettylights-syntax-markup-deleted-text:#ffdcd7;--color-prettylights-syntax-markup-deleted-bg:#67060c;--color-prettylights-syntax-markup-inserted-text:#aff5b4;--color-prettylights-syntax-markup-inserted-bg:#033a16;--color-prettylights-syntax-markup-changed-text:#ffdfb6;--color-prettylights-syntax-markup-changed-bg:#5a1e02;--color-prettylights-syntax-markup-ignored-text:#f0f6fc;--color-prettylights-syntax-markup-ignored-bg:#1158c7;--color-prettylights-syntax-meta-diff-range:#d2a8ff;--color-prettylights-syntax-sublimelinter-gutter-mark:#3d444d;
}
.theme-toggle-btn .theme-icon-sun {
    display: none;
}
/* Theme glyphs run edge-to-edge in their viewBox, so scale them down
   slightly to match the optical size of the other header icons.
   Layout box (and the 36px button) is unchanged. */
.theme-toggle-btn svg {
    transform: scale(0.85);
}
html[data-theme="dark"] .theme-toggle-btn .theme-icon-moon {
    display: none;
}
html[data-theme="dark"] .theme-toggle-btn .theme-icon-sun {
    display: block;
}
/* Manual-theme counterparts for the hardcoded prefers-color-scheme
   comment/shortcuts rules in both templates. The head script always sets
data-theme, so these must reproduce the OS rendering exactly when the
   manual choice matches the OS. Values mirror the light bases and dark
   media blocks; both templates resolve to the same colors except where
   noted with body.mdp-single. */
html[data-theme="dark"] .comment-btn {
    background: #58a6ff;
}
html[data-theme="dark"] .comment-highlight {
    background-color: rgba(255, 212, 59, 0.25);
}
html[data-theme="dark"] .comment-highlight:hover,
html[data-theme="dark"] .comment-highlight.active {
    background-color: rgba(255, 212, 59, 0.45);
}
html[data-theme="dark"] .comments-panel {
    background: #161b22;
    border-left-color: #3d444d;
}
html[data-theme="dark"] .comments-list::-webkit-scrollbar-thumb {
    background-color: #3d444d;
}
html[data-theme="dark"] .comment-entry {
    border-bottom-color: #3d444d;
}
html[data-theme="dark"] .comment-entry:hover {
    background: #21262d;
}
html[data-theme="dark"] .comment-entry.active {
    background: #388bfd26;
}
html[data-theme="dark"] .comment-quote {
    background: #21262d;
    border-left-color: #58a6ff;
    color: #9198a1;
}
html[data-theme="dark"] .comment-text {
    color: #e6edf3;
}
html[data-theme="dark"] .comment-copy-btn,
html[data-theme="dark"] .comment-edit-btn,
html[data-theme="dark"] .comment-delete-btn {
    color: #9198a1;
}
html[data-theme="dark"] .comment-copy-btn:hover,
html[data-theme="dark"] .comment-edit-btn:hover {
    color: #58a6ff;
}
html[data-theme="dark"] .comment-copy-btn.copied {
    color: #3fb950;
}
html[data-theme="dark"] .comment-delete-btn:hover {
    color: #f85149;
    background: rgba(248, 81, 73, 0.1);
}
html[data-theme="dark"] .comment-input-form {
    background: #161b22;
    border-bottom-color: #3d444d;
}
html[data-theme="dark"] .comment-input-quote {
    background: #21262d;
    border-left-color: #58a6ff;
    color: #9198a1;
}
html[data-theme="dark"] .comment-input-textarea {
    background: #0d1117;
    border-color: #3d444d;
    color: #e6edf3;
}
html[data-theme="dark"] .comment-input-textarea:focus {
    border-color: #58a6ff;
    box-shadow: 0 0 0 3px rgba(88, 166, 255, 0.1);
}
html[data-theme="dark"] .comment-cancel-btn {
    border-color: #3d444d;
    color: #e6edf3;
}
html[data-theme="dark"] .comment-cancel-btn:hover {
    background: #21262d;
}
html[data-theme="dark"] .comment-save-btn {
    background: #58a6ff;
}
html[data-theme="dark"] .open-comments-btn {
    background: #161b22;
    border-color: #3d444d;
    color: #9198a1;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.3);
}
html[data-theme="dark"] .comment-count {
    background: #58a6ff;
}
html[data-theme="dark"] .comments-panel-footer {
    border-top-color: #3d444d;
}
html[data-theme="dark"] .copy-comments-btn {
    background: transparent;
    color: #9198a1;
}
html[data-theme="dark"] .copy-comments-btn.copied {
    color: #3fb950;
    background: rgba(63, 185, 80, 0.1);
}
html[data-theme="dark"] .comments-empty {
    color: #9198a1;
}
html[data-theme="dark"] .comments-empty kbd {
    background: #21262d;
}
html[data-theme="dark"] .shortcuts-modal-overlay {
    background: rgba(0, 0, 0, 0.7);
}
html[data-theme="dark"] .shortcuts-modal {
    background: #161b22;
    box-shadow: 0 8px 30px rgba(0, 0, 0, 0.4);
}
html[data-theme="dark"] .shortcuts-modal-header {
    border-bottom-color: #3d444d;
}
html[data-theme="dark"] .shortcuts-modal-header h2 {
    color: #e6edf3;
}
html[data-theme="dark"] .shortcuts-modal-close {
    color: #9198a1;
}
html[data-theme="dark"] .shortcuts-modal-close:hover {
    color: #e6edf3;
    background: #21262d;
}
html[data-theme="dark"] .shortcut-row:not(:last-child) {
    border-bottom-color: #3d444d;
}
html[data-theme="dark"] .shortcut-action {
    color: #e6edf3;
}
html[data-theme="dark"] .shortcut-keys kbd {
    background: #21262d;
    border-color: #3d444d;
    color: #e6edf3;
}
/* Single-template values that differ from the multi-file template:
   the copy-button hover wash and the shortcuts modal background stay
   variable-driven in multi-file and already follow the toggle there. */
html[data-theme="dark"] body.mdp-single .comment-copy-btn:hover,
html[data-theme="dark"] body.mdp-single .comment-edit-btn:hover {
    background: rgba(88, 166, 255, 0.1);
}
html[data-theme="dark"] body.mdp-single .copy-comments-btn:hover {
    color: #58a6ff;
    background: rgba(88, 166, 255, 0.1);
}
html[data-theme="light"] .comment-btn {
    background: #0969da;
}
html[data-theme="light"] .comment-highlight {
    background-color: rgba(255, 212, 59, 0.4);
}
html[data-theme="light"] .comment-highlight:hover,
html[data-theme="light"] .comment-highlight.active {
    background-color: rgba(255, 212, 59, 0.7);
}
html[data-theme="light"] .comments-panel {
    background: #f6f8fa;
    border-left-color: #d1d9e0;
}
html[data-theme="light"] .comments-list::-webkit-scrollbar-thumb {
    background-color: #d1d9e0;
}
html[data-theme="light"] .comment-entry {
    border-bottom-color: #d1d9e0;
}
html[data-theme="light"] .comment-entry:hover {
    background: #e6e8eb;
}
html[data-theme="light"] .comment-entry.active {
    background: #ddf4ff;
}
html[data-theme="light"] .comment-quote {
    background: #e6e8eb;
    border-left-color: #0969da;
    color: #59636e;
}
html[data-theme="light"] .comment-text {
    color: #1f2328;
}
html[data-theme="light"] .comment-copy-btn,
html[data-theme="light"] .comment-edit-btn,
html[data-theme="light"] .comment-delete-btn {
    color: #59636e;
}
html[data-theme="light"] .comment-copy-btn:hover,
html[data-theme="light"] .comment-edit-btn:hover {
    color: #0969da;
}
html[data-theme="light"] .comment-copy-btn.copied {
    color: #1a7f37;
}
html[data-theme="light"] .comment-delete-btn:hover {
    color: #cf222e;
    background: rgba(207, 34, 46, 0.1);
}
html[data-theme="light"] .comment-input-form {
    background: #f6f8fa;
    border-bottom-color: #d1d9e0;
}
html[data-theme="light"] .comment-input-quote {
    background: #e6e8eb;
    border-left-color: #0969da;
    color: #59636e;
}
html[data-theme="light"] .comment-input-textarea {
    background: #ffffff;
    border-color: #d1d9e0;
    color: #1f2328;
}
html[data-theme="light"] .comment-input-textarea:focus {
    border-color: #0969da;
    box-shadow: 0 0 0 3px rgba(9, 105, 218, 0.1);
}
html[data-theme="light"] .comment-cancel-btn {
    border-color: #d1d9e0;
    color: #1f2328;
}
html[data-theme="light"] .comment-cancel-btn:hover {
    background: #e6e8eb;
}
html[data-theme="light"] .comment-save-btn {
    background: #0969da;
}
html[data-theme="light"] .open-comments-btn {
    background: #f6f8fa;
    border-color: #d1d9e0;
    color: #59636e;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
}
html[data-theme="light"] .comment-count {
    background: #0969da;
}
html[data-theme="light"] .comments-panel-footer {
    border-top-color: #d1d9e0;
}
html[data-theme="light"] .copy-comments-btn {
    background: transparent;
    color: #59636e;
}
html[data-theme="light"] .copy-comments-btn.copied {
    color: #1a7f37;
    background: rgba(26, 127, 55, 0.08);
}
html[data-theme="light"] .comments-empty {
    color: #59636e;
}
html[data-theme="light"] .comments-empty kbd {
    background: #e6e8eb;
}
html[data-theme="light"] .shortcuts-modal-overlay {
    background: rgba(0, 0, 0, 0.5);
}
html[data-theme="light"] .shortcuts-modal {
    box-shadow: 0 8px 30px rgba(0, 0, 0, 0.2);
}
html[data-theme="light"] .shortcuts-modal-header {
    border-bottom-color: #d1d9e0;
}
html[data-theme="light"] .shortcuts-modal-header h2 {
    color: #1f2328;
}
html[data-theme="light"] .shortcuts-modal-close {
    color: #59636e;
}
html[data-theme="light"] .shortcuts-modal-close:hover {
    color: #1f2328;
    background: #e6e8eb;
}
html[data-theme="light"] .shortcut-row:not(:last-child) {
    border-bottom-color: #e6e8eb;
}
html[data-theme="light"] .shortcut-action {
    color: #1f2328;
}
html[data-theme="light"] .shortcut-keys kbd {
    background: #f6f8fa;
    border-color: #d1d9e0;
    color: #1f2328;
}
html[data-theme="light"] body.mdp-single .comment-copy-btn:hover,
html[data-theme="light"] body.mdp-single .comment-edit-btn:hover {
    background: rgba(9, 105, 218, 0.08);
}
html[data-theme="light"] body.mdp-single .copy-comments-btn:hover {
    color: #0969da;
    background: rgba(9, 105, 218, 0.08);
}
html[data-theme="light"] body.mdp-single .shortcuts-modal {
    background: #ffffff;
}
`

// themeJSInline wires up the toggle button: applies the saved theme (or the OS
// default) without persisting, stores the choice only on click, and follows
// the OS only while no explicit choice was saved.
// themeJSInline is the raw toggle script without <script> tags, for templates
// that already provide their own script wrapper (e.g. the multi-file
// sidebar script block). The leading semicolon guards against concatenation
// with a previous statement missing its trailing semicolon.
const themeJSInline = `;(function() {
            var root = document.documentElement;
            function systemTheme() {
                return (window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches) ? 'dark' : 'light';
            }
            function readSavedTheme() {
                try { return localStorage.getItem('mdp-theme'); } catch (e) { return null; }
            }
            function applyTheme(t, persist) {
                root.dataset.theme = t;
                if (persist) {
                    try { localStorage.setItem('mdp-theme', t); } catch (e) {}
                }
                var label = t === 'dark' ? 'Switch to light mode' : 'Switch to dark mode';
                document.querySelectorAll('.theme-toggle-btn').forEach(function(btn) {
                    btn.setAttribute('aria-label', label);
                    btn.setAttribute('title', label);
                });
            }
            applyTheme(readSavedTheme() || root.dataset.theme || systemTheme(), false);
            document.querySelectorAll('.theme-toggle-btn').forEach(function(btn) {
                btn.addEventListener('click', function() {
                    applyTheme(root.dataset.theme === 'dark' ? 'light' : 'dark', true);
                    window.dispatchEvent(new Event('mdp-theme-change'));
                });
            });
            if (window.matchMedia) {
                window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', function() {
                    if (!readSavedTheme()) applyTheme(systemTheme(), false);
                });
            }
        })();`

// themeJS is the standalone script block for the single-file template.
const themeJS = "\n    <script>\n        " + themeJSInline + "\n    </script>"
