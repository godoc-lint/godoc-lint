package util

import (
	"path/filepath"
	"strings"
)

// IsPathUnderBaseDir determines whether the given path is a sub-directory of
// the given base, lexicographically.
func IsPathUnderBaseDir(baseDir, path string) (bool, error) {
	rel, err := filepath.Rel(baseDir, path)
	if err != nil {
		return false, err
	}
	return rel == "." || !(rel == ".." || strings.HasPrefix(filepath.ToSlash(rel), "../")), nil
}
