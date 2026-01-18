// Package linkrewriter rewrites relative markdown links in HTML content
// to work within multi-file HTML output by converting them to fragment identifiers.
package linkrewriter

import (
	"net/url"
	"path"
	"regexp"
	"strings"

	"mdp/internal/filetree"
)

// LinkRewriter rewrites relative .md links to section fragment identifiers.
type LinkRewriter struct {
	pathToID map[string]string // normalized relative path -> section ID
}

// New creates a new LinkRewriter from a list of file entries.
func New(entries []filetree.FileEntry) *LinkRewriter {
	pathToID := make(map[string]string)
	for _, entry := range entries {
		// Normalize the path for lookups (use forward slashes, lowercase)
		normalized := normalizePath(entry.RelPath)
		pathToID[normalized] = entry.ID
	}
	return &LinkRewriter{pathToID: pathToID}
}

// RewriteLinks rewrites relative .md links in HTML content to fragment identifiers.
// sourceRelPath is the relative path of the source file (used to resolve relative links).
func (lr *LinkRewriter) RewriteLinks(html string, sourceRelPath string) string {
	// Match <a href="..."> patterns
	// This regex captures the href value including any surrounding quotes
	re := regexp.MustCompile(`(<a\s+[^>]*href=")([^"]+)("[^>]*>)`)

	sourceDir := path.Dir(sourceRelPath)
	if sourceDir == "." {
		sourceDir = ""
	}

	return re.ReplaceAllStringFunc(html, func(match string) string {
		parts := re.FindStringSubmatch(match)
		if len(parts) != 4 {
			return match
		}

		prefix := parts[1]  // <a href="
		href := parts[2]    // the link
		suffix := parts[3]  // ">

		rewritten := lr.rewriteHref(href, sourceDir)
		return prefix + rewritten + suffix
	})
}

// rewriteHref rewrites a single href value if it's a relative .md link.
func (lr *LinkRewriter) rewriteHref(href string, sourceDir string) string {
	// Skip external links
	if strings.HasPrefix(href, "http://") || strings.HasPrefix(href, "https://") {
		return href
	}

	// Skip fragment-only links
	if strings.HasPrefix(href, "#") {
		return href
	}

	// Skip mailto and other protocols
	if strings.Contains(href, ":") {
		return href
	}

	// URL decode the href for proper path resolution
	decodedHref, err := url.PathUnescape(href)
	if err != nil {
		decodedHref = href
	}

	// Strip any fragment identifier from the link
	linkPath, _, _ := strings.Cut(decodedHref, "#")

	// Only process .md links
	if !strings.HasSuffix(strings.ToLower(linkPath), ".md") {
		return href
	}

	// Resolve the relative path from the source file's directory
	var resolvedPath string
	if sourceDir == "" {
		resolvedPath = linkPath
	} else {
		resolvedPath = path.Join(sourceDir, linkPath)
	}

	// Clean the path (handles ../ and ./)
	resolvedPath = path.Clean(resolvedPath)

	// Normalize for lookup
	normalized := normalizePath(resolvedPath)

	// Look up the section ID
	if sectionID, ok := lr.pathToID[normalized]; ok {
		return "#" + sectionID
	}

	// Not found in our file set, leave unchanged
	return href
}

// normalizePath normalizes a path for consistent lookups.
func normalizePath(p string) string {
	// Use forward slashes
	normalized := strings.ReplaceAll(p, "\\", "/")
	// Remove leading ./
	normalized = strings.TrimPrefix(normalized, "./")
	// Lowercase for case-insensitive matching
	normalized = strings.ToLower(normalized)
	return normalized
}
