package browser

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestOpen_ValidFile(t *testing.T) {
	// Skip on CI environments where browser opening won't work
	if os.Getenv("CI") != "" {
		t.Skip("Skipping browser test in CI environment")
	}

	// Create a temporary HTML file
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.html")
	err := os.WriteFile(tmpFile, []byte("<html><body>Test</body></html>"), 0644)
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	// Test opening the file (will actually open in browser on local machine)
	// This test verifies the command doesn't error out
	err = Open(tmpFile)
	if err != nil {
		t.Errorf("Open() returned error: %v", err)
	}
}

func TestOpen_NonexistentFile(t *testing.T) {
	// Skip on CI environments
	if os.Getenv("CI") != "" {
		t.Skip("Skipping browser test in CI environment")
	}

	// Opening a nonexistent file should succeed on most platforms
	// (the browser will show an error, but the command itself succeeds)
	// This behavior is platform-specific
	err := Open("/nonexistent/path/that/does/not/exist.html")

	// On macOS, 'open' succeeds even for nonexistent files
	// On Linux, xdg-open may fail
	if runtime.GOOS != "darwin" && err == nil {
		// It's acceptable if it doesn't error on darwin
		t.Log("Note: Open() did not error on nonexistent file (platform-specific behavior)")
	}
}

func TestPlatformSupport(t *testing.T) {
	// Verify the current platform is supported
	supported := []string{"darwin", "linux", "windows"}
	currentOS := runtime.GOOS

	found := false
	for _, os := range supported {
		if currentOS == os {
			found = true
			break
		}
	}

	if !found {
		t.Logf("Current platform %s may not be fully supported", currentOS)
	}
}
