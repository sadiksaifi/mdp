# CLAUDE.md

This file provides guidance to AI agents like you when working with code in this repository.

## Build, Test, and Run

```bash
make build            # Build the binary
make test             # Run all tests
make coverage         # Run tests with coverage report
./mdp README.md       # Preview a markdown file in browser
make help             # Show all available targets
```

## Architecture

mdp is a CLI tool that converts Markdown files to styled HTML and opens them in the default browser.

**Project structure:**
```
cmd/mdp/              # Entry point
internal/
  converter/          # Markdown to HTML conversion (goldmark)
  template/           # HTML document generation with embedded CSS
  browser/            # Platform-specific browser opening
assets/               # Original CSS file (also embedded in template package)
```

**Flow**: Read .md file → Convert to HTML via goldmark (GFM) → Generate HTML document with embedded GitHub CSS → Write to /tmp → Open in browser

**Key implementation details**:
- CSS is embedded into the binary using `//go:embed` in `internal/template/`
- goldmark is configured with GFM extension, auto heading IDs, and unsafe HTML rendering
- Output files go to `/tmp/mdpreview-{filename}.html`
- Cross-platform: supports macOS (`open`), Linux (`xdg-open`), and Windows (`rundll32`)
