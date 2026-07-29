//go:build windows

package updater

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWindowsReplacementInstallsNewBinary(t *testing.T) {
	tmpDir := t.TempDir()
	helperPath := filepath.Join(tmpDir, "mdp.exe.update.exe")
	targetPath := filepath.Join(tmpDir, "mdp.exe")

	if err := os.WriteFile(helperPath, []byte("new binary"), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(targetPath, []byte("old binary"), 0755); err != nil {
		t.Fatal(err)
	}

	if err := installWindowsReplacement(helperPath, targetPath); err != nil {
		t.Fatalf("installWindowsReplacement() error = %v", err)
	}

	got, err := os.ReadFile(targetPath)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != "new binary" {
		t.Errorf("installed binary = %q, want %q", got, "new binary")
	}
	if _, err := os.Stat(targetPath + ".old"); !os.IsNotExist(err) {
		t.Errorf("backup file still exists, stat error = %v", err)
	}
}
