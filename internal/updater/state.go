package updater

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

// State represents the cached update state stored in ~/.cache/mdp/state.json
type State struct {
	CurrentVersion string    `json:"current_version"`
	LatestVersion  string    `json:"latest_version"`
	LastCheckTime  time.Time `json:"last_check_time"`
	InstallMethod  string    `json:"install_method"`
}

// GetCacheDir returns the path to the mdp cache directory (~/.cache/mdp/).
func GetCacheDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".cache", "mdp"), nil
}

// GetStatePath returns the path to the state file.
func GetStatePath() (string, error) {
	cacheDir, err := GetCacheDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(cacheDir, "state.json"), nil
}

// LoadState reads the state from the cache file.
// Returns an empty state if the file doesn't exist or is corrupted.
func LoadState() (*State, error) {
	statePath, err := GetStatePath()
	if err != nil {
		return &State{}, nil
	}

	data, err := os.ReadFile(statePath)
	if os.IsNotExist(err) {
		return &State{}, nil
	}
	if err != nil {
		return nil, err
	}

	var state State
	if err := json.Unmarshal(data, &state); err != nil {
		// Corrupted file, return empty state
		return &State{}, nil
	}

	return &state, nil
}

// SaveState writes the state to the cache file.
// Creates the cache directory if it doesn't exist.
func SaveState(state *State) error {
	cacheDir, err := GetCacheDir()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}

	statePath := filepath.Join(cacheDir, "state.json")
	return os.WriteFile(statePath, data, 0644)
}

// ShouldCheckRemote returns true if we should check for updates from GitHub.
// Returns true if last check was more than 24 hours ago or never checked.
func ShouldCheckRemote(state *State) bool {
	if state.LastCheckTime.IsZero() {
		return true
	}
	return time.Since(state.LastCheckTime) > 24*time.Hour
}
