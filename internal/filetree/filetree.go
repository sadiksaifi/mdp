package filetree

import (
	"path/filepath"
	"sort"
	"strings"
)

// FileEntry represents a single markdown file.
type FileEntry struct {
	ID      string // Sanitized identifier for HTML id attribute
	Path    string // Original file path
	Name    string // Display name (filename without extension)
	RelPath string // Relative path for display in tree
	Content string // Converted HTML content
}

// TreeNode represents a node in the file tree (file or directory).
type TreeNode struct {
	Name     string
	IsDir    bool
	File     *FileEntry  // nil if IsDir is true
	Children []*TreeNode // empty if not a directory
}

// BuildTree creates a tree structure from a list of file entries.
// The baseDir is used to compute relative paths for display.
func BuildTree(files []FileEntry) *TreeNode {
	root := &TreeNode{
		Name:     "root",
		IsDir:    true,
		Children: make([]*TreeNode, 0),
	}

	for i := range files {
		file := &files[i]
		parts := strings.Split(file.RelPath, string(filepath.Separator))
		insertIntoTree(root, parts, file)
	}

	sortTree(root)
	return root
}

// insertIntoTree inserts a file into the tree at the correct location.
func insertIntoTree(node *TreeNode, pathParts []string, file *FileEntry) {
	if len(pathParts) == 0 {
		return
	}

	if len(pathParts) == 1 {
		// This is the file itself
		node.Children = append(node.Children, &TreeNode{
			Name:  file.Name,
			IsDir: false,
			File:  file,
		})
		return
	}

	// Find or create the directory
	dirName := pathParts[0]
	var dirNode *TreeNode
	for _, child := range node.Children {
		if child.IsDir && child.Name == dirName {
			dirNode = child
			break
		}
	}

	if dirNode == nil {
		dirNode = &TreeNode{
			Name:     dirName,
			IsDir:    true,
			Children: make([]*TreeNode, 0),
		}
		node.Children = append(node.Children, dirNode)
	}

	insertIntoTree(dirNode, pathParts[1:], file)
}

// sortTree sorts the tree nodes: directories first, then files, both alphabetically.
func sortTree(node *TreeNode) {
	if !node.IsDir {
		return
	}

	sort.Slice(node.Children, func(i, j int) bool {
		// Directories come before files
		if node.Children[i].IsDir != node.Children[j].IsDir {
			return node.Children[i].IsDir
		}
		// Alphabetical within same type
		return strings.ToLower(node.Children[i].Name) < strings.ToLower(node.Children[j].Name)
	})

	for _, child := range node.Children {
		sortTree(child)
	}
}
