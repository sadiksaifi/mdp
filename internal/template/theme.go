package template

// themeInitScript runs in <head> before paint to avoid a theme flash.
// It restores the saved choice, falling back to the OS preference.
const themeInitScript = `<script>try{var t=localStorage.getItem('mdp-theme');if(!t){t=(window.matchMedia&&window.matchMedia('(prefers-color-scheme: dark)').matches)?'dark':'light';}document.documentElement.dataset.theme=t;}catch(e){}</script>`

// themeToggleButton is rendered first in .topbar-right in both templates.
const themeToggleButton = `<button class="topbar-btn topbar-theme-btn" aria-label="Toggle theme" title="Toggle theme">
                <svg class="theme-icon-moon" xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z"></path></svg>
                <svg class="theme-icon-sun" xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="5"></circle><line x1="12" y1="1" x2="12" y2="3"></line><line x1="12" y1="21" x2="12" y2="23"></line><line x1="4.22" y1="4.22" x2="5.64" y2="5.64"></line><line x1="18.36" y1="18.36" x2="19.78" y2="19.78"></line><line x1="1" y1="12" x2="3" y2="12"></line><line x1="21" y1="12" x2="23" y2="12"></line><line x1="4.22" y1="19.78" x2="5.64" y2="18.36"></line><line x1="18.36" y1="5.64" x2="19.78" y2="4.22"></line></svg>
            </button>
            <div class="topbar-divider"></div>`

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
.topbar-theme-btn .theme-icon-sun {
    display: none;
}
/* Theme glyphs run edge-to-edge in their viewBox, so scale them down
   slightly to match the optical size of the other header icons.
   Layout box (and the 36px button) is unchanged. */
.topbar-theme-btn svg {
    transform: scale(0.85);
}
html[data-theme="dark"] .topbar-theme-btn .theme-icon-moon {
    display: none;
}
html[data-theme="dark"] .topbar-theme-btn .theme-icon-sun {
    display: block;
}
`

// themeJS wires up the toggle button: applies the saved theme (or the OS
// default) without persisting, stores the choice only on click, and follows
// the OS only while no explicit choice was saved.
const themeJS = `
    <script>
        (function() {
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
                var btn = document.querySelector('.topbar-theme-btn');
                if (btn) {
                    var label = t === 'dark' ? 'Switch to light mode' : 'Switch to dark mode';
                    btn.setAttribute('aria-label', label);
                    btn.setAttribute('title', label);
                }
            }
            applyTheme(readSavedTheme() || root.dataset.theme || systemTheme(), false);
            var btn = document.querySelector('.topbar-theme-btn');
            if (btn) {
                btn.addEventListener('click', function() {
                    applyTheme(root.dataset.theme === 'dark' ? 'light' : 'dark', true);
                });
            }
            if (window.matchMedia) {
                window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', function() {
                    if (!readSavedTheme()) applyTheme(systemTheme(), false);
                });
            }
        })();
    </script>`
