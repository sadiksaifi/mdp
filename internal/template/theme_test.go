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

func TestGenerate_MobileThemeToggle(t *testing.T) {
	result := Generate("Test", "<p>Content</p>")

	if !strings.Contains(result, `<button class="mobile-theme-btn theme-toggle-btn"`) {
		t.Fatal("expected an accessible mobile theme toggle in single-file output")
	}
}

func TestGenerate_MermaidFollowsManualTheme(t *testing.T) {
	result := Generate("Test", "<p>Content</p>")

	// Diagrams must honor the manual toggle, not just the OS preference.
	for _, want := range []string{
		"manualTheme",
		"mdp-theme-change",
	} {
		if !strings.Contains(result, want) {
			t.Errorf("expected mermaid manual-theme hook %q in single-file output", want)
		}
	}
}

func TestGenerateWithLiveReload_ThemeToggle(t *testing.T) {
	result := GenerateWithLiveReload("Test", "<p>Content</p>", 8080)

	for _, want := range []string{
		"topbar-theme-btn",
		"readSavedTheme",
		"mdp-theme-change",
		`html[data-theme="dark"] .markdown-body`,
	} {
		if !strings.Contains(result, want) {
			t.Errorf("expected theme toggle %q in single-file live-reload output", want)
		}
	}
}

func TestGenerateMulti_ThemeScriptNotNested(t *testing.T) {
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

	// The sidebar <script> block must not contain a nested <script> tag:
	// browsers end the outer script at the first </script>, which kills
	// both the sidebar JS and the theme toggle.
	openIdx := strings.Index(result, "var sidebar = document.querySelector")
	if openIdx == -1 {
		t.Fatal("expected sidebar JS in multi-file output")
	}
	blockOpen := strings.LastIndex(result[:openIdx], "<script")
	blockClose := strings.Index(result[blockOpen:], "</script>")
	block := result[blockOpen : blockOpen+blockClose]
	if strings.Contains(block[len("<script"):], "<script") {
		t.Error("nested <script> tag inside sidebar script block; theme toggle would be dead")
	}
	if !strings.Contains(block, "mdp-theme") {
		t.Error("expected theme toggle code inside sidebar script block")
	}
}
func TestGenerateMulti_MobileThemeToggle(t *testing.T) {
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

	mobileHeaderStart := strings.Index(result, `<div class="floating-buttons">`)
	mobileHeaderEnd := strings.Index(result, `<div class="sidebar-overlay">`)
	if mobileHeaderStart == -1 || mobileHeaderEnd == -1 || mobileHeaderEnd <= mobileHeaderStart {
		t.Fatal("expected mobile header in multi-file output")
	}
	mobileHeader := result[mobileHeaderStart:mobileHeaderEnd]
	if !strings.Contains(mobileHeader, `<button class="mobile-theme-btn theme-toggle-btn"`) {
		t.Fatal("expected an accessible theme toggle in the multi-file mobile header")
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

func TestGenerateMultiWithLiveReload_ThemeToggle(t *testing.T) {
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
	result := GenerateMultiWithLiveReload("Test", tree, files, 8080)

	for _, want := range []string{
		"topbar-theme-btn",
		"readSavedTheme",
		"mdp-theme-change",
	} {
		if !strings.Contains(result, want) {
			t.Errorf("expected theme toggle %q in multi-file live-reload output", want)
		}
	}
}

func TestGenerate_ThemeSensitiveControlsUseSharedVariables(t *testing.T) {
	result := Generate("Test", "<p>Content</p>")

	for _, want := range []string{
		`--theme-light-github-bg: #f6f8fa`,
		`--theme-dark-github-bg: #21262d`,
		`background: var(--github-link-bg)`,
		`background-color: var(--comment-highlight-bg)`,
	} {
		if !strings.Contains(result, want) {
			t.Errorf("expected shared theme variable %q in single-file output", want)
		}
	}
	if strings.Contains(result, `html[data-theme="dark"] .comments-panel`) {
		t.Error("manual theme should change shared variables instead of duplicating component selectors")
	}
}

func TestGenerate_CommentsUseThemeVariables(t *testing.T) {
	result := Generate("Test", "<p>Content</p>")

	for _, want := range []string{
		`background: var(--panel-bg)`,
		`background: var(--input-bg)`,
		`background: var(--accent-wash)`,
	} {
		if !strings.Contains(result, want) {
			t.Errorf("expected variable-driven comment style %q in single-file output", want)
		}
	}
}

func TestGenerateMulti_CommentsUseThemeVariables(t *testing.T) {
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
		`background: var(--sidebar-bg)`,
		`background: var(--input-bg)`,
		`background: var(--accent-wash)`,
	} {
		if !strings.Contains(result, want) {
			t.Errorf("expected variable-driven comment style %q in multi-file output", want)
		}
	}
}
