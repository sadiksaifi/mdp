package updater

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"golang.org/x/mod/semver"
)

// UpdateInfo contains information about available updates
type UpdateInfo struct {
	CurrentVersion string
	LatestVersion  string
	HasUpdate      bool
	InstallMethod  InstallMethod
}

// CheckForUpdateQuick performs a quick, non-blocking update check.
// Returns nil if check fails or times out (2 second timeout).
// This is designed to be called at the start of every mdp invocation.
func CheckForUpdateQuick(currentVersion string) *UpdateInfo {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	info, err := CheckForUpdateWithContext(ctx, currentVersion)
	if err != nil {
		return nil
	}
	return info
}

// CheckForUpdateWithContext checks for updates with context support.
func CheckForUpdateWithContext(ctx context.Context, currentVersion string) (*UpdateInfo, error) {
	state, err := LoadState()
	if err != nil {
		state = &State{}
	}

	installMethod := DetectInstallMethod()

	// Use cached information if we checked recently
	if !ShouldCheckRemote(state) && state.LatestVersion != "" {
		return &UpdateInfo{
			CurrentVersion: currentVersion,
			LatestVersion:  state.LatestVersion,
			HasUpdate:      IsNewerVersion(state.LatestVersion, currentVersion),
			InstallMethod:  installMethod,
		}, nil
	}

	// Query GitHub API
	release, err := FetchLatestReleaseWithContext(ctx)
	if err != nil {
		return nil, err
	}

	latestVersion := strings.TrimPrefix(release.TagName, "v")

	// Update state
	state.CurrentVersion = currentVersion
	state.LatestVersion = latestVersion
	state.LastCheckTime = time.Now()
	if state.InstallMethod == "" {
		state.InstallMethod = string(installMethod)
	}

	// Ignore save errors - non-critical
	_ = SaveState(state)

	return &UpdateInfo{
		CurrentVersion: currentVersion,
		LatestVersion:  latestVersion,
		HasUpdate:      IsNewerVersion(latestVersion, currentVersion),
		InstallMethod:  installMethod,
	}, nil
}

// IsNewerVersion compares two version strings and returns true if latest > current.
// Handles versions with or without 'v' prefix.
// Returns false for "dev" builds to avoid upgrade prompts during development.
func IsNewerVersion(latest, current string) bool {
	// Don't prompt dev builds to upgrade
	if current == "dev" || current == "vdev" {
		return false
	}

	// Ensure versions have 'v' prefix for semver
	if !strings.HasPrefix(latest, "v") {
		latest = "v" + latest
	}
	if !strings.HasPrefix(current, "v") {
		current = "v" + current
	}

	return semver.Compare(latest, current) > 0
}

// PerformUpgrade upgrades mdp to the latest version.
// Behavior depends on installation method:
// - brew: prints upgrade instructions
// - curl/PowerShell/unknown: downloads and replaces the binary
// - source: prints upgrade instructions
func PerformUpgrade(currentVersion string, force bool) error {
	method := DetectInstallMethod()

	switch method {
	case InstallMethodBrew:
		return handleBrewUpgrade()
	case InstallMethodSource:
		return handleSourceUpgrade()
	case InstallMethodCurl, InstallMethodUnknown:
		return handleDirectUpgrade(currentVersion, force)
	}

	return nil
}

func handleBrewUpgrade() error {
	fmt.Println("mdp was installed via Homebrew.")
	fmt.Println()
	fmt.Println("To upgrade, run:")
	fmt.Println()
	fmt.Println("  brew upgrade sadiksaifi/tap/mdp")
	fmt.Println()
	return nil
}

func handleSourceUpgrade() error {
	fmt.Println("mdp was installed from source.")
	fmt.Println()
	fmt.Println("To upgrade, run:")
	fmt.Println()
	fmt.Println("  cd /path/to/mdp && git pull && make install")
	fmt.Println()
	return nil
}

func handleDirectUpgrade(currentVersion string, force bool) error {
	fmt.Println("Checking for updates...")

	release, err := FetchLatestRelease()
	if err != nil {
		return fmt.Errorf("failed to check for updates: %w", err)
	}

	latestVersion := strings.TrimPrefix(release.TagName, "v")

	if !force && !IsNewerVersion(latestVersion, currentVersion) {
		fmt.Printf("mdp is up to date (version %s)\n", currentVersion)
		return nil
	}

	fmt.Printf("Upgrading mdp %s → %s\n", currentVersion, latestVersion)

	asset, err := releaseAssetFor(runtime.GOOS, runtime.GOARCH)
	if err != nil {
		return err
	}

	// Find matching asset
	downloadURL := release.GetAssetURL(asset.archiveName)
	if downloadURL == "" {
		return fmt.Errorf("no release found for %s/%s", runtime.GOOS, runtime.GOARCH)
	}

	// Download to temp directory
	tmpDir, err := os.MkdirTemp("", "mdp-upgrade-*")
	if err != nil {
		return fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	archivePath := filepath.Join(tmpDir, asset.archiveName)
	fmt.Println("Downloading...")
	if err := downloadFile(downloadURL, archivePath); err != nil {
		return fmt.Errorf("download failed: %w", err)
	}

	// Extract
	fmt.Println("Extracting...")
	if err := extractReleaseBinary(archivePath, tmpDir, asset); err != nil {
		return fmt.Errorf("extraction failed: %w", err)
	}

	// Get current executable path
	currentExe, err := GetExecutablePath()
	if err != nil {
		return fmt.Errorf("failed to get executable path: %w", err)
	}

	newBinary := filepath.Join(tmpDir, asset.binaryName)

	// Replace binary. Windows defers replacement until this process exits because
	// a running .exe cannot be overwritten.
	fmt.Println("Installing...")
	deferred, err := replaceBinary(newBinary, currentExe, latestVersion)
	if err != nil {
		return fmt.Errorf("failed to replace binary: %w", err)
	}
	if deferred {
		fmt.Printf("mdp %s downloaded; installation will complete momentarily.\n", latestVersion)
		return nil
	}

	updateInstalledVersion(latestVersion)
	fmt.Printf("Successfully upgraded to mdp %s\n", latestVersion)
	return nil
}

func downloadFile(url, destPath string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("download returned status %d", resp.StatusCode)
	}

	out, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func updateInstalledVersion(version string) {
	state, err := LoadState()
	if err != nil {
		state = &State{}
	}
	state.CurrentVersion = version
	state.LatestVersion = version
	state.LastCheckTime = time.Now()
	_ = SaveState(state)
}
