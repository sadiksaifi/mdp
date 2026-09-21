package template

import (
	"strings"
	"testing"

	"mdp/internal/filetree"
)

func TestGenerate_ThemeToggle(t *testing.T) {
	result := Generate("Test", "<p>Content</p>")

	for _, want := range []string{
		"topbar-theme-btn",
		"mdp-theme",
		"data-theme",
		`html[data-theme="dark"] .markdown-body`,
	} {
		if !strings.Contains(result, want) {
			t.Errorf("expected theme toggle %q in single-file output", want)
		}
	}
}

func TestGenerate_ThemeDefersToSystemUntilManualChoice(t *testing.T) {
	result := Generate("Test", "<p>Content</p>")

	// The OS-derived default must be applied without persisting, so the
	// matchMedia follower keeps working until the user clicks the toggle.
	for _, want := range []string{
		"readSavedTheme",
		"applyTheme(readSavedTheme()",
		", false)",
		", true)",
	} {
		if !strings.Contains(result, want) {
			t.Errorf("expected non-persisting theme init %q in single-file output", want)
		}
	}
}

func TestGenerateMulti_ThemeToggle(t *testing.T) {
	tree := &filetree.TreeNode{
		Name:  "root",
		IsDir: true,
		Children: []*filetree.TreeNode{
			{
				Name:  "a.md",
				IsDir: false,
				File:  &filetree.FileEntry{ID: "a-md", Name: "a.md", Path: "a.md"},
			},
		},
	}
	files := []filetree.FileEntry{{ID: "a-md", Name: "a.md", Path: "a.md", Content: "<p>hi</p>"}}
	result := GenerateMulti("Test", tree, files)

	for _, want := range []string{
		"topbar-theme-btn",
		"mdp-theme",
		"data-theme",
	} {
		if !strings.Contains(result, want) {
			t.Errorf("expected theme toggle %q in multi-file output", want)
		}
	}
}
