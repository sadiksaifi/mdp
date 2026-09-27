package template

// themeCSS defines theme tokens once and maps them to the active theme. UI
// components consume the semantic tokens instead of duplicating light/dark
// selectors. The media query remains as a no-JavaScript fallback; data-theme
// has higher specificity and therefore wins after a manual choice.
const themeCSS = `
:root {
    --theme-light-page-bg: #ffffff;
    --theme-light-fg: #1f2328;
    --theme-light-muted: #59636e;
    --theme-light-panel-bg: #f6f8fa;
    --theme-light-border: #d1d9e0;
    --theme-light-hover: #e6e8eb;
    --theme-light-accent: #0969da;
    --theme-light-accent-bg: #ddf4ff;
    --theme-light-input-bg: #ffffff;
    --theme-light-focus-ring: rgba(9, 105, 218, 0.1);
    --theme-light-accent-wash: rgba(9, 105, 218, 0.08);
    --theme-light-highlight: rgba(255, 212, 59, 0.4);
    --theme-light-highlight-active: rgba(255, 212, 59, 0.7);
    --theme-light-danger: #cf222e;
    --theme-light-danger-bg: rgba(207, 34, 46, 0.1);
    --theme-light-success: #1a7f37;
    --theme-light-success-bg: rgba(26, 127, 55, 0.08);
    --theme-light-overlay: rgba(0, 0, 0, 0.5);
    --theme-light-floating-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
    --theme-light-modal-shadow: 0 8px 30px rgba(0, 0, 0, 0.2);
    --theme-light-modal-bg: #ffffff;
    --theme-light-github-bg: #f6f8fa;
    --theme-light-github-border: #d0d7de;
    --theme-light-github-color: #656d76;
    --theme-light-github-hover-bg: #eaeef2;
    --theme-light-github-hover-color: #1f2328;

    --theme-dark-page-bg: #0d1117;
    --theme-dark-fg: #e6edf3;
    --theme-dark-muted: #9198a1;
    --theme-dark-panel-bg: #161b22;
    --theme-dark-border: #3d444d;
    --theme-dark-hover: #21262d;
    --theme-dark-accent: #58a6ff;
    --theme-dark-accent-bg: #388bfd26;
    --theme-dark-input-bg: #0d1117;
    --theme-dark-focus-ring: rgba(88, 166, 255, 0.1);
    --theme-dark-accent-wash: rgba(88, 166, 255, 0.1);
    --theme-dark-highlight: rgba(255, 212, 59, 0.25);
    --theme-dark-highlight-active: rgba(255, 212, 59, 0.45);
    --theme-dark-danger: #f85149;
    --theme-dark-danger-bg: rgba(248, 81, 73, 0.1);
    --theme-dark-success: #3fb950;
    --theme-dark-success-bg: rgba(63, 185, 80, 0.1);
    --theme-dark-overlay: rgba(0, 0, 0, 0.7);
    --theme-dark-floating-shadow: 0 2px 8px rgba(0, 0, 0, 0.3);
    --theme-dark-modal-shadow: 0 8px 30px rgba(0, 0, 0, 0.4);
    --theme-dark-modal-bg: #161b22;
    --theme-dark-github-bg: #21262d;
    --theme-dark-github-border: #30363d;
    --theme-dark-github-color: #8b949e;
    --theme-dark-github-hover-bg: #30363d;
    --theme-dark-github-hover-color: #e6edf3;

    --sidebar-bg: var(--panel-bg);
    --sidebar-border: var(--panel-border);
    --sidebar-hover: var(--panel-hover);
    --sidebar-active: var(--accent-color);
    --sidebar-active-bg: var(--accent-bg);
}

:root,
html[data-theme="light"] {
    color-scheme: light;
    --page-bg: var(--theme-light-page-bg);
    --fg-color: var(--theme-light-fg);
    --fg-muted: var(--theme-light-muted);
    --panel-bg: var(--theme-light-panel-bg);
    --panel-border: var(--theme-light-border);
    --panel-hover: var(--theme-light-hover);
    --accent-color: var(--theme-light-accent);
    --accent-bg: var(--theme-light-accent-bg);
    --input-bg: var(--theme-light-input-bg);
    --focus-ring: var(--theme-light-focus-ring);
    --accent-wash: var(--theme-light-accent-wash);
    --comment-highlight-bg: var(--theme-light-highlight);
    --comment-highlight-active-bg: var(--theme-light-highlight-active);
    --danger-color: var(--theme-light-danger);
    --danger-bg: var(--theme-light-danger-bg);
    --success-color: var(--theme-light-success);
    --success-bg: var(--theme-light-success-bg);
    --overlay-bg: var(--theme-light-overlay);
    --floating-shadow: var(--theme-light-floating-shadow);
    --modal-shadow: var(--theme-light-modal-shadow);
    --modal-bg: var(--theme-light-modal-bg);
    --github-link-bg: var(--theme-light-github-bg);
    --github-link-border: var(--theme-light-github-border);
    --github-link-color: var(--theme-light-github-color);
    --github-link-hover-bg: var(--theme-light-github-hover-bg);
    --github-link-hover-color: var(--theme-light-github-hover-color);
}

@media (prefers-color-scheme: dark) {
    :root {
        color-scheme: dark;
        --page-bg: var(--theme-dark-page-bg);
        --fg-color: var(--theme-dark-fg);
        --fg-muted: var(--theme-dark-muted);
        --panel-bg: var(--theme-dark-panel-bg);
        --panel-border: var(--theme-dark-border);
        --panel-hover: var(--theme-dark-hover);
        --accent-color: var(--theme-dark-accent);
        --accent-bg: var(--theme-dark-accent-bg);
        --input-bg: var(--theme-dark-input-bg);
        --focus-ring: var(--theme-dark-focus-ring);
        --accent-wash: var(--theme-dark-accent-wash);
        --comment-highlight-bg: var(--theme-dark-highlight);
        --comment-highlight-active-bg: var(--theme-dark-highlight-active);
        --danger-color: var(--theme-dark-danger);
        --danger-bg: var(--theme-dark-danger-bg);
        --success-color: var(--theme-dark-success);
        --success-bg: var(--theme-dark-success-bg);
        --overlay-bg: var(--theme-dark-overlay);
        --floating-shadow: var(--theme-dark-floating-shadow);
        --modal-shadow: var(--theme-dark-modal-shadow);
        --modal-bg: var(--theme-dark-modal-bg);
        --github-link-bg: var(--theme-dark-github-bg);
        --github-link-border: var(--theme-dark-github-border);
        --github-link-color: var(--theme-dark-github-color);
        --github-link-hover-bg: var(--theme-dark-github-hover-bg);
        --github-link-hover-color: var(--theme-dark-github-hover-color);
    }
}

html[data-theme="dark"] {
    color-scheme: dark;
    --page-bg: var(--theme-dark-page-bg);
    --fg-color: var(--theme-dark-fg);
    --fg-muted: var(--theme-dark-muted);
    --panel-bg: var(--theme-dark-panel-bg);
    --panel-border: var(--theme-dark-border);
    --panel-hover: var(--theme-dark-hover);
    --accent-color: var(--theme-dark-accent);
    --accent-bg: var(--theme-dark-accent-bg);
    --input-bg: var(--theme-dark-input-bg);
    --focus-ring: var(--theme-dark-focus-ring);
    --accent-wash: var(--theme-dark-accent-wash);
    --comment-highlight-bg: var(--theme-dark-highlight);
    --comment-highlight-active-bg: var(--theme-dark-highlight-active);
    --danger-color: var(--theme-dark-danger);
    --danger-bg: var(--theme-dark-danger-bg);
    --success-color: var(--theme-dark-success);
    --success-bg: var(--theme-dark-success-bg);
    --overlay-bg: var(--theme-dark-overlay);
    --floating-shadow: var(--theme-dark-floating-shadow);
    --modal-shadow: var(--theme-dark-modal-shadow);
    --modal-bg: var(--theme-dark-modal-bg);
    --github-link-bg: var(--theme-dark-github-bg);
    --github-link-border: var(--theme-dark-github-border);
    --github-link-color: var(--theme-dark-github-color);
    --github-link-hover-bg: var(--theme-dark-github-hover-bg);
    --github-link-hover-color: var(--theme-dark-github-hover-color);
}

body {
    background-color: var(--page-bg);
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
.theme-toggle-btn svg {
    transform: scale(0.85);
}
html[data-theme="dark"] .theme-toggle-btn .theme-icon-moon {
    display: none;
}
html[data-theme="dark"] .theme-toggle-btn .theme-icon-sun {
    display: block;
}
`

// themeJSInline wires up every theme toggle: applies the saved theme (or the
// OS default) without persisting, stores the choice only on click, and follows
// the OS only while no explicit choice was saved. It omits <script> tags for
// templates that already provide a script wrapper. The leading semicolon
// guards against concatenation with a preceding unterminated statement.
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
