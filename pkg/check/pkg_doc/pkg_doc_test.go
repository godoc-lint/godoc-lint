package pkg_doc_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/godoc-lint/godoc-lint/pkg/check/pkg_doc"
)

func TestCheckPkgDocPrefix(t *testing.T) {
	tests := []struct {
		want        bool
		doc         string
		packageName string
		expected    string
	}{
		{false, "", "", "Package "},
		{false, "foo", "foo", "Package foo"},
		{false, "Package", "foo", "Package foo"},
		{false, "Package foofoo", "foo", "Package foo"},
		{false, "Package foobar", "foo", "Package foo"},
		{false, "Package\nfoo", "foo", "Package foo"},
		{false, " Package foo", "foo", "Package foo"},
		{false, "\nPackage foo", "foo", "Package foo"},
		{false, "\tPackage foo", "foo", "Package foo"},
		{false, "Package\tfoo", "foo", "Package foo"},

		{true, "Package foo", "foo", "Package foo"},
		{true, "Package foo does nothing", "foo", "Package foo"},
		{true, "Package foo\ndoes nothing", "foo", "Package foo"},
		{true, "Package foo\r\ndoes nothing", "foo", "Package foo"},
	}

	for _, tt := range tests {
		prefix, got := pkg_doc.CheckPkgDocPrefix(tt.doc, tt.packageName)
		assert.Equal(t, tt.expected, prefix, "doc: %q", tt.doc)
		assert.Equal(t, tt.want, got, "doc: %q", tt.doc)
	}
}
