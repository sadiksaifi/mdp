# mdp - Markdown Previewer

A fast CLI tool that previews Markdown files in your browser with GitHub-styled rendering.

## Features

- GitHub Flavored Markdown (GFM) support
- Dark mode support (follows system preference)
- Syntax highlighting for code blocks
- Tables, task lists, strikethrough, and autolinks
- Cross-platform: macOS, Linux, and Windows
- **Multi-file support** with sidebar navigation
- **Directory support** - preview all markdown files in a directory
- Respects `.gitignore` files (including nested) when scanning directories
- Collapsible sidebar with keyboard shortcut (Cmd+B / Ctrl+B)
- Mobile responsive design with hamburger menu

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
mdp <file.md>                    # Preview single file
mdp <file1.md> <file2.md>        # Preview multiple files with sidebar
mdp <directory>                  # Preview all .md files in directory
```

### Examples

**Single file:**
```bash
mdp README.md
```

**Multiple files:**
```bash
mdp README.md CHANGELOG.md docs/guide.md
```

**Entire directory:**
```bash
mdp ./docs/
```

### Output

- **Single file**: Opens `/tmp/mdpreview-{filename}.html` in your default browser
- **Multiple files/directory**: Opens `/tmp/mdpreview-multi.html` with a sidebar for navigation

### Keyboard Shortcuts (Multi-file mode)

| Shortcut | Action |
|----------|--------|
| `Cmd+B` (Mac) / `Ctrl+B` (Win/Linux) | Toggle sidebar |
| `Escape` | Close sidebar (mobile) |

## Development

### Project Structure

```
cmd/mdp/              # CLI entry point
internal/
  converter/          # Markdown to HTML conversion
  template/           # HTML document generation (single & multi-file)
  filetree/           # File tree data structure for sidebar
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
