package updater

import (
	"archive/tar"
	"compress/gzip"
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
// - curl/unknown: downloads and replaces binary
// - source: prints upgrade instructions
func PerformUpgrade(currentVersion string, force bool) error {
	method := DetectInstallMethod()

	switch method {
	case InstallMethodBrew:
		return handleBrewUpgrade()
	case InstallMethodSource:
		return handleSourceUpgrade()
	case InstallMethodCurl, InstallMethodUnknown:
		return handleCurlUpgrade(currentVersion, force)
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

func handleCurlUpgrade(currentVersion string, force bool) error {
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

	// Detect OS and architecture
	osName := runtime.GOOS
	archName := runtime.GOARCH

	// Find matching asset
	assetName := fmt.Sprintf("mdp-%s-%s.tar.gz", osName, archName)
	downloadURL := release.GetAssetURL(assetName)

	if downloadURL == "" {
		return fmt.Errorf("no release found for %s/%s", osName, archName)
	}

	// Download to temp directory
	tmpDir, err := os.MkdirTemp("", "mdp-upgrade-*")
	if err != nil {
		return fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	archivePath := filepath.Join(tmpDir, assetName)
	fmt.Println("Downloading...")
	if err := downloadFile(downloadURL, archivePath); err != nil {
		return fmt.Errorf("download failed: %w", err)
	}

	// Extract
	fmt.Println("Extracting...")
	if err := extractTarGz(archivePath, tmpDir); err != nil {
		return fmt.Errorf("extraction failed: %w", err)
	}

	// Get current executable path
	currentExe, err := GetExecutablePath()
	if err != nil {
		return fmt.Errorf("failed to get executable path: %w", err)
	}

	newBinary := filepath.Join(tmpDir, "mdp")

	// Replace binary
	fmt.Println("Installing...")
	if err := replaceBinary(newBinary, currentExe); err != nil {
		return fmt.Errorf("failed to replace binary: %w", err)
	}

	// Update state
	state, _ := LoadState()
	state.CurrentVersion = latestVersion
	state.LatestVersion = latestVersion
	state.LastCheckTime = time.Now()
	_ = SaveState(state)

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

func extractTarGz(archivePath, destDir string) error {
	file, err := os.Open(archivePath)
	if err != nil {
		return err
	}
	defer file.Close()

	gzr, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// Only extract the mdp binary
		if header.Name != "mdp" {
			continue
		}

		target := filepath.Join(destDir, header.Name)
		outFile, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
		if err != nil {
			return err
		}

		if _, err := io.Copy(outFile, tr); err != nil {
			outFile.Close()
			return err
		}
		outFile.Close()
	}

	return nil
}

func replaceBinary(src, dst string) error {
	// Read source file
	srcData, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	// Get destination file info for permissions
	dstInfo, err := os.Stat(dst)
	if err != nil {
		return err
	}

	// On Unix, we can often rename over the running binary
	// First try atomic rename
	if err := os.Rename(src, dst); err == nil {
		return nil
	}

	// Fallback: write directly (this works on most systems)
	return os.WriteFile(dst, srcData, dstInfo.Mode())
}
