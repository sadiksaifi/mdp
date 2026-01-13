package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRun_NoArgs(t *testing.T) {
	err := run([]string{})
	if err == nil {
		t.Error("expected error for no arguments")
	}
	if !strings.Contains(err.Error(), "Usage") {
		t.Errorf("expected usage message, got: %v", err)
	}
}

func TestRun_HelpFlag(t *testing.T) {
	err := run([]string{"--help"})
	if err != nil {
		t.Errorf("--help should not return error, got: %v", err)
	}
}

func TestRun_VersionFlag(t *testing.T) {
	err := run([]string{"--version"})
	if err != nil {
		t.Errorf("--version should not return error, got: %v", err)
	}
}

func TestRun_InvalidExtension(t *testing.T) {
	// Create a temp file with wrong extension
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "file.txt")
	err := os.WriteFile(tmpFile, []byte("test"), 0644)
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	err = run([]string{tmpFile})
	if err == nil {
		t.Error("expected error for invalid extension")
	}
	if !strings.Contains(err.Error(), ".md extension") {
		t.Errorf("expected extension error, got: %v", err)
	}
}

func TestRun_NonexistentFile(t *testing.T) {
	err := run([]string{"nonexistent.md"})
	if err == nil {
		t.Error("expected error for nonexistent file")
	}
	if !strings.Contains(err.Error(), "Error accessing") {
		t.Errorf("expected file access error, got: %v", err)
	}
}

func TestRun_ValidMarkdownFile(t *testing.T) {
	// Skip browser opening in tests
	if os.Getenv("CI") != "" {
		t.Skip("Skipping integration test in CI environment")
	}

	// Create a temporary markdown file
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.md")
	content := "# Test\n\nThis is a test."
	err := os.WriteFile(tmpFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	// Run the conversion
	err = run([]string{tmpFile})
	if err != nil {
		t.Errorf("run() returned error: %v", err)
	}

	// Verify output file was created
	outputPath := filepath.Join("/tmp", "mdpreview-test.html")
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		t.Error("expected output file to be created")
	}

	// Verify output content
	outputContent, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("failed to read output file: %v", err)
	}

	checks := []string{
		"<!DOCTYPE html>",
		"<title>test</title>",
		"<h1",
		"Test",
		"This is a test.",
	}

	for _, check := range checks {
		if !strings.Contains(string(outputContent), check) {
			t.Errorf("expected output to contain %q", check)
		}
	}
}

func TestRun_CaseInsensitiveExtension(t *testing.T) {
	// Skip browser opening in tests
	if os.Getenv("CI") != "" {
		t.Skip("Skipping integration test in CI environment")
	}

	// Create a temporary markdown file with uppercase extension
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "TEST.MD")
	content := "# Test"
	err := os.WriteFile(tmpFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	// Run should succeed with .MD extension
	err = run([]string{tmpFile})
	if err != nil {
		t.Errorf("run() should accept .MD extension, got error: %v", err)
	}
}

func TestRun_MultipleFiles(t *testing.T) {
	// Skip browser opening in tests
	if os.Getenv("CI") != "" {
		t.Skip("Skipping integration test in CI environment")
	}

	tmpDir := t.TempDir()

	// Create multiple test files
	files := []string{"one.md", "two.md", "three.md"}
	var paths []string
	for _, f := range files {
		path := filepath.Join(tmpDir, f)
		content := "# " + f + "\n\nContent for " + f
		err := os.WriteFile(path, []byte(content), 0644)
		if err != nil {
			t.Fatalf("failed to create temp file: %v", err)
		}
		paths = append(paths, path)
	}

	err := run(paths)
	if err != nil {
		t.Errorf("run() with multiple files failed: %v", err)
	}

	// Verify output file was created
	outputPath := filepath.Join("/tmp", "mdpreview-multi.html")
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		t.Error("expected output file to be created")
	}

	// Verify output content contains sidebar elements
	outputContent, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("failed to read output file: %v", err)
	}

	checks := []string{
		"<!DOCTYPE html>",
		"sidebar",
		"sidebar-open-btn",
		"file-tree",
		"one",
		"two",
		"three",
		"content-section",
	}

	for _, check := range checks {
		if !strings.Contains(string(outputContent), check) {
			t.Errorf("expected output to contain %q", check)
		}
	}
}

func TestRun_Directory(t *testing.T) {
	// Skip browser opening in tests
	if os.Getenv("CI") != "" {
		t.Skip("Skipping integration test in CI environment")
	}

	tmpDir := t.TempDir()

	// Create nested structure
	subDir := filepath.Join(tmpDir, "docs")
	err := os.MkdirAll(subDir, 0755)
	if err != nil {
		t.Fatalf("failed to create subdir: %v", err)
	}

	// Create files
	err = os.WriteFile(filepath.Join(tmpDir, "README.md"), []byte("# Root"), 0644)
	if err != nil {
		t.Fatalf("failed to create file: %v", err)
	}
	err = os.WriteFile(filepath.Join(subDir, "guide.md"), []byte("# Guide"), 0644)
	if err != nil {
		t.Fatalf("failed to create file: %v", err)
	}

	err = run([]string{tmpDir})
	if err != nil {
		t.Errorf("run() with directory failed: %v", err)
	}

	// Verify output content contains both files
	outputPath := filepath.Join("/tmp", "mdpreview-multi.html")
	outputContent, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("failed to read output file: %v", err)
	}

	checks := []string{
		"README",
		"guide",
		"docs",
	}

	for _, check := range checks {
		if !strings.Contains(string(outputContent), check) {
			t.Errorf("expected output to contain %q", check)
		}
	}
}

func TestRun_EmptyDirectory(t *testing.T) {
	tmpDir := t.TempDir()

	err := run([]string{tmpDir})
	if err == nil {
		t.Error("expected error for empty directory")
	}
	if !strings.Contains(err.Error(), "No markdown files") {
		t.Errorf("expected 'no markdown files' error, got: %v", err)
	}
}

func TestResolveFiles_SingleFile(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.md")
	err := os.WriteFile(tmpFile, []byte("# Test"), 0644)
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	files, err := resolveFiles([]string{tmpFile})
	if err != nil {
		t.Fatalf("resolveFiles failed: %v", err)
	}

	if len(files) != 1 {
		t.Errorf("expected 1 file, got %d", len(files))
	}
}

func TestResolveFiles_Directory(t *testing.T) {
	tmpDir := t.TempDir()

	// Create nested structure
	subDir := filepath.Join(tmpDir, "sub")
	err := os.MkdirAll(subDir, 0755)
	if err != nil {
		t.Fatalf("failed to create subdir: %v", err)
	}

	err = os.WriteFile(filepath.Join(tmpDir, "a.md"), []byte("# A"), 0644)
	if err != nil {
		t.Fatalf("failed to create file: %v", err)
	}
	err = os.WriteFile(filepath.Join(subDir, "b.md"), []byte("# B"), 0644)
	if err != nil {
		t.Fatalf("failed to create file: %v", err)
	}

	files, err := resolveFiles([]string{tmpDir})
	if err != nil {
		t.Fatalf("resolveFiles failed: %v", err)
	}

	if len(files) != 2 {
		t.Errorf("expected 2 files, got %d", len(files))
	}
}

func TestSanitizeID(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"README.md", "readme-md"},
		{"docs/guide.md", "docs-guide-md"},
		{"path with spaces/file.md", "path-with-spaces-file-md"},
		{"test", "test"},
	}

	for _, tc := range tests {
		result := sanitizeID(tc.input)
		if result != tc.expected {
			t.Errorf("sanitizeID(%q) = %q, want %q", tc.input, result, tc.expected)
		}
	}
}

func TestFindCommonBase(t *testing.T) {
	tests := []struct {
		paths    []string
		expected string
	}{
		{[]string{"/a/b/c.md"}, "/a/b"},
		{[]string{"/a/b/c.md", "/a/b/d.md"}, "/a/b"},
		{[]string{"/a/b/c.md", "/a/d/e.md"}, "/a"},
		{[]string{}, ""},
	}

	for _, tc := range tests {
		result := findCommonBase(tc.paths)
		if result != tc.expected {
			t.Errorf("findCommonBase(%v) = %q, want %q", tc.paths, result, tc.expected)
		}
	}
}

func TestRun_OutputFlag_SingleFile(t *testing.T) {
	tmpDir := t.TempDir()

	// Create a markdown file
	inputFile := filepath.Join(tmpDir, "test.md")
	err := os.WriteFile(inputFile, []byte("# Test\n\nHello world"), 0644)
	if err != nil {
		t.Fatalf("failed to create input file: %v", err)
	}

	// Specify custom output path
	outputFile := filepath.Join(tmpDir, "output.html")

	err = run([]string{"--output", outputFile, inputFile})
	if err != nil {
		t.Errorf("run() with --output flag failed: %v", err)
	}

	// Verify output file was created at specified path
	if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		t.Error("expected output file to be created at specified path")
	}

	// Verify content
	content, err := os.ReadFile(outputFile)
	if err != nil {
		t.Fatalf("failed to read output file: %v", err)
	}

	checks := []string{
		"<!DOCTYPE html>",
		"<title>test</title>",
		"Hello world",
	}

	for _, check := range checks {
		if !strings.Contains(string(content), check) {
			t.Errorf("expected output to contain %q", check)
		}
	}
}

func TestRun_OutputFlag_Shorthand(t *testing.T) {
	tmpDir := t.TempDir()

	// Create a markdown file
	inputFile := filepath.Join(tmpDir, "test.md")
	err := os.WriteFile(inputFile, []byte("# Shorthand Test"), 0644)
	if err != nil {
		t.Fatalf("failed to create input file: %v", err)
	}

	// Use -O shorthand
	outputFile := filepath.Join(tmpDir, "shorthand-output.html")

	err = run([]string{"-O", outputFile, inputFile})
	if err != nil {
		t.Errorf("run() with -O flag failed: %v", err)
	}

	// Verify output file was created
	if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		t.Error("expected output file to be created with -O shorthand")
	}
}

func TestRun_OutputFlag_MultipleFiles(t *testing.T) {
	tmpDir := t.TempDir()

	// Create multiple markdown files
	file1 := filepath.Join(tmpDir, "one.md")
	file2 := filepath.Join(tmpDir, "two.md")
	err := os.WriteFile(file1, []byte("# One"), 0644)
	if err != nil {
		t.Fatalf("failed to create file: %v", err)
	}
	err = os.WriteFile(file2, []byte("# Two"), 0644)
	if err != nil {
		t.Fatalf("failed to create file: %v", err)
	}

	// Specify custom output path
	outputFile := filepath.Join(tmpDir, "multi-output.html")

	err = run([]string{"--output", outputFile, file1, file2})
	if err != nil {
		t.Errorf("run() with --output and multiple files failed: %v", err)
	}

	// Verify output file was created
	if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		t.Error("expected output file to be created for multiple files")
	}

	// Verify content contains both files and sidebar
	content, err := os.ReadFile(outputFile)
	if err != nil {
		t.Fatalf("failed to read output file: %v", err)
	}

	checks := []string{
		"<!DOCTYPE html>",
		"sidebar",
		"One",
		"Two",
	}

	for _, check := range checks {
		if !strings.Contains(string(content), check) {
			t.Errorf("expected output to contain %q", check)
		}
	}
}

func TestRun_OutputFlag_WithServe_MutualExclusion(t *testing.T) {
	tmpDir := t.TempDir()

	// Create a markdown file
	inputFile := filepath.Join(tmpDir, "test.md")
	err := os.WriteFile(inputFile, []byte("# Test"), 0644)
	if err != nil {
		t.Fatalf("failed to create input file: %v", err)
	}

	outputFile := filepath.Join(tmpDir, "output.html")

	// Using both --output and --serve should fail
	err = run([]string{"--serve", "--output", outputFile, inputFile})
	if err == nil {
		t.Error("expected error when using --output with --serve")
	}
	if !strings.Contains(err.Error(), "cannot use --output with --serve") {
		t.Errorf("expected mutual exclusion error, got: %v", err)
	}
}

func TestRun_OutputFlag_Directory(t *testing.T) {
	tmpDir := t.TempDir()

	// Create a subdirectory with markdown files
	docsDir := filepath.Join(tmpDir, "docs")
	err := os.MkdirAll(docsDir, 0755)
	if err != nil {
		t.Fatalf("failed to create docs dir: %v", err)
	}

	err = os.WriteFile(filepath.Join(docsDir, "readme.md"), []byte("# Readme"), 0644)
	if err != nil {
		t.Fatalf("failed to create file: %v", err)
	}
	err = os.WriteFile(filepath.Join(docsDir, "guide.md"), []byte("# Guide"), 0644)
	if err != nil {
		t.Fatalf("failed to create file: %v", err)
	}

	// Output to custom path
	outputFile := filepath.Join(tmpDir, "docs-output.html")

	err = run([]string{"-O", outputFile, docsDir})
	if err != nil {
		t.Errorf("run() with -O and directory failed: %v", err)
	}

	// Verify output
	if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		t.Error("expected output file to be created for directory")
	}

	content, err := os.ReadFile(outputFile)
	if err != nil {
		t.Fatalf("failed to read output file: %v", err)
	}

	checks := []string{
		"Readme",
		"Guide",
		"sidebar",
	}

	for _, check := range checks {
		if !strings.Contains(string(content), check) {
			t.Errorf("expected output to contain %q", check)
		}
	}
}
