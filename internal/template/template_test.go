package template

import (
	"strings"
	"testing"

	"mdp/internal/filetree"
)

func TestGenerate_ContainsTitle(t *testing.T) {
	result := Generate("Test Document", "<p>Content</p>")

	if !strings.Contains(result, "<title>Test Document</title>") {
		t.Errorf("expected title in output, got: %s", result)
	}
}

func TestGenerate_ContainsContent(t *testing.T) {
	content := "<p>Hello World</p>"
	result := Generate("Test", content)

	if !strings.Contains(result, content) {
		t.Errorf("expected content in output, got: %s", result)
	}
}

func TestGenerate_ContainsCSS(t *testing.T) {
	result := Generate("Test", "<p>Content</p>")

	// Check for markdown-body class styling
	if !strings.Contains(result, ".markdown-body") {
		t.Error("expected CSS to be embedded in output")
	}
}

func TestGenerate_ValidHTMLStructure(t *testing.T) {
	result := Generate("Test", "<p>Content</p>")

	checks := []string{
		"<!DOCTYPE html>",
		"<html lang=\"en\">",
		"<head>",
		"</head>",
		"<body>",
		"</body>",
		"</html>",
		"<meta charset=\"UTF-8\">",
		"<meta name=\"viewport\"",
		"<article class=\"markdown-body\">",
	}

	for _, check := range checks {
		if !strings.Contains(result, check) {
			t.Errorf("expected %q in output", check)
		}
	}
}

func TestGenerate_DarkModeSupport(t *testing.T) {
	result := Generate("Test", "<p>Content</p>")

	if !strings.Contains(result, "prefers-color-scheme: dark") {
		t.Error("expected dark mode media query in output")
	}

	if !strings.Contains(result, "#0d1117") {
		t.Error("expected dark mode background color in output")
	}
}

func TestGenerate_ResponsiveLayout(t *testing.T) {
	result := Generate("Test", "<p>Content</p>")

	if !strings.Contains(result, "max-width: 980px") {
		t.Error("expected max-width constraint in output")
	}

	if !strings.Contains(result, "margin: 0 auto") {
		t.Error("expected centered layout in output")
	}
}

func TestGenerate_ContainsChromaCSS(t *testing.T) {
	result := Generate("Test", "<p>Content</p>")

	// Check for chroma syntax highlighting class styling
	if !strings.Contains(result, ".hl-k") {
		t.Error("expected Chroma keyword class styling in output")
	}

	if !strings.Contains(result, ".hl-s") {
		t.Error("expected Chroma string class styling in output")
	}

	if !strings.Contains(result, "color-prettylights-syntax") {
		t.Error("expected CSS to reference PrettyLights variables for theming")
	}
}

func TestGenerate_ContainsCommentsFeature(t *testing.T) {
	result := Generate("Test", "<p>Content</p>")

	// Check for comment button
	if !strings.Contains(result, "comment-btn") {
		t.Error("expected comment button in output")
	}

	// Check for comments panel
	if !strings.Contains(result, "comments-panel") {
		t.Error("expected comments panel in output")
	}

	// Check for copy comments button
	if !strings.Contains(result, "copy-comments-btn") {
		t.Error("expected copy comments button in output")
	}

	// Check for comment highlight CSS
	if !strings.Contains(result, "comment-highlight") {
		t.Error("expected comment highlight class in output")
	}
}

func TestGenerate_CommentsJavaScript(t *testing.T) {
	result := Generate("Test", "<p>Content</p>")

	// Check for key JS functions
	checks := []string{
		"getSelection",
		"highlightRange",
		"copyCommentsAsMarkdown",
		"openCommentsPanel",
		"closeCommentsPanel",
	}

	for _, check := range checks {
		if !strings.Contains(result, check) {
			t.Errorf("expected %q in JavaScript for comments feature", check)
		}
	}
}

func TestGenerate_CommentsCSS(t *testing.T) {
	result := Generate("Test", "<p>Content</p>")

	// Check for comment-related CSS
	checks := []string{
		".comment-btn",
		".comment-highlight",
		".comments-panel",
		".comment-entry",
		".comment-input-form",
		".copy-comments-btn",
	}

	for _, check := range checks {
		if !strings.Contains(result, check) {
			t.Errorf("expected CSS class %q in output", check)
		}
	}
}

func TestGenerate_CommentsHTML(t *testing.T) {
	result := Generate("Test", "<p>Content</p>")

	// Check for comment HTML elements
	checks := []string{
		`<button class="comment-btn"`,
		`<aside class="comments-panel"`,
		`<button class="copy-comments-btn"`,
		`class="comment-input-textarea"`,
		`class="comment-save-btn"`,
		`class="comment-cancel-btn"`,
	}

	for _, check := range checks {
		if !strings.Contains(result, check) {
			t.Errorf("expected HTML element %q in output", check)
		}
	}
}

func TestGenerate_KeyboardShortcutsModal(t *testing.T) {
	result := Generate("Test", "<p>Content</p>")

	// Check for keyboard shortcuts modal HTML elements
	checks := []string{
		`class="shortcuts-modal-overlay"`,
		`class="shortcuts-modal"`,
		`class="shortcuts-modal-header"`,
		`class="shortcuts-modal-content"`,
		`class="shortcut-row"`,
		`class="shortcut-keys"`,
	}

	for _, check := range checks {
		if !strings.Contains(result, check) {
			t.Errorf("expected keyboard shortcuts modal element %q in output", check)
		}
	}
}

func TestGenerate_KeyboardShortcutsJavaScript(t *testing.T) {
	result := Generate("Test", "<p>Content</p>")

	// Check for keyboard shortcuts JS functions
	checks := []string{
		"openShortcutsModal",
		"closeShortcutsModal",
		"shortcutsOverlay",
		"shortcutsModal",
	}

	for _, check := range checks {
		if !strings.Contains(result, check) {
			t.Errorf("expected %q in JavaScript for keyboard shortcuts", check)
		}
	}
}

func TestGenerate_CommentsEmptyState(t *testing.T) {
	result := Generate("Test", "<p>Content</p>")

	// Check for empty state HTML elements
	checks := []string{
		`class="comments-empty"`,
		"No comments yet",
		"Select text and press",
	}

	for _, check := range checks {
		if !strings.Contains(result, check) {
			t.Errorf("expected empty state element %q in output", check)
		}
	}
}

func TestGenerate_MermaidScriptIncluded(t *testing.T) {
	result := Generate("Test", "<p>Content</p>")

	// Check that mermaid script is included
	if !strings.Contains(result, "language-mermaid") {
		t.Error("expected mermaid language class detection in output")
	}

	if !strings.Contains(result, "mermaid.esm.min.mjs") {
		t.Error("expected Mermaid.js CDN import in output")
	}
}

func TestGenerate_MermaidCSSIncluded(t *testing.T) {
	result := Generate("Test", "<p>Content</p>")

	// Check for mermaid CSS classes
	checks := []string{
		".mermaid-wrapper",
		".mermaid-rendered",
		".mermaid-source",
		".mermaid-error",
	}

	for _, check := range checks {
		if !strings.Contains(result, check) {
			t.Errorf("expected mermaid CSS class %q in output", check)
		}
	}
}

func TestGenerate_MermaidCSSThemeSupport(t *testing.T) {
	result := Generate("Test", "<p>Content</p>")

	// Check for mermaid dark mode CSS
	if !strings.Contains(result, ".mermaid-wrapper") {
		t.Error("expected mermaid wrapper class in CSS")
	}

	// The dark mode CSS for mermaid should be in the chroma.css
	if !strings.Contains(result, ".mermaid-error") {
		t.Error("expected mermaid error class in CSS")
	}
}

func TestGenerate_CopyButtonSkipsMermaid(t *testing.T) {
	result := Generate("Test", "<p>Content</p>")

	// Check that the copy button script has the mermaid skip logic
	if !strings.Contains(result, "language-mermaid") {
		t.Error("expected mermaid skip logic in copy button script")
	}

	// Check that the skip logic exists
	if !strings.Contains(result, "classList.contains('language-mermaid')") {
		t.Error("expected classList check for mermaid in copy button script")
	}
}

func TestGenerate_MermaidJavaScript(t *testing.T) {
	result := Generate("Test", "<p>Content</p>")

	// Check for key mermaid JS functions and features
	checks := []string{
		"isDarkMode",
		"mermaid-wrapper",
		"mermaid-rendered",
		"mermaid.render",
		"prefers-color-scheme",
	}

	for _, check := range checks {
		if !strings.Contains(result, check) {
			t.Errorf("expected %q in mermaid JavaScript", check)
		}
	}
}

func TestGenerateWithLiveReload_MermaidIncluded(t *testing.T) {
	result := GenerateWithLiveReload("Test", "<p>Content</p>", 8080)

	// Check that mermaid is included in live reload mode
	if !strings.Contains(result, "mermaid.esm.min.mjs") {
		t.Error("expected Mermaid.js CDN import in live reload output")
	}

	if !strings.Contains(result, ".mermaid-wrapper") {
		t.Error("expected mermaid CSS in live reload output")
	}
}

func TestGenerateMulti_MermaidIncluded(t *testing.T) {
	tree := &filetree.TreeNode{
		Name:  "root",
		IsDir: true,
		Children: []*filetree.TreeNode{
			{
				Name:  "test.md",
				IsDir: false,
				File: &filetree.FileEntry{
					ID:   "test-md",
					Name: "test.md",
					Path: "test.md",
				},
			},
		},
	}
	files := []filetree.FileEntry{
		{
			ID:      "test-md",
			Name:    "test.md",
			Path:    "test.md",
			Content: "<p>Content</p>",
		},
	}

	result := GenerateMulti("Test", tree, files)

	// Check that mermaid script is included
	if !strings.Contains(result, "mermaid.esm.min.mjs") {
		t.Error("expected Mermaid.js CDN import in multifile output")
	}

	// Check that mermaid CSS classes are included
	checks := []string{
		".mermaid-wrapper",
		".mermaid-rendered",
		".mermaid-source",
		".mermaid-error",
	}

	for _, check := range checks {
		if !strings.Contains(result, check) {
			t.Errorf("expected mermaid CSS class %q in multifile output", check)
		}
	}
}

func TestGenerateMulti_CopyButtonSkipsMermaid(t *testing.T) {
	tree := &filetree.TreeNode{
		Name:  "root",
		IsDir: true,
		Children: []*filetree.TreeNode{
			{
				Name:  "test.md",
				IsDir: false,
				File: &filetree.FileEntry{
					ID:   "test-md",
					Name: "test.md",
					Path: "test.md",
				},
			},
		},
	}
	files := []filetree.FileEntry{
		{
			ID:      "test-md",
			Name:    "test.md",
			Path:    "test.md",
			Content: "<p>Content</p>",
		},
	}

	result := GenerateMulti("Test", tree, files)

	// Check that the copy button script has the mermaid skip logic
	if !strings.Contains(result, "classList.contains('language-mermaid')") {
		t.Error("expected classList check for mermaid in multifile copy button script")
	}
}

func TestGenerateMultiWithLiveReload_MermaidIncluded(t *testing.T) {
	tree := &filetree.TreeNode{
		Name:  "root",
		IsDir: true,
		Children: []*filetree.TreeNode{
			{
				Name:  "test.md",
				IsDir: false,
				File: &filetree.FileEntry{
					ID:   "test-md",
					Name: "test.md",
					Path: "test.md",
				},
			},
		},
	}
	files := []filetree.FileEntry{
		{
			ID:      "test-md",
			Name:    "test.md",
			Path:    "test.md",
			Content: "<p>Content</p>",
		},
	}

	result := GenerateMultiWithLiveReload("Test", tree, files, 8080)

	// Check that mermaid is included in live reload mode
	if !strings.Contains(result, "mermaid.esm.min.mjs") {
		t.Error("expected Mermaid.js CDN import in multifile live reload output")
	}

	// Check that live reload script is also included
	if !strings.Contains(result, "WebSocket") {
		t.Error("expected WebSocket in multifile live reload output")
	}
}
