package updater

import (
	"os"
	"path/filepath"
	"strings"
)

// InstallMethod represents how mdp was installed
type InstallMethod string

const (
	// InstallMethodBrew indicates installation via Homebrew
	InstallMethodBrew InstallMethod = "brew"
	// InstallMethodCurl indicates installation via curl script
	InstallMethodCurl InstallMethod = "curl"
	// InstallMethodSource indicates installation from source
	InstallMethodSource InstallMethod = "source"
	// InstallMethodUnknown indicates unknown installation method
	InstallMethodUnknown InstallMethod = "unknown"
)

// DetectInstallMethod determines how mdp was installed.
// Detection order:
// 1. Check if path contains "Cellar" or "homebrew" -> brew
// 2. Check for marker file ~/.local/share/mdp/.curl-installed -> curl
// 3. Check if binary in ~/.local/bin -> curl
// 4. Check if binary in GOPATH -> source
// 5. Default -> unknown
func DetectInstallMethod() InstallMethod {
	execPath, err := os.Executable()
	if err != nil {
		return InstallMethodUnknown
	}

	// Resolve symlinks to get actual binary location
	realPath, err := filepath.EvalSymlinks(execPath)
	if err != nil {
		realPath = execPath
	}

	return DetectInstallMethodFromPath(realPath)
}

// DetectInstallMethodFromPath determines install method from a given binary path.
// Useful for testing with mock paths.
func DetectInstallMethodFromPath(binaryPath string) InstallMethod {
	// Normalize path separators
	normalizedPath := filepath.ToSlash(binaryPath)
	lowerPath := strings.ToLower(normalizedPath)

	// 1. Check for Homebrew installation
	if strings.Contains(lowerPath, "cellar") ||
		strings.Contains(lowerPath, "homebrew") ||
		strings.Contains(lowerPath, "linuxbrew") {
		return InstallMethodBrew
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return InstallMethodUnknown
	}

	// 2. Check for curl installation marker file
	markerPath := filepath.Join(home, ".local", "share", "mdp", ".curl-installed")
	if _, err := os.Stat(markerPath); err == nil {
		return InstallMethodCurl
	}

	// 3. Check if binary is in ~/.local/bin (likely curl install)
	localBin := filepath.Join(home, ".local", "bin", "mdp")
	if binaryPath == localBin {
		return InstallMethodCurl
	}

	// 4. Check if binary is in GOPATH (source install via go install)
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = filepath.Join(home, "go")
	}
	if strings.HasPrefix(binaryPath, gopath) {
		return InstallMethodSource
	}

	// 5. Check if in /usr/local/bin (likely make install)
	if strings.HasPrefix(binaryPath, "/usr/local/bin") {
		return InstallMethodSource
	}

	return InstallMethodUnknown
}

// GetExecutablePath returns the path to the current executable.
// Resolves symlinks to get the actual binary location.
func GetExecutablePath() (string, error) {
	execPath, err := os.Executable()
	if err != nil {
		return "", err
	}

	realPath, err := filepath.EvalSymlinks(execPath)
	if err != nil {
		return execPath, nil
	}

	return realPath, nil
}

// GetMarkerFilePath returns the path to the curl installation marker file.
func GetMarkerFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".local", "share", "mdp", ".curl-installed"), nil
}

// CreateMarkerFile creates the curl installation marker file.
func CreateMarkerFile() error {
	markerPath, err := GetMarkerFilePath()
	if err != nil {
		return err
	}

	dir := filepath.Dir(markerPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	return os.WriteFile(markerPath, []byte("curl-installed\n"), 0644)
}
