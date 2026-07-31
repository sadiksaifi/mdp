package updater

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const maxReleaseBinarySize = 128 << 20 // 128 MiB

type releaseAsset struct {
	archiveName string
	binaryName  string
	format      string
}

func releaseAssetFor(goos, goarch string) (releaseAsset, error) {
	if goarch != "amd64" && goarch != "arm64" {
		return releaseAsset{}, fmt.Errorf("unsupported architecture: %s", goarch)
	}

	switch goos {
	case "darwin", "linux":
		return releaseAsset{
			archiveName: fmt.Sprintf("mdp-%s-%s.tar.gz", goos, goarch),
			binaryName:  "mdp",
			format:      "tar.gz",
		}, nil
	case "windows":
		return releaseAsset{
			archiveName: fmt.Sprintf("mdp-windows-%s.zip", goarch),
			binaryName:  "mdp.exe",
			format:      "zip",
		}, nil
	default:
		return releaseAsset{}, fmt.Errorf("unsupported operating system: %s", goos)
	}
}

func extractReleaseBinary(archivePath, destDir string, asset releaseAsset) error {
	switch asset.format {
	case "tar.gz":
		return extractTarGz(archivePath, destDir, asset.binaryName)
	case "zip":
		return extractZip(archivePath, destDir, asset.binaryName)
	default:
		return fmt.Errorf("unsupported archive format: %s", asset.format)
	}
}

func extractTarGz(archivePath, destDir, expectedName string) error {
	file, err := os.Open(archivePath)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	gzr, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer func() { _ = gzr.Close() }()

	tr := tar.NewReader(gzr)
	found := false
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if err := validateArchivePath(header.Name); err != nil {
			return err
		}
		if header.Name != expectedName {
			continue
		}
		if found {
			return fmt.Errorf("archive contains duplicate %s", expectedName)
		}
		if header.Typeflag != tar.TypeReg && header.Typeflag != 0 {
			return fmt.Errorf("archive entry %s is not a regular file", expectedName)
		}
		if header.Size < 0 || header.Size > maxReleaseBinarySize {
			return fmt.Errorf("archive entry %s is too large", expectedName)
		}

		if err := writeReleaseBinary(filepath.Join(destDir, expectedName), tr, os.FileMode(header.Mode)); err != nil {
			return err
		}
		found = true
	}

	if !found {
		return fmt.Errorf("archive does not contain %s", expectedName)
	}
	return nil
}

func extractZip(archivePath, destDir, expectedName string) error {
	zr, err := zip.OpenReader(archivePath)
	if err != nil {
		return err
	}
	defer func() { _ = zr.Close() }()

	found := false
	for _, entry := range zr.File {
		if err := validateArchivePath(entry.Name); err != nil {
			return err
		}
		if entry.Name != expectedName {
			continue
		}
		if found {
			return fmt.Errorf("archive contains duplicate %s", expectedName)
		}
		if !entry.FileInfo().Mode().IsRegular() {
			return fmt.Errorf("archive entry %s is not a regular file", expectedName)
		}
		if entry.UncompressedSize64 > maxReleaseBinarySize {
			return fmt.Errorf("archive entry %s is too large", expectedName)
		}

		src, err := entry.Open()
		if err != nil {
			return err
		}
		writeErr := writeReleaseBinary(filepath.Join(destDir, expectedName), src, 0755)
		closeErr := src.Close()
		if writeErr != nil {
			return writeErr
		}
		if closeErr != nil {
			return closeErr
		}
		found = true
	}

	if !found {
		return fmt.Errorf("archive does not contain %s", expectedName)
	}
	return nil
}

func validateArchivePath(name string) error {
	cleaned := path.Clean(name)
	if name == "" || strings.HasPrefix(name, "/") || strings.Contains(name, "\\") ||
		strings.Contains(name, ":") || cleaned == ".." || strings.HasPrefix(cleaned, "../") {
		return fmt.Errorf("archive contains invalid path: %s", name)
	}
	return nil
}

func writeReleaseBinary(target string, src io.Reader, mode os.FileMode) error {
	if mode.Perm() == 0 {
		mode = 0755
	}
	out, err := os.OpenFile(target, os.O_CREATE|os.O_EXCL|os.O_WRONLY, mode.Perm())
	if err != nil {
		return err
	}

	written, copyErr := io.Copy(out, io.LimitReader(src, maxReleaseBinarySize+1))
	closeErr := out.Close()
	if copyErr != nil {
		_ = os.Remove(target)
		return copyErr
	}
	if closeErr != nil {
		_ = os.Remove(target)
		return closeErr
	}
	if written > maxReleaseBinarySize {
		_ = os.Remove(target)
		return fmt.Errorf("release binary is too large")
	}
	return nil
}
