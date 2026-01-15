package updater

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestLoadState_NonExistent(t *testing.T) {
	// Save original home and restore after test
	origHome := os.Getenv("HOME")
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", origHome)

	state, err := LoadState()
	if err != nil {
		t.Fatalf("LoadState() error = %v, want nil", err)
	}

	// Should return empty state
	if state.CurrentVersion != "" {
		t.Errorf("CurrentVersion = %q, want empty", state.CurrentVersion)
	}
	if state.LatestVersion != "" {
		t.Errorf("LatestVersion = %q, want empty", state.LatestVersion)
	}
}

func TestLoadState_Corrupted(t *testing.T) {
	origHome := os.Getenv("HOME")
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", origHome)

	// Create corrupted state file
	cacheDir := filepath.Join(tmpDir, ".cache", "mdp")
	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		t.Fatal(err)
	}
	statePath := filepath.Join(cacheDir, "state.json")
	if err := os.WriteFile(statePath, []byte("not valid json"), 0644); err != nil {
		t.Fatal(err)
	}

	state, err := LoadState()
	if err != nil {
		t.Fatalf("LoadState() error = %v, want nil for corrupted file", err)
	}

	// Should return empty state
	if state.CurrentVersion != "" {
		t.Errorf("CurrentVersion = %q, want empty for corrupted file", state.CurrentVersion)
	}
}

func TestSaveState_CreatesDirectory(t *testing.T) {
	origHome := os.Getenv("HOME")
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", origHome)

	state := &State{
		CurrentVersion: "1.0.0",
		LatestVersion:  "1.1.0",
		LastCheckTime:  time.Now(),
		InstallMethod:  "curl",
	}

	if err := SaveState(state); err != nil {
		t.Fatalf("SaveState() error = %v", err)
	}

	// Verify directory was created
	cacheDir := filepath.Join(tmpDir, ".cache", "mdp")
	if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
		t.Error("Cache directory was not created")
	}

	// Verify file was created
	statePath := filepath.Join(cacheDir, "state.json")
	if _, err := os.Stat(statePath); os.IsNotExist(err) {
		t.Error("State file was not created")
	}
}

func TestSaveState_RoundTrip(t *testing.T) {
	origHome := os.Getenv("HOME")
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", origHome)

	checkTime := time.Now().Truncate(time.Second)
	original := &State{
		CurrentVersion: "1.2.3",
		LatestVersion:  "1.3.0",
		LastCheckTime:  checkTime,
		InstallMethod:  "brew",
	}

	if err := SaveState(original); err != nil {
		t.Fatalf("SaveState() error = %v", err)
	}

	loaded, err := LoadState()
	if err != nil {
		t.Fatalf("LoadState() error = %v", err)
	}

	if loaded.CurrentVersion != original.CurrentVersion {
		t.Errorf("CurrentVersion = %q, want %q", loaded.CurrentVersion, original.CurrentVersion)
	}
	if loaded.LatestVersion != original.LatestVersion {
		t.Errorf("LatestVersion = %q, want %q", loaded.LatestVersion, original.LatestVersion)
	}
	if loaded.InstallMethod != original.InstallMethod {
		t.Errorf("InstallMethod = %q, want %q", loaded.InstallMethod, original.InstallMethod)
	}
	// Time comparison with tolerance for JSON serialization
	if loaded.LastCheckTime.Unix() != original.LastCheckTime.Unix() {
		t.Errorf("LastCheckTime = %v, want %v", loaded.LastCheckTime, original.LastCheckTime)
	}
}

func TestShouldCheckRemote(t *testing.T) {
	tests := []struct {
		name     string
		state    *State
		expected bool
	}{
		{
			name:     "never checked",
			state:    &State{},
			expected: true,
		},
		{
			name: "checked 1 hour ago",
			state: &State{
				LastCheckTime: time.Now().Add(-1 * time.Hour),
			},
			expected: false,
		},
		{
			name: "checked 23 hours ago",
			state: &State{
				LastCheckTime: time.Now().Add(-23 * time.Hour),
			},
			expected: false,
		},
		{
			name: "checked 25 hours ago",
			state: &State{
				LastCheckTime: time.Now().Add(-25 * time.Hour),
			},
			expected: true,
		},
		{
			name: "checked 48 hours ago",
			state: &State{
				LastCheckTime: time.Now().Add(-48 * time.Hour),
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ShouldCheckRemote(tt.state)
			if got != tt.expected {
				t.Errorf("ShouldCheckRemote() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestGetCacheDir(t *testing.T) {
	origHome := os.Getenv("HOME")
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", origHome)

	cacheDir, err := GetCacheDir()
	if err != nil {
		t.Fatalf("GetCacheDir() error = %v", err)
	}

	expected := filepath.Join(tmpDir, ".cache", "mdp")
	if cacheDir != expected {
		t.Errorf("GetCacheDir() = %q, want %q", cacheDir, expected)
	}
}
