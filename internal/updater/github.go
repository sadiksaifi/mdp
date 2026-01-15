package updater

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	// GitHubRepo is the repository path for mdp
	GitHubRepo = "sadiksaifi/mdp"
	// GitHubAPIURL is the base URL for GitHub API
	GitHubAPIURL = "https://api.github.com"
)

// Release represents a GitHub release
type Release struct {
	TagName string  `json:"tag_name"`
	Name    string  `json:"name"`
	Assets  []Asset `json:"assets"`
}

// Asset represents a release asset (downloadable file)
type Asset struct {
	Name               string `json:"name"`
	BrowserDownloadURL string `json:"browser_download_url"`
	Size               int64  `json:"size"`
}

// FetchLatestRelease queries the GitHub API to get the latest release.
// Returns nil and error if the request fails or times out.
func FetchLatestRelease() (*Release, error) {
	return FetchLatestReleaseWithContext(context.Background())
}

// FetchLatestReleaseWithContext queries the GitHub API with a context for cancellation.
func FetchLatestReleaseWithContext(ctx context.Context) (*Release, error) {
	url := fmt.Sprintf("%s/repos/%s/releases/latest", GitHubAPIURL, GitHubRepo)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("User-Agent", "mdp-updater")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch release: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 403 {
		return nil, fmt.Errorf("rate limited by GitHub API")
	}

	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("no releases found")
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("GitHub API returned status %d", resp.StatusCode)
	}

	var release Release
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &release, nil
}

// GetAssetURL returns the download URL for a specific asset name.
// Returns empty string if not found.
func (r *Release) GetAssetURL(assetName string) string {
	for _, asset := range r.Assets {
		if asset.Name == assetName {
			return asset.BrowserDownloadURL
		}
	}
	return ""
}
