package updater

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectInstallMethodFromPath(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected InstallMethod
	}{
		{
			name:     "homebrew cellar path",
			path:     "/opt/homebrew/Cellar/mdp/1.0.0/bin/mdp",
			expected: InstallMethodBrew,
		},
		{
			name:     "homebrew path lowercase",
			path:     "/usr/local/homebrew/bin/mdp",
			expected: InstallMethodBrew,
		},
		{
			name:     "linuxbrew path",
			path:     "/home/user/.linuxbrew/bin/mdp",
			expected: InstallMethodBrew,
		},
		{
			name:     "usr local bin (source install)",
			path:     "/usr/local/bin/mdp",
			expected: InstallMethodSource,
		},
		{
			name:     "random path",
			path:     "/some/random/path/mdp",
			expected: InstallMethodUnknown,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DetectInstallMethodFromPath(tt.path)
			if got != tt.expected {
				t.Errorf("DetectInstallMethodFromPath(%q) = %v, want %v", tt.path, got, tt.expected)
			}
		})
	}
}

func TestDetectInstallMethodFromPath_LocalBin(t *testing.T) {
	origHome := os.Getenv("HOME")
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", origHome)

	localBinPath := filepath.Join(tmpDir, ".local", "bin", "mdp")
	got := DetectInstallMethodFromPath(localBinPath)
	if got != InstallMethodCurl {
		t.Errorf("DetectInstallMethodFromPath(%q) = %v, want %v", localBinPath, got, InstallMethodCurl)
	}
}

func TestDetectInstallMethodFromPath_GOPATH(t *testing.T) {
	origHome := os.Getenv("HOME")
	origGopath := os.Getenv("GOPATH")
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	os.Setenv("GOPATH", filepath.Join(tmpDir, "go"))
	defer func() {
		os.Setenv("HOME", origHome)
		os.Setenv("GOPATH", origGopath)
	}()

	gopathBin := filepath.Join(tmpDir, "go", "bin", "mdp")
	got := DetectInstallMethodFromPath(gopathBin)
	if got != InstallMethodSource {
		t.Errorf("DetectInstallMethodFromPath(%q) = %v, want %v", gopathBin, got, InstallMethodSource)
	}
}

func TestDetectInstallMethodFromPath_MarkerFile(t *testing.T) {
	origHome := os.Getenv("HOME")
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", origHome)

	// Create marker file
	markerDir := filepath.Join(tmpDir, ".local", "share", "mdp")
	if err := os.MkdirAll(markerDir, 0755); err != nil {
		t.Fatal(err)
	}
	markerPath := filepath.Join(markerDir, ".curl-installed")
	if err := os.WriteFile(markerPath, []byte("curl-installed\n"), 0644); err != nil {
		t.Fatal(err)
	}

	// Even with a random path, marker file should be detected
	got := DetectInstallMethodFromPath("/some/random/path/mdp")
	if got != InstallMethodCurl {
		t.Errorf("DetectInstallMethodFromPath with marker file = %v, want %v", got, InstallMethodCurl)
	}
}

func TestGetMarkerFilePath(t *testing.T) {
	origHome := os.Getenv("HOME")
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", origHome)

	markerPath, err := GetMarkerFilePath()
	if err != nil {
		t.Fatalf("GetMarkerFilePath() error = %v", err)
	}

	expected := filepath.Join(tmpDir, ".local", "share", "mdp", ".curl-installed")
	if markerPath != expected {
		t.Errorf("GetMarkerFilePath() = %q, want %q", markerPath, expected)
	}
}

func TestCreateMarkerFile(t *testing.T) {
	origHome := os.Getenv("HOME")
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", origHome)

	if err := CreateMarkerFile(); err != nil {
		t.Fatalf("CreateMarkerFile() error = %v", err)
	}

	markerPath := filepath.Join(tmpDir, ".local", "share", "mdp", ".curl-installed")
	if _, err := os.Stat(markerPath); os.IsNotExist(err) {
		t.Error("Marker file was not created")
	}
}

func TestInstallMethod_String(t *testing.T) {
	tests := []struct {
		method   InstallMethod
		expected string
	}{
		{InstallMethodBrew, "brew"},
		{InstallMethodCurl, "curl"},
		{InstallMethodSource, "source"},
		{InstallMethodUnknown, "unknown"},
	}

	for _, tt := range tests {
		t.Run(string(tt.method), func(t *testing.T) {
			if string(tt.method) != tt.expected {
				t.Errorf("InstallMethod = %q, want %q", string(tt.method), tt.expected)
			}
		})
	}
}
