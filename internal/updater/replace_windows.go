//go:build windows

package updater

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	upgradeHelperEnv    = "MDP_UPGRADE_HELPER"
	upgradeParentPIDEnv = "MDP_UPGRADE_PARENT_PID"
	upgradeTargetEnv    = "MDP_UPGRADE_TARGET"
	upgradeVersionEnv   = "MDP_UPGRADE_VERSION"
)

// RunUpgradeHelper completes a deferred Windows executable replacement.
// Windows locks a running .exe, so the downloaded mdp binary waits for the
// original process to exit before installing itself.
func RunUpgradeHelper() (bool, error) {
	if os.Getenv(upgradeHelperEnv) != "1" {
		cleanupStaleUpgradeFiles()
		return false, nil
	}

	parentPID, err := strconv.ParseUint(os.Getenv(upgradeParentPIDEnv), 10, 32)
	if err != nil {
		return true, fmt.Errorf("invalid upgrade parent PID: %w", err)
	}
	target := filepath.Clean(os.Getenv(upgradeTargetEnv))
	version := os.Getenv(upgradeVersionEnv)
	if target == "." || version == "" {
		return true, fmt.Errorf("incomplete Windows upgrade helper configuration")
	}

	helperPath, err := os.Executable()
	if err != nil {
		return true, fmt.Errorf("find upgrade helper executable: %w", err)
	}
	helperPath, err = filepath.Abs(helperPath)
	if err != nil {
		return true, fmt.Errorf("resolve upgrade helper path: %w", err)
	}
	target, err = filepath.Abs(target)
	if err != nil {
		return true, fmt.Errorf("resolve upgrade target path: %w", err)
	}
	if !strings.EqualFold(filepath.Clean(helperPath), target+".update.exe") {
		return true, fmt.Errorf("invalid Windows upgrade helper path")
	}

	if err := waitForProcessExit(uint32(parentPID), 30*time.Second); err != nil {
		return true, err
	}
	if err := installWindowsReplacement(helperPath, target); err != nil {
		return true, err
	}

	updateInstalledVersion(version)
	fmt.Printf("Successfully upgraded to mdp %s\n", version)
	return true, nil
}

func replaceBinary(src, dst, version string) (bool, error) {
	otherPID, err := findOtherRunningExecutable(dst)
	if err != nil {
		return false, fmt.Errorf("check for running mdp processes: %w", err)
	}
	if otherPID != 0 {
		return false, fmt.Errorf("another mdp process (PID %d) is using %s; stop it and retry", otherPID, dst)
	}

	stagePath := dst + ".update.exe"
	if err := os.Remove(stagePath); err != nil && !os.IsNotExist(err) {
		return false, fmt.Errorf("remove stale upgrade helper: %w", err)
	}
	if err := copyExecutable(src, stagePath, 0755); err != nil {
		return false, fmt.Errorf("stage Windows upgrade helper: %w", err)
	}

	cmd := exec.Command(stagePath)
	cmd.Env = append(os.Environ(),
		upgradeHelperEnv+"=1",
		upgradeParentPIDEnv+"="+strconv.Itoa(os.Getpid()),
		upgradeTargetEnv+"="+dst,
		upgradeVersionEnv+"="+version,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		_ = os.Remove(stagePath)
		return false, fmt.Errorf("start Windows upgrade helper: %w", err)
	}
	if err := cmd.Process.Release(); err != nil {
		return false, fmt.Errorf("release Windows upgrade helper: %w", err)
	}
	return true, nil
}

func findOtherRunningExecutable(target string) (uint32, error) {
	target, err := filepath.Abs(target)
	if err != nil {
		return 0, fmt.Errorf("resolve executable path: %w", err)
	}

	snapshot, err := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPPROCESS, 0)
	if err != nil {
		return 0, fmt.Errorf("list Windows processes: %w", err)
	}
	defer windows.CloseHandle(snapshot)

	var entry windows.ProcessEntry32
	entry.Size = uint32(unsafe.Sizeof(entry))
	if err := windows.Process32First(snapshot, &entry); err != nil {
		return 0, fmt.Errorf("read Windows process list: %w", err)
	}

	currentPID := uint32(os.Getpid())
	for {
		if entry.ProcessID != currentPID {
			processPath, ok := processExecutablePath(entry.ProcessID)
			if ok && equalWindowsPaths(processPath, target) {
				return entry.ProcessID, nil
			}
		}

		err = windows.Process32Next(snapshot, &entry)
		if errors.Is(err, windows.ERROR_NO_MORE_FILES) {
			return 0, nil
		}
		if err != nil {
			return 0, fmt.Errorf("read Windows process list: %w", err)
		}
	}
}

func processExecutablePath(pid uint32) (string, bool) {
	handle, err := windows.OpenProcess(windows.PROCESS_QUERY_LIMITED_INFORMATION, false, pid)
	if err != nil {
		return "", false
	}
	defer windows.CloseHandle(handle)

	buffer := make([]uint16, 32768)
	size := uint32(len(buffer))
	if err := windows.QueryFullProcessImageName(handle, 0, &buffer[0], &size); err != nil {
		return "", false
	}
	return windows.UTF16ToString(buffer[:size]), true
}

func equalWindowsPaths(left, right string) bool {
	left = strings.TrimPrefix(filepath.Clean(left), `\\?\`)
	right = strings.TrimPrefix(filepath.Clean(right), `\\?\`)
	return strings.EqualFold(left, right)
}

func waitForProcessExit(pid uint32, timeout time.Duration) error {
	handle, err := windows.OpenProcess(windows.SYNCHRONIZE, false, pid)
	if errors.Is(err, windows.ERROR_INVALID_PARAMETER) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("open upgrade parent process: %w", err)
	}
	defer windows.CloseHandle(handle)

	result, err := windows.WaitForSingleObject(handle, uint32(timeout/time.Millisecond))
	if err != nil {
		return fmt.Errorf("wait for upgrade parent process: %w", err)
	}
	if result == uint32(windows.WAIT_TIMEOUT) {
		return fmt.Errorf("timed out waiting for the original mdp process to exit")
	}
	if result != uint32(windows.WAIT_OBJECT_0) {
		return fmt.Errorf("unexpected wait result while upgrading: %d", result)
	}
	return nil
}

func installWindowsReplacement(helperPath, target string) error {
	backupPath := target + ".old"
	readyPath := target + ".new"
	for _, stalePath := range []string{backupPath, readyPath} {
		if err := os.Remove(stalePath); err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("remove stale upgrade file %s: %w", stalePath, err)
		}
	}

	info, err := os.Stat(target)
	if err != nil {
		return fmt.Errorf("inspect current executable: %w", err)
	}
	if err := copyExecutable(helperPath, readyPath, info.Mode()); err != nil {
		return fmt.Errorf("prepare replacement executable: %w", err)
	}

	var renameErr error
	deadline := time.Now().Add(10 * time.Second)
	for {
		renameErr = os.Rename(target, backupPath)
		if renameErr == nil || time.Now().After(deadline) {
			break
		}
		time.Sleep(200 * time.Millisecond)
	}
	if renameErr != nil {
		_ = os.Remove(readyPath)
		return fmt.Errorf("move current executable aside: %w", renameErr)
	}

	if err := os.Rename(readyPath, target); err != nil {
		_ = os.Remove(readyPath)
		restoreErr := os.Rename(backupPath, target)
		if restoreErr != nil {
			return fmt.Errorf("install replacement executable: %w (restore failed: %v; previous binary remains at %s)", err, restoreErr, backupPath)
		}
		return fmt.Errorf("install replacement executable: %w", err)
	}

	// Antivirus software can briefly retain the previous executable. A future
	// mdp invocation also removes this backup, so cleanup failure is non-fatal.
	_ = os.Remove(backupPath)
	return nil
}

func copyExecutable(src, dst string, mode os.FileMode) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.OpenFile(dst, os.O_CREATE|os.O_EXCL|os.O_WRONLY, mode.Perm())
	if err != nil {
		return err
	}
	_, copyErr := io.Copy(out, in)
	syncErr := out.Sync()
	closeErr := out.Close()
	if copyErr != nil {
		_ = os.Remove(dst)
		return copyErr
	}
	if syncErr != nil {
		_ = os.Remove(dst)
		return syncErr
	}
	if closeErr != nil {
		_ = os.Remove(dst)
		return closeErr
	}
	return nil
}

func cleanupStaleUpgradeFiles() {
	target, err := os.Executable()
	if err != nil {
		return
	}
	for _, suffix := range []string{".update.exe", ".old", ".new"} {
		_ = os.Remove(target + suffix)
	}
}
