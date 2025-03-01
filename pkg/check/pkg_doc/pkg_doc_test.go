package pkg_doc_test

import (
	"testing"

	"github.com/godoc-lint/godoc-lint/pkg/check/pkg_doc"
	"github.com/stretchr/testify/assert"
)

func TestCheckPkgDocPrefix(t *testing.T) {
	assert := assert.New(t)

	const packageName = "foo"

	tests := []struct {
		name           string
		text           string
		startWith      string
		expectedPrefix string
		expectedOK     bool
	}{{
		name:           "empty",
		text:           "",
		startWith:      "Package",
		expectedPrefix: "",
		expectedOK:     true,
	}, {
		name:           "good: empty startWith, empty rest",
		text:           "foo",
		startWith:      "",
		expectedPrefix: "foo",
		expectedOK:     true,
	}, {
		name:           "good: empty startWith",
		text:           "foo is a package",
		startWith:      "",
		expectedPrefix: "foo",
		expectedOK:     true,
	}, {
		name:           "bad: empty startWith, bad rest",
		text:           "foo-is-a-package",
		startWith:      "",
		expectedPrefix: "foo",
	}, {
		name:           "bad: empty startWith, bad start",
		text:           " foo",
		startWith:      "",
		expectedPrefix: "foo",
	}, {
		name:           "good: empty rest",
		text:           "Package foo",
		startWith:      "Package",
		expectedPrefix: "Package foo",
		expectedOK:     true,
	}, {
		name:           "good",
		text:           "Package foo is a package",
		startWith:      "Package",
		expectedPrefix: "Package foo",
		expectedOK:     true,
	}, {
		name:           "bad rest",
		text:           "Package foo-is-a-package",
		startWith:      "Package",
		expectedPrefix: "Package foo",
	}, {
		name:           "bad start",
		text:           " Package foo",
		startWith:      "Package",
		expectedPrefix: "Package foo",
	},
	}

	for _, tt := range tests {
		prefix, ok := pkg_doc.CheckPkgDocPrefix(tt.text, tt.startWith, packageName)
		assert.Equal(tt.expectedOK, ok, "case: %q", tt.name)
		assert.Equal(tt.expectedPrefix, prefix, "case: %q", tt.name)
	}
}
