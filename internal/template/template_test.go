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
