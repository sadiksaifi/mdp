# mdp - Markdown Previewer

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Platform](https://img.shields.io/badge/Platform-macOS%20|%20Linux%20|%20Windows-lightgrey)]()

A fast CLI tool that previews Markdown files in your browser with GitHub-styled rendering.

<img width="1512" height="949" alt="Screenshot 2026-01-12 at 14 12 06" src="https://github.com/user-attachments/assets/7e8652d1-1387-4873-bfe3-6f0749f6ae47" />

---

## Features

| Feature | Description |
|---------|-------------|
| **GitHub Flavored Markdown** | Tables, task lists, strikethrough, and autolinks |
| **Syntax Highlighting** | Beautiful code blocks with language support |
| **Dark Mode** | Automatically follows system preference |
| **Multi-file Support** | Preview multiple files with sidebar navigation |
| **Directory Support** | Preview all `.md` files in a directory |
| **Live Reload Server** | Watch files and auto-refresh on changes |
| **Respects `.gitignore`** | Automatically skips ignored files |
| **Mobile Responsive** | Hamburger menu on smaller screens |
| **Fuzzy Search** | Quick file search with `Cmd/Ctrl+K` |

---

## Installation

### Homebrew (macOS)

```bash
brew install sadiksaifi/tap/mdp
```

### From Source

```bash
git clone https://github.com/sadiksaifi/mdp.git
cd mdp
make install
```

### Manual Build

```bash
make build
./mdp README.md
```

---

## Usage

```bash
mdp <file.md>                    # Preview single file
mdp <file1.md> <file2.md>        # Preview multiple files with sidebar
mdp <directory>                  # Preview all .md files in directory
mdp --serve <file.md>            # Start live reload server
mdp --serve --port 3000 <dir>    # Live reload on custom port
```

### Options

| Option | Description |
|--------|-------------|
| `--serve` | Start live reload server instead of opening browser |
| `--port <port>` | Port for live reload server (default: `8080`) |
| `-h, --help` | Show help message |
| `-v, --version` | Show version |

---

## Examples

### Single File

```bash
mdp README.md
```

### Multiple Files

```bash
mdp README.md CHANGELOG.md docs/guide.md
```

### Entire Directory

```bash
mdp ./docs/
```

### Live Reload Server

```bash
mdp --serve README.md              # Start server on port 8080
mdp --serve --port 3000 ./docs/    # Start server on port 3000
```

> [!TIP]
> Use `--serve` when writing documentation to see changes in real-time without manually refreshing the browser.

---

## Output

| Mode | Output |
|------|--------|
| **Single file** | Opens `/tmp/mdpreview-{filename}.html` in your default browser |
| **Multiple files/directory** | Opens `/tmp/mdpreview-multi.html` with sidebar navigation |
| **Live reload mode** | Starts HTTP server at `http://localhost:<port>` with WebSocket auto-refresh |

---

## Keyboard Shortcuts

> [!NOTE]
> These shortcuts work in multi-file mode.

| Shortcut | Action |
|----------|--------|
| <kbd>Cmd/Ctrl</kbd> + <kbd>K</kbd> | Open fuzzy search palette |
| <kbd>Cmd</kbd> + <kbd>B</kbd> (Mac) | Toggle sidebar |
| <kbd>Ctrl</kbd> + <kbd>B</kbd> (Win/Linux) | Toggle sidebar |
| <kbd>Escape</kbd> | Close sidebar/search palette |

### Search Palette Navigation

| Shortcut | Action |
|----------|--------|
| <kbd>↑</kbd> / <kbd>↓</kbd> | Navigate results |
| <kbd>Ctrl</kbd> + <kbd>P</kbd> / <kbd>N</kbd> | Navigate results (alternative) |
| <kbd>Enter</kbd> | Open selected file |
| <kbd>Escape</kbd> | Close palette |

---

## Development

### Project Structure

```
cmd/mdp/              # CLI entry point
internal/
  converter/          # Markdown to HTML conversion
  template/           # HTML document generation (single & multi-file)
  filetree/           # File tree data structure for sidebar
  browser/            # Platform-specific browser opening
  server/             # Live reload HTTP server with WebSocket
assets/               # CSS assets
```

### Commands

| Command | Description |
|---------|-------------|
| `make build` | Build the binary |
| `make test` | Run all tests |
| `make coverage` | Run tests with coverage |
| `make install` | Install to `/usr/local/bin` |
| `make release` | Build for all platforms |
| `make clean` | Remove build artifacts |

### Running Tests

```bash
make test
```

> [!IMPORTANT]
> Always run tests before submitting a pull request.

---

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

---

## License

MIT
