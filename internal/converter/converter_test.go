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
			contains: "<pre",
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

func TestConvert_SyntaxHighlighting(t *testing.T) {
	conv := New()

	tests := []struct {
		name     string
		input    string
		contains []string
	}{
		{
			name:  "go code block has highlighting spans",
			input: "```go\nfunc main() {\n\tfmt.Println(\"hello\")\n}\n```",
			contains: []string{
				"hl-kd", // keyword declaration (func)
				"hl-nf", // function name (main, Println)
				"hl-s",  // string
			},
		},
		{
			name:  "python code block has highlighting spans",
			input: "```python\ndef hello():\n    print(\"world\")\n```",
			contains: []string{
				"hl-k",  // keyword (def)
				"hl-nf", // function name
				"hl-s",  // string
			},
		},
		{
			name:  "javascript code block has highlighting spans",
			input: "```javascript\nconst x = 42;\nconsole.log(x);\n```",
			contains: []string{
				"hl-kr", // keyword reserved (const)
				"hl-mi", // integer
			},
		},
		{
			name:  "bash code block has highlighting spans",
			input: "```bash\necho \"hello world\"\n```",
			contains: []string{
				"hl-nb", // builtin (echo)
				"hl-s",  // string
			},
		},
		{
			name:     "code block without language still renders",
			input:    "```\nplain code\n```",
			contains: []string{"<pre>", "<code>"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := conv.Convert([]byte(tt.input))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			for _, want := range tt.contains {
				if !strings.Contains(result, want) {
					t.Errorf("expected output to contain %q, got: %s", want, result)
				}
			}
		})
	}
}

func TestConvert_SupportedLanguages(t *testing.T) {
	conv := New()

	// Test that common languages are recognized and produce highlighting
	languages := []string{
		"go", "python", "javascript", "typescript",
		"bash", "sh", "json", "yaml", "html", "css",
		"rust", "java", "c", "cpp", "ruby", "sql",
	}

	for _, lang := range languages {
		t.Run(lang, func(t *testing.T) {
			input := "```" + lang + "\ncode here\n```"
			result, err := conv.Convert([]byte(input))
			if err != nil {
				t.Fatalf("unexpected error for %s: %v", lang, err)
			}
			// Should contain chroma wrapper class
			if !strings.Contains(result, "hl-chroma") {
				t.Errorf("expected %s code block to have chroma wrapper, got: %s", lang, result)
			}
		})
	}
}
