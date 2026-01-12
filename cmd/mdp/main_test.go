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

func TestRun_TooManyArgs(t *testing.T) {
	err := run([]string{"file1.md", "file2.md"})
	if err == nil {
		t.Error("expected error for too many arguments")
	}
	if !strings.Contains(err.Error(), "Usage") {
		t.Errorf("expected usage message, got: %v", err)
	}
}

func TestRun_InvalidExtension(t *testing.T) {
	err := run([]string{"file.txt"})
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
	if !strings.Contains(err.Error(), "reading file") {
		t.Errorf("expected file reading error, got: %v", err)
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
