//go:build !windows

package updater

import "os"

// RunUpgradeHelper reports whether this process is a deferred upgrade helper.
// Deferred replacement is only required on Windows.
func RunUpgradeHelper() (bool, error) {
	return false, nil
}

func replaceBinary(src, dst, _ string) (bool, error) {
	srcData, err := os.ReadFile(src)
	if err != nil {
		return false, err
	}

	dstInfo, err := os.Stat(dst)
	if err != nil {
		return false, err
	}

	if err := os.Rename(src, dst); err == nil {
		return false, nil
	}

	return false, os.WriteFile(dst, srcData, dstInfo.Mode())
}
