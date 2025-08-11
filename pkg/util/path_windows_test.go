//go:build windows

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
		baseDir:  "C:\\foo",
		path:     "C:\\foo",
		expected: true,
	}, {
		name:     "same, indirect",
		baseDir:  "C:\\foo",
		path:     "C:\\..\\foo",
		expected: true,
	}, {
		name:     "under",
		baseDir:  "C:\\foo",
		path:     "C:\\foo\\bar",
		expected: true,
	}, {
		name:     "under, indirect",
		baseDir:  "C:\\foo",
		path:     "C:\\foo\\..\\foo\\bar",
		expected: true,
	}, {
		name:     "under, indirect 2",
		baseDir:  "C:\\foo",
		path:     "C:\\foo\\bar\\..\\bar",
		expected: true,
	}, {
		name:     "under, indirect 3",
		baseDir:  "C:\\foo",
		path:     "C:\\foo\\..\\bar\\baz\\..\\..\\foo\\bar",
		expected: true,
	}, {
		name:     "under 2",
		baseDir:  "C:\\foo",
		path:     "C:\\foo\\bar\\baz\\foo",
		expected: true,
	}, {
		name:     "parent",
		baseDir:  "C:\\foo\\bar",
		path:     "C:\\foo",
		expected: false,
	}, {
		name:     "parent, indirect",
		baseDir:  "C:\\foo\\bar",
		path:     "C:\\foo\\bar\\..\\..",
		expected: false,
	}, {
		name:     "not under",
		baseDir:  "C:\\foo",
		path:     "C:\\bar",
		expected: false,
	}, {
		name:     "not under, indirect",
		baseDir:  "C:\\foo",
		path:     "C:\\foo\\..\\bar",
		expected: false,
	}, {
		name:     "not under, different drives", // (See issue #16)
		baseDir:  "C:\\foo",
		path:     "D:\\bar",
		expected: false,
	},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, util.IsPathUnderBaseDir(tt.baseDir, tt.path), "case %q", tt.name)
	}
}
