# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build and Run

```bash
make build            # Build the binary
./mdp README.md       # Preview a markdown file in browser
make help             # Show all available targets
```

## Architecture

mdp is a single-file CLI tool that converts Markdown files to styled HTML and opens them in the default browser.

**Flow**: Read .md file -> Convert to HTML via goldmark (GFM) -> Wrap in HTML template with embedded GitHub CSS -> Write to /tmp -> Open with `open` command

**Key implementation details**:
- CSS is embedded into the binary using `//go:embed assets/github-markdown.min.css`
- goldmark is configured with GFM extension, auto heading IDs, and unsafe HTML rendering
- Output files go to `/tmp/mdpreview-{filename}.html`
- macOS-only: uses the `open` command for browser launching
