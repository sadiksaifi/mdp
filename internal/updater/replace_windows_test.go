//go:build windows

package updater

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestWindowsReplacementRejectsAnotherRunningInstance(t *testing.T) {
	tmpDir := t.TempDir()
	targetPath := filepath.Join(tmpDir, "mdp.exe")
	copyRunningTestExecutable(t, targetPath)
	startExecutableHolder(t, targetPath)

	replacementPath := filepath.Join(tmpDir, "replacement.exe")
	if err := os.WriteFile(replacementPath, []byte("replacement"), 0755); err != nil {
		t.Fatal(err)
	}

	deferred, err := replaceBinary(replacementPath, targetPath, "v1.2.3")
	if err == nil {
		t.Fatal("replaceBinary() error = nil, want another running process error")
	}
	if deferred {
		t.Error("replaceBinary() deferred = true, want false")
	}
	if !strings.Contains(err.Error(), "another mdp process") {
		t.Fatalf("replaceBinary() error = %q, want another mdp process error", err)
	}
}

func TestWindowsReplacementProcessHolder(t *testing.T) {
	if os.Getenv("MDP_TEST_EXECUTABLE_HOLDER") != "1" {
		return
	}

	readyPath := os.Getenv("MDP_TEST_EXECUTABLE_HOLDER_READY")
	if err := os.WriteFile(readyPath, []byte("ready"), 0600); err != nil {
		t.Fatal(err)
	}
	if _, err := io.Copy(io.Discard, os.Stdin); err != nil {
		t.Fatal(err)
	}
}

func copyRunningTestExecutable(t *testing.T, targetPath string) {
	t.Helper()

	executable, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}
	if err := copyExecutable(executable, targetPath, 0755); err != nil {
		t.Fatal(err)
	}
}

func startExecutableHolder(t *testing.T, targetPath string) {
	t.Helper()

	readyPath := filepath.Join(t.TempDir(), "ready")
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	var output bytes.Buffer
	cmd := exec.Command(targetPath, "-test.run=^TestWindowsReplacementProcessHolder$")
	cmd.Env = append(os.Environ(),
		"MDP_TEST_EXECUTABLE_HOLDER=1",
		"MDP_TEST_EXECUTABLE_HOLDER_READY="+readyPath,
	)
	cmd.Stdin = reader
	cmd.Stdout = &output
	cmd.Stderr = &output
	if err := cmd.Start(); err != nil {
		_ = reader.Close()
		_ = writer.Close()
		t.Fatal(err)
	}
	_ = reader.Close()

	t.Cleanup(func() {
		_ = writer.Close()
		if err := cmd.Wait(); err != nil {
			t.Errorf("executable holder failed: %v\n%s", err, output.String())
		}
	})

	deadline := time.Now().Add(5 * time.Second)
	for {
		if _, err := os.Stat(readyPath); err == nil {
			return
		} else if !os.IsNotExist(err) {
			t.Fatal(err)
		}
		if time.Now().After(deadline) {
			t.Fatalf("timed out waiting for executable holder\n%s", output.String())
		}
		time.Sleep(25 * time.Millisecond)
	}
}

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
