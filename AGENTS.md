# AGENTS.md

This file provides guidance to AI agents like you when working with code in this repository.

## Build, Test, and Run

```bash
make build            # Build the binary
make test             # Run all tests
make coverage         # Run tests with coverage report
make help             # Show all available targets

# Preview modes
./mdp README.md                     # Single file preview
./mdp README.md CHANGELOG.md        # Multi-file with sidebar
./mdp ./docs/                       # All .md files in directory
```

## Architecture

mdp is a CLI tool that converts Markdown files to styled HTML and opens them in the default browser.

**Project structure:**
```
cmd/mdp/              # Entry point, CLI argument parsing, file discovery
internal/
  converter/          # Markdown to HTML conversion (goldmark)
  template/           # HTML document generation with embedded CSS
    template.go       # Single-file HTML template
    multifile.go      # Multi-file HTML template with sidebar CSS/JS
  filetree/           # File tree data structure for sidebar navigation
  browser/            # Platform-specific browser opening
assets/               # Original CSS file (also embedded in template package)
```

**Single-file flow**: Read .md file → Convert to HTML via goldmark (GFM) → Generate HTML document with embedded GitHub CSS → Write to `/tmp/mdpreview-{filename}.html` → Open in browser

**Multi-file flow**: Resolve files (expand directories, respect .gitignore) → Convert all to HTML → Build file tree structure → Generate HTML with sidebar navigation → Write to `/tmp/mdpreview-multi.html` → Open in browser

**Key implementation details**:
- CSS and JS are embedded into the binary using `//go:embed` in `internal/template/`
- goldmark is configured with GFM extension, auto heading IDs, and unsafe HTML rendering
- Single file output: `/tmp/mdpreview-{filename}.html`
- Multi-file output: `/tmp/mdpreview-multi.html`
- Cross-platform: supports macOS (`open`), Linux (`xdg-open`), and Windows (`rundll32`)
- Directory scanning respects `.gitignore` files at all levels using `github.com/sabhiram/go-gitignore`
- Multi-file sidebar is collapsible on desktop (state persisted in localStorage)
- Mobile responsive: sidebar becomes overlay with hamburger menu at ≤768px
- Keyboard shortcut: Cmd+B (Mac) / Ctrl+B (Win/Linux) toggles sidebar on desktop
