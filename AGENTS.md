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

# Live reload server
./mdp --serve README.md             # Start server on port 8080
./mdp --serve --port 3000 ./docs/   # Start server on custom port
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
  server/             # Live reload HTTP server with WebSocket support
assets/               # Original CSS file (also embedded in template package)
```

**Single-file flow**: Read .md file → Convert to HTML via goldmark (GFM) → Generate HTML document with embedded GitHub CSS → Write to `/tmp/mdpreview-{filename}.html` → Open in browser

**Multi-file flow**: Resolve files (expand directories, respect .gitignore) → Convert all to HTML → Build file tree structure → Generate HTML with sidebar navigation → Write to `/tmp/mdpreview-multi.html` → Open in browser

**Live reload flow (--serve)**: Resolve files → Start HTTP server → Watch files with fsnotify → On change: re-convert markdown → Notify clients via WebSocket → Browser auto-refreshes

**Key implementation details**:
- CSS and JS are embedded into the binary using `//go:embed` in `internal/template/`
- goldmark is configured with GFM extension, auto heading IDs, and unsafe HTML rendering
- Single file output: `/tmp/mdpreview-{filename}.html`
- Multi-file output: `/tmp/mdpreview-multi.html`
- Live reload uses WebSocket for instant browser refresh on file changes
- Cross-platform: supports macOS (`open`), Linux (`xdg-open`), and Windows (`rundll32`)
- Directory scanning respects `.gitignore` files at all levels using `github.com/sabhiram/go-gitignore`
- Multi-file sidebar is collapsible on desktop (state persisted in localStorage)
- Mobile responsive: sidebar becomes overlay with hamburger menu at ≤768px
- Keyboard shortcut: Cmd+B (Mac) / Ctrl+B (Win/Linux) toggles sidebar on desktop

## Development Guidelines

**Testing Requirements:**
- Always write tests when adding new features or modifying existing functionality
- Tests should be placed in `*_test.go` files alongside the code being tested
- Run `make test` to execute all tests before submitting changes
- Run `make coverage` to ensure adequate test coverage

**Test Structure:**
- Use table-driven tests where appropriate
- Test both success and error cases
- Use `t.TempDir()` for tests requiring temporary files
- Mock external dependencies (file system, network) when possible

**Before Completing Any Feature:**
1. Ensure all new functions have corresponding test cases
2. Verify existing tests still pass (`make test`)
3. Check for any regressions in functionality
