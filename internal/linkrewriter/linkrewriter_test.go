package linkrewriter

import (
	"testing"

	"mdp/internal/filetree"
)

func TestNew(t *testing.T) {
	entries := []filetree.FileEntry{
		{ID: "readme-md", RelPath: "README.md"},
		{ID: "docs-guide-md", RelPath: "docs/guide.md"},
		{ID: "docs-api-types-md", RelPath: "docs/api/types.md"},
	}

	lr := New(entries)

	// Verify the path to ID mapping was created correctly
	if lr.pathToID["readme.md"] != "readme-md" {
		t.Errorf("expected pathToID['readme.md'] = 'readme-md', got %q", lr.pathToID["readme.md"])
	}
	if lr.pathToID["docs/guide.md"] != "docs-guide-md" {
		t.Errorf("expected pathToID['docs/guide.md'] = 'docs-guide-md', got %q", lr.pathToID["docs/guide.md"])
	}
}

func TestRewriteLinks(t *testing.T) {
	entries := []filetree.FileEntry{
		{ID: "readme-md", RelPath: "README.md"},
		{ID: "docs-guide-md", RelPath: "docs/guide.md"},
		{ID: "docs-api-md", RelPath: "docs/api.md"},
		{ID: "docs-api-types-md", RelPath: "docs/api/types.md"},
	}

	lr := New(entries)

	tests := []struct {
		name          string
		html          string
		sourceRelPath string
		expected      string
	}{
		{
			name:          "simple relative link from root",
			html:          `<a href="./docs/guide.md">Guide</a>`,
			sourceRelPath: "README.md",
			expected:      `<a href="#docs-guide-md">Guide</a>`,
		},
		{
			name:          "relative link without ./",
			html:          `<a href="docs/guide.md">Guide</a>`,
			sourceRelPath: "README.md",
			expected:      `<a href="#docs-guide-md">Guide</a>`,
		},
		{
			name:          "parent directory link",
			html:          `<a href="../README.md">Back</a>`,
			sourceRelPath: "docs/guide.md",
			expected:      `<a href="#readme-md">Back</a>`,
		},
		{
			name:          "sibling file link",
			html:          `<a href="./api.md">API</a>`,
			sourceRelPath: "docs/guide.md",
			expected:      `<a href="#docs-api-md">API</a>`,
		},
		{
			name:          "nested relative link",
			html:          `<a href="./api/types.md">Types</a>`,
			sourceRelPath: "docs/guide.md",
			expected:      `<a href="#docs-api-types-md">Types</a>`,
		},
		{
			name:          "external http link unchanged",
			html:          `<a href="https://example.com">Example</a>`,
			sourceRelPath: "README.md",
			expected:      `<a href="https://example.com">Example</a>`,
		},
		{
			name:          "external https link unchanged",
			html:          `<a href="https://github.com/test">GitHub</a>`,
			sourceRelPath: "README.md",
			expected:      `<a href="https://github.com/test">GitHub</a>`,
		},
		{
			name:          "anchor-only link unchanged",
			html:          `<a href="#section">Section</a>`,
			sourceRelPath: "README.md",
			expected:      `<a href="#section">Section</a>`,
		},
		{
			name:          "link to file not in project unchanged",
			html:          `<a href="./missing.md">Missing</a>`,
			sourceRelPath: "README.md",
			expected:      `<a href="./missing.md">Missing</a>`,
		},
		{
			name:          "link with anchor stripped",
			html:          `<a href="./docs/guide.md#installation">Guide Install</a>`,
			sourceRelPath: "README.md",
			expected:      `<a href="#docs-guide-md">Guide Install</a>`,
		},
		{
			name:          "non-md link unchanged",
			html:          `<a href="./image.png">Image</a>`,
			sourceRelPath: "README.md",
			expected:      `<a href="./image.png">Image</a>`,
		},
		{
			name:          "mailto link unchanged",
			html:          `<a href="mailto:test@example.com">Email</a>`,
			sourceRelPath: "README.md",
			expected:      `<a href="mailto:test@example.com">Email</a>`,
		},
		{
			name:          "multiple links in content",
			html:          `<p><a href="./docs/guide.md">Guide</a> and <a href="./docs/api.md">API</a></p>`,
			sourceRelPath: "README.md",
			expected:      `<p><a href="#docs-guide-md">Guide</a> and <a href="#docs-api-md">API</a></p>`,
		},
		{
			name:          "link with extra attributes",
			html:          `<a class="link" href="./docs/guide.md" target="_blank">Guide</a>`,
			sourceRelPath: "README.md",
			expected:      `<a class="link" href="#docs-guide-md" target="_blank">Guide</a>`,
		},
		{
			name:          "URL encoded path",
			html:          `<a href="./docs/guide.md">Guide</a>`,
			sourceRelPath: "README.md",
			expected:      `<a href="#docs-guide-md">Guide</a>`,
		},
		{
			name:          "case insensitive matching",
			html:          `<a href="../README.MD">Readme</a>`,
			sourceRelPath: "docs/guide.md",
			expected:      `<a href="#readme-md">Readme</a>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lr.RewriteLinks(tt.html, tt.sourceRelPath)
			if result != tt.expected {
				t.Errorf("RewriteLinks(%q, %q) = %q, want %q", tt.html, tt.sourceRelPath, result, tt.expected)
			}
		})
	}
}

func TestNormalizePath(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"README.md", "readme.md"},
		{"docs/guide.md", "docs/guide.md"},
		{"./docs/guide.md", "docs/guide.md"},
		{"docs\\guide.md", "docs/guide.md"},
		{"DOCS/GUIDE.MD", "docs/guide.md"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := normalizePath(tt.input)
			if result != tt.expected {
				t.Errorf("normalizePath(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestRewriteHref(t *testing.T) {
	entries := []filetree.FileEntry{
		{ID: "readme-md", RelPath: "README.md"},
		{ID: "docs-guide-md", RelPath: "docs/guide.md"},
	}

	lr := New(entries)

	tests := []struct {
		name      string
		href      string
		sourceDir string
		expected  string
	}{
		{"external http", "http://example.com", "", "http://example.com"},
		{"external https", "https://example.com", "", "https://example.com"},
		{"fragment only", "#section", "", "#section"},
		{"mailto", "mailto:test@example.com", "", "mailto:test@example.com"},
		{"non-md file", "image.png", "", "image.png"},
		{"valid md link", "README.md", "", "#readme-md"},
		{"valid nested md link", "docs/guide.md", "", "#docs-guide-md"},
		{"parent relative", "../README.md", "docs", "#readme-md"},
		{"not found", "missing.md", "", "missing.md"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lr.rewriteHref(tt.href, tt.sourceDir)
			if result != tt.expected {
				t.Errorf("rewriteHref(%q, %q) = %q, want %q", tt.href, tt.sourceDir, result, tt.expected)
			}
		})
	}
}
