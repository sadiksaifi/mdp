package updater

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"os"
	"path/filepath"
	"testing"
)

func TestReleaseAssetFor(t *testing.T) {
	tests := []struct {
		name        string
		goos        string
		goarch      string
		archiveName string
		binaryName  string
		format      string
		wantErr     bool
	}{
		{name: "macOS AMD64", goos: "darwin", goarch: "amd64", archiveName: "mdp-darwin-amd64.tar.gz", binaryName: "mdp", format: "tar.gz"},
		{name: "macOS ARM64", goos: "darwin", goarch: "arm64", archiveName: "mdp-darwin-arm64.tar.gz", binaryName: "mdp", format: "tar.gz"},
		{name: "Linux AMD64", goos: "linux", goarch: "amd64", archiveName: "mdp-linux-amd64.tar.gz", binaryName: "mdp", format: "tar.gz"},
		{name: "Linux ARM64", goos: "linux", goarch: "arm64", archiveName: "mdp-linux-arm64.tar.gz", binaryName: "mdp", format: "tar.gz"},
		{name: "Windows AMD64", goos: "windows", goarch: "amd64", archiveName: "mdp-windows-amd64.zip", binaryName: "mdp.exe", format: "zip"},
		{name: "Windows ARM64", goos: "windows", goarch: "arm64", archiveName: "mdp-windows-arm64.zip", binaryName: "mdp.exe", format: "zip"},
		{name: "unsupported OS", goos: "freebsd", goarch: "amd64", wantErr: true},
		{name: "unsupported architecture", goos: "windows", goarch: "386", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			asset, err := releaseAssetFor(tt.goos, tt.goarch)
			if tt.wantErr {
				if err == nil {
					t.Fatal("releaseAssetFor() error = nil, want error")
				}
				return
			}
			if err != nil {
				t.Fatalf("releaseAssetFor() error = %v", err)
			}
			if asset.archiveName != tt.archiveName || asset.binaryName != tt.binaryName || asset.format != tt.format {
				t.Errorf("releaseAssetFor() = %+v, want archive=%q binary=%q format=%q", asset, tt.archiveName, tt.binaryName, tt.format)
			}
		})
	}
}

func TestExtractReleaseBinary(t *testing.T) {
	const content = "mdp test binary"
	tests := []struct {
		name   string
		asset  releaseAsset
		create func(*testing.T, string, []archiveTestEntry)
	}{
		{
			name:   "tar.gz",
			asset:  releaseAsset{binaryName: "mdp", format: "tar.gz"},
			create: createTarGz,
		},
		{
			name:   "zip",
			asset:  releaseAsset{binaryName: "mdp.exe", format: "zip"},
			create: createZip,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir := t.TempDir()
			archivePath := filepath.Join(tmpDir, "release.archive")
			tt.create(t, archivePath, []archiveTestEntry{{name: tt.asset.binaryName, content: content}})

			destDir := filepath.Join(tmpDir, "extract")
			if err := os.Mkdir(destDir, 0755); err != nil {
				t.Fatal(err)
			}
			if err := extractReleaseBinary(archivePath, destDir, tt.asset); err != nil {
				t.Fatalf("extractReleaseBinary() error = %v", err)
			}

			got, err := os.ReadFile(filepath.Join(destDir, tt.asset.binaryName))
			if err != nil {
				t.Fatal(err)
			}
			if string(got) != content {
				t.Errorf("extracted content = %q, want %q", got, content)
			}
		})
	}
}

func TestExtractZipRejectsUnsafeOrInvalidArchives(t *testing.T) {
	tests := []struct {
		name    string
		entries []archiveTestEntry
	}{
		{name: "missing binary", entries: []archiveTestEntry{{name: "README.txt", content: "missing"}}},
		{name: "duplicate binary", entries: []archiveTestEntry{{name: "mdp.exe", content: "one"}, {name: "mdp.exe", content: "two"}}},
		{name: "path traversal", entries: []archiveTestEntry{{name: "../mdp.exe", content: "unsafe"}}},
		{name: "Windows absolute path", entries: []archiveTestEntry{{name: `C:\mdp.exe`, content: "unsafe"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir := t.TempDir()
			archivePath := filepath.Join(tmpDir, "release.zip")
			createZip(t, archivePath, tt.entries)

			err := extractZip(archivePath, tmpDir, "mdp.exe")
			if err == nil {
				t.Fatal("extractZip() error = nil, want error")
			}
		})
	}
}

type archiveTestEntry struct {
	name    string
	content string
}

func createZip(t *testing.T, archivePath string, entries []archiveTestEntry) {
	t.Helper()
	file, err := os.Create(archivePath)
	if err != nil {
		t.Fatal(err)
	}
	zw := zip.NewWriter(file)
	for _, entry := range entries {
		writer, err := zw.Create(entry.name)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := writer.Write([]byte(entry.content)); err != nil {
			t.Fatal(err)
		}
	}
	if err := zw.Close(); err != nil {
		t.Fatal(err)
	}
	if err := file.Close(); err != nil {
		t.Fatal(err)
	}
}

func createTarGz(t *testing.T, archivePath string, entries []archiveTestEntry) {
	t.Helper()
	file, err := os.Create(archivePath)
	if err != nil {
		t.Fatal(err)
	}
	gz := gzip.NewWriter(file)
	tw := tar.NewWriter(gz)
	for _, entry := range entries {
		header := &tar.Header{Name: entry.name, Mode: 0755, Size: int64(len(entry.content)), Typeflag: tar.TypeReg}
		if err := tw.WriteHeader(header); err != nil {
			t.Fatal(err)
		}
		if _, err := tw.Write([]byte(entry.content)); err != nil {
			t.Fatal(err)
		}
	}
	if err := tw.Close(); err != nil {
		t.Fatal(err)
	}
	if err := gz.Close(); err != nil {
		t.Fatal(err)
	}
	if err := file.Close(); err != nil {
		t.Fatal(err)
	}
}
