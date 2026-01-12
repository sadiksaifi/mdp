package converter

import (
	"strings"
	"testing"
)

func TestConvert_BasicMarkdown(t *testing.T) {
	conv := New()

	tests := []struct {
		name     string
		input    string
		contains string
	}{
		{
			name:     "heading",
			input:    "# Hello World",
			contains: "<h1",
		},
		{
			name:     "paragraph",
			input:    "This is a paragraph.",
			contains: "<p>This is a paragraph.</p>",
		},
		{
			name:     "bold",
			input:    "**bold text**",
			contains: "<strong>bold text</strong>",
		},
		{
			name:     "italic",
			input:    "*italic text*",
			contains: "<em>italic text</em>",
		},
		{
			name:     "link",
			input:    "[link](https://example.com)",
			contains: `<a href="https://example.com">link</a>`,
		},
		{
			name:     "code block",
			input:    "```go\nfmt.Println(\"hello\")\n```",
			contains: "<pre><code",
		},
		{
			name:     "inline code",
			input:    "`code`",
			contains: "<code>code</code>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := conv.Convert([]byte(tt.input))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !strings.Contains(result, tt.contains) {
				t.Errorf("expected output to contain %q, got: %s", tt.contains, result)
			}
		})
	}
}

func TestConvert_GFMFeatures(t *testing.T) {
	conv := New()

	tests := []struct {
		name     string
		input    string
		contains string
	}{
		{
			name:     "strikethrough",
			input:    "~~strikethrough~~",
			contains: "<del>strikethrough</del>",
		},
		{
			name:     "table",
			input:    "| A | B |\n|---|---|\n| 1 | 2 |",
			contains: "<table>",
		},
		{
			name:     "task list",
			input:    "- [ ] unchecked\n- [x] checked",
			contains: `type="checkbox"`,
		},
		{
			name:     "autolink",
			input:    "https://example.com",
			contains: `<a href="https://example.com">`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := conv.Convert([]byte(tt.input))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !strings.Contains(result, tt.contains) {
				t.Errorf("expected output to contain %q, got: %s", tt.contains, result)
			}
		})
	}
}

func TestConvert_HeadingIDs(t *testing.T) {
	conv := New()

	input := "# Hello World"
	result, err := conv.Convert([]byte(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(result, `id="hello-world"`) {
		t.Errorf("expected heading to have auto-generated ID, got: %s", result)
	}
}

func TestConvert_UnsafeHTML(t *testing.T) {
	conv := New()

	input := "<div class=\"custom\">Raw HTML</div>"
	result, err := conv.Convert([]byte(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(result, `<div class="custom">Raw HTML</div>`) {
		t.Errorf("expected raw HTML to be preserved, got: %s", result)
	}
}

func TestConvert_EmptyInput(t *testing.T) {
	conv := New()

	result, err := conv.Convert([]byte(""))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != "" {
		t.Errorf("expected empty output for empty input, got: %s", result)
	}
}
