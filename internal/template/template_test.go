package template

import (
	"strings"
	"testing"
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
