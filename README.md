# mdp - Markdown Previewer

A fast CLI tool that previews Markdown files in your browser with GitHub-styled rendering.

## Features

- GitHub Flavored Markdown (GFM) support
- Dark mode support (follows system preference)
- Syntax highlighting for code blocks
- Tables, task lists, strikethrough, and autolinks
- Cross-platform: macOS, Linux, and Windows

## Installation

### Homebrew (macOS)

```bash
brew install sadiksaifi/tap/mdp
```

### From source

```bash
git clone https://github.com/sadiksaifi/mdp.git
cd mdp
make install
```

### Manual build

```bash
make build
./mdp README.md
```

## Usage

```bash
mdp <markdown-file.md>
```

Example:

```bash
mdp README.md
```

This will:
1. Convert the Markdown to styled HTML
2. Save it to `/tmp/mdpreview-README.html`
3. Open it in your default browser

## Development

### Project Structure

```
cmd/mdp/              # CLI entry point
internal/
  converter/          # Markdown to HTML conversion
  template/           # HTML document generation
  browser/            # Platform-specific browser opening
assets/               # CSS assets
```

### Commands

```bash
make build      # Build the binary
make test       # Run all tests
make coverage   # Run tests with coverage
make install    # Install to /usr/local/bin
make release    # Build for all platforms
make clean      # Remove build artifacts
```

### Running Tests

```bash
make test
```

## License

MIT
