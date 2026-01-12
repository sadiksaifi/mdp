package converter

import (
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

// Converter handles markdown to HTML conversion.
type Converter struct {
	md goldmark.Markdown
}

// New creates a new Converter with GFM support.
func New() *Converter {
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithUnsafe(), // Allow raw HTML in markdown
		),
	)
	return &Converter{md: md}
}

// Convert transforms markdown content into HTML.
func (c *Converter) Convert(markdown []byte) (string, error) {
	var buf strings.Builder
	if err := c.md.Convert(markdown, &buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}
