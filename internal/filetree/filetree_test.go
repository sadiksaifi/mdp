package filetree

import (
	"testing"
)

func TestBuildTree_FlatFiles(t *testing.T) {
	files := []FileEntry{
		{ID: "a-md", Name: "a", RelPath: "a.md"},
		{ID: "b-md", Name: "b", RelPath: "b.md"},
	}

	tree := BuildTree(files)

	if len(tree.Children) != 2 {
		t.Errorf("expected 2 children, got %d", len(tree.Children))
	}

	// Should be sorted alphabetically
	if tree.Children[0].Name != "a" {
		t.Errorf("expected first child to be 'a', got %q", tree.Children[0].Name)
	}
	if tree.Children[1].Name != "b" {
		t.Errorf("expected second child to be 'b', got %q", tree.Children[1].Name)
	}
}

func TestBuildTree_NestedDirectories(t *testing.T) {
	files := []FileEntry{
		{ID: "docs-guide-md", Name: "guide", RelPath: "docs/guide.md"},
		{ID: "docs-api-md", Name: "api", RelPath: "docs/api.md"},
		{ID: "readme-md", Name: "README", RelPath: "README.md"},
	}

	tree := BuildTree(files)

	// Should have: docs/ directory and README.md at root
	if len(tree.Children) != 2 {
		t.Errorf("expected 2 top-level children (docs + README), got %d", len(tree.Children))
	}

	// Directories should come first
	if !tree.Children[0].IsDir {
		t.Error("expected first child to be directory")
	}
	if tree.Children[0].Name != "docs" {
		t.Errorf("expected first child to be 'docs', got %q", tree.Children[0].Name)
	}

	// Check docs directory has 2 children
	docsDir := tree.Children[0]
	if len(docsDir.Children) != 2 {
		t.Errorf("expected docs to have 2 children, got %d", len(docsDir.Children))
	}
}

func TestBuildTree_DeeplyNested(t *testing.T) {
	files := []FileEntry{
		{ID: "a-b-c-d-md", Name: "d", RelPath: "a/b/c/d.md"},
	}

	tree := BuildTree(files)

	// Navigate down the tree
	current := tree
	expectedDirs := []string{"a", "b", "c"}
	for _, dir := range expectedDirs {
		if len(current.Children) != 1 {
			t.Fatalf("expected 1 child at %s level", dir)
		}
		if current.Children[0].Name != dir {
			t.Errorf("expected %q, got %q", dir, current.Children[0].Name)
		}
		current = current.Children[0]
	}

	// Final level should have the file
	if len(current.Children) != 1 {
		t.Fatal("expected 1 file at deepest level")
	}
	if current.Children[0].Name != "d" {
		t.Errorf("expected file 'd', got %q", current.Children[0].Name)
	}
}

func TestBuildTree_Sorting(t *testing.T) {
	files := []FileEntry{
		{ID: "z-md", Name: "z", RelPath: "z.md"},
		{ID: "a-md", Name: "a", RelPath: "a.md"},
		{ID: "dir-b-md", Name: "b", RelPath: "dir/b.md"},
		{ID: "m-md", Name: "m", RelPath: "m.md"},
	}

	tree := BuildTree(files)

	// Should be: dir (directory first), then a, m, z (files alphabetically)
	if len(tree.Children) != 4 {
		t.Fatalf("expected 4 children, got %d", len(tree.Children))
	}

	if tree.Children[0].Name != "dir" || !tree.Children[0].IsDir {
		t.Errorf("expected first child to be directory 'dir'")
	}
	if tree.Children[1].Name != "a" {
		t.Errorf("expected second child to be 'a', got %q", tree.Children[1].Name)
	}
	if tree.Children[2].Name != "m" {
		t.Errorf("expected third child to be 'm', got %q", tree.Children[2].Name)
	}
	if tree.Children[3].Name != "z" {
		t.Errorf("expected fourth child to be 'z', got %q", tree.Children[3].Name)
	}
}

func TestBuildTree_EmptyFiles(t *testing.T) {
	files := []FileEntry{}

	tree := BuildTree(files)

	if tree == nil {
		t.Fatal("expected non-nil tree")
	}
	if len(tree.Children) != 0 {
		t.Errorf("expected 0 children, got %d", len(tree.Children))
	}
}

func TestBuildTree_FileEntryPreserved(t *testing.T) {
	entry := FileEntry{
		ID:      "test-id",
		Name:    "test",
		RelPath: "test.md",
		Path:    "/full/path/test.md",
		Content: "<h1>Test</h1>",
	}
	files := []FileEntry{entry}

	tree := BuildTree(files)

	if len(tree.Children) != 1 {
		t.Fatal("expected 1 child")
	}

	node := tree.Children[0]
	if node.File == nil {
		t.Fatal("expected File to be set")
	}
	if node.File.ID != "test-id" {
		t.Errorf("expected ID 'test-id', got %q", node.File.ID)
	}
	if node.File.Content != "<h1>Test</h1>" {
		t.Errorf("expected content to be preserved")
	}
}
