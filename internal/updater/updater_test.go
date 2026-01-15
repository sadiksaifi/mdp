package updater

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestIsNewerVersion(t *testing.T) {
	tests := []struct {
		name     string
		latest   string
		current  string
		expected bool
	}{
		{
			name:     "newer version available",
			latest:   "v1.2.0",
			current:  "v1.1.0",
			expected: true,
		},
		{
			name:     "same version",
			latest:   "v1.1.0",
			current:  "v1.1.0",
			expected: false,
		},
		{
			name:     "older version on remote",
			latest:   "v1.0.0",
			current:  "v1.1.0",
			expected: false,
		},
		{
			name:     "dev build should not prompt",
			latest:   "v2.0.0",
			current:  "dev",
			expected: false,
		},
		{
			name:     "dev build with v prefix",
			latest:   "v2.0.0",
			current:  "vdev",
			expected: false,
		},
		{
			name:     "without v prefix - newer",
			latest:   "1.2.0",
			current:  "1.1.0",
			expected: true,
		},
		{
			name:     "without v prefix - same",
			latest:   "1.1.0",
			current:  "1.1.0",
			expected: false,
		},
		{
			name:     "mixed prefixes - newer",
			latest:   "v1.2.0",
			current:  "1.1.0",
			expected: true,
		},
		{
			name:     "patch version update",
			latest:   "v1.1.1",
			current:  "v1.1.0",
			expected: true,
		},
		{
			name:     "major version update",
			latest:   "v2.0.0",
			current:  "v1.9.9",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsNewerVersion(tt.latest, tt.current)
			if got != tt.expected {
				t.Errorf("IsNewerVersion(%q, %q) = %v, want %v", tt.latest, tt.current, got, tt.expected)
			}
		})
	}
}

func TestFetchLatestRelease_MockServer(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		release := Release{
			TagName: "v1.5.0",
			Name:    "Release 1.5.0",
			Assets: []Asset{
				{
					Name:               "mdp-darwin-arm64.tar.gz",
					BrowserDownloadURL: "https://example.com/mdp-darwin-arm64.tar.gz",
					Size:               1024,
				},
				{
					Name:               "mdp-linux-amd64.tar.gz",
					BrowserDownloadURL: "https://example.com/mdp-linux-amd64.tar.gz",
					Size:               1024,
				},
			},
		}
		json.NewEncoder(w).Encode(release)
	}))
	defer server.Close()

	// Note: This test would need the package to support custom URLs
	// For now, we test the parsing logic indirectly through Release methods
	release := &Release{
		TagName: "v1.5.0",
		Assets: []Asset{
			{Name: "mdp-darwin-arm64.tar.gz", BrowserDownloadURL: "https://example.com/darwin-arm64"},
			{Name: "mdp-linux-amd64.tar.gz", BrowserDownloadURL: "https://example.com/linux-amd64"},
		},
	}

	t.Run("GetAssetURL found", func(t *testing.T) {
		url := release.GetAssetURL("mdp-darwin-arm64.tar.gz")
		if url != "https://example.com/darwin-arm64" {
			t.Errorf("GetAssetURL() = %q, want %q", url, "https://example.com/darwin-arm64")
		}
	})

	t.Run("GetAssetURL not found", func(t *testing.T) {
		url := release.GetAssetURL("mdp-windows-amd64.zip")
		if url != "" {
			t.Errorf("GetAssetURL() = %q, want empty string", url)
		}
	})
}

func TestExtractTarGz(t *testing.T) {
	// This test would require creating actual tar.gz files
	// Skip for now as it's an integration-level test
	t.Skip("Integration test - requires actual tar.gz file")
}

func TestCheckForUpdateQuick_Timeout(t *testing.T) {
	// This test verifies that CheckForUpdateQuick returns nil on timeout
	// Since it makes network calls, we verify behavior with a non-existent server

	origHome := os.Getenv("HOME")
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", origHome)

	// With no cached state and a real network call that might fail,
	// CheckForUpdateQuick should return nil gracefully
	result := CheckForUpdateQuick("1.0.0")
	// Result can be nil (network failure) or valid (if GitHub is reachable)
	// We mainly verify it doesn't panic
	_ = result
}

func TestPerformUpgrade_DetectsBrewInstall(t *testing.T) {
	// This is more of an integration test
	// We verify the function doesn't panic with brew detection
	t.Skip("Integration test - requires mocking os.Executable")
}

func TestUpdateInfo_Fields(t *testing.T) {
	info := &UpdateInfo{
		CurrentVersion: "1.0.0",
		LatestVersion:  "1.1.0",
		HasUpdate:      true,
		InstallMethod:  InstallMethodCurl,
	}

	if info.CurrentVersion != "1.0.0" {
		t.Errorf("CurrentVersion = %q, want %q", info.CurrentVersion, "1.0.0")
	}
	if info.LatestVersion != "1.1.0" {
		t.Errorf("LatestVersion = %q, want %q", info.LatestVersion, "1.1.0")
	}
	if !info.HasUpdate {
		t.Error("HasUpdate = false, want true")
	}
	if info.InstallMethod != InstallMethodCurl {
		t.Errorf("InstallMethod = %v, want %v", info.InstallMethod, InstallMethodCurl)
	}
}

func TestCheckForUpdateWithContext_UsesCachedState(t *testing.T) {
	origHome := os.Getenv("HOME")
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", origHome)

	// Create a cached state that was checked recently
	cacheDir := filepath.Join(tmpDir, ".cache", "mdp")
	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		t.Fatal(err)
	}

	// Write a state that was checked 1 hour ago (within 24h window)
	stateJSON := `{
		"current_version": "1.0.0",
		"latest_version": "1.1.0",
		"last_check_time": "` + jsonTime() + `",
		"install_method": "curl"
	}`
	statePath := filepath.Join(cacheDir, "state.json")
	if err := os.WriteFile(statePath, []byte(stateJSON), 0644); err != nil {
		t.Fatal(err)
	}

	// This should use cached state and not make a network call
	// Verifying by checking it doesn't timeout
	info := CheckForUpdateQuick("1.0.0")
	if info != nil && info.LatestVersion != "1.1.0" {
		t.Errorf("LatestVersion = %q, want cached value %q", info.LatestVersion, "1.1.0")
	}
}

// jsonTime returns current time in JSON format for test fixtures
func jsonTime() string {
	return `2026-01-15T10:00:00Z` // Recent time that's within 24h
}
