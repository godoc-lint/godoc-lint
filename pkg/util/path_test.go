//go:build !windows

package util_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/godoc-lint/godoc-lint/pkg/util"
)

func TestIsPathUnderBaseDir(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		name     string
		baseDir  string
		path     string
		expected bool
	}{{
		name:     "same",
		baseDir:  "/foo",
		path:     "/foo",
		expected: true,
	}, {
		name:     "same, indirect",
		baseDir:  "/foo",
		path:     "/../foo",
		expected: true,
	}, {
		name:     "under",
		baseDir:  "/foo",
		path:     "/foo/bar",
		expected: true,
	}, {
		name:     "under, indirect",
		baseDir:  "/foo",
		path:     "/foo/../foo/bar",
		expected: true,
	}, {
		name:     "under, indirect 2",
		baseDir:  "/foo",
		path:     "/foo/bar/../bar",
		expected: true,
	}, {
		name:     "under, indirect 3",
		baseDir:  "/foo",
		path:     "/foo/../bar/baz/../../foo/bar",
		expected: true,
	}, {
		name:     "under 2",
		baseDir:  "/foo",
		path:     "/foo/bar/baz/foo",
		expected: true,
	}, {
		name:     "parent",
		baseDir:  "/foo/bar",
		path:     "/foo",
		expected: false,
	}, {
		name:     "parent, indirect",
		baseDir:  "/foo/bar",
		path:     "/foo/bar/../..",
		expected: false,
	}, {
		name:     "not under",
		baseDir:  "/foo",
		path:     "/bar",
		expected: false,
	}, {
		name:     "not under, indirect",
		baseDir:  "/foo",
		path:     "/foo/../bar",
		expected: false,
	},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, util.IsPathUnderBaseDir(tt.baseDir, tt.path), "case %q", tt.name)
	}
}
