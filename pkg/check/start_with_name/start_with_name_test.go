package start_with_name_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/godoc-lint/godoc-lint/pkg/check/start_with_name"
)

func TestMatchSymbolName(t *testing.T) {
	tests := []struct {
		want   bool
		doc    string
		symbol string
	}{
		{false, "", "foo"},
		{false, "  ", "foo"},
		{false, "foobar", "foo"},
		{false, "foobar", "bar"},
		{false, "\nA foo is a symbol", "foo"},
		{false, "\na foo is a\nsymbol", "foo"},
		{false, "\nAn foo is a symbol", "foo"},
		{false, "\nAN foo is a symbol", "foo"},
		{false, "\nan foo is a symbol", "foo"},
		{false, "\nTHE foo is a symbol", "foo"},
		{false, "\nThe foo is a symbol", "foo"},
		{false, "\nthe foo is a symbol", "foo"},

		{true, "foo", "foo"},
		{true, "A", "A"},
		{true, "A sth", "A"},
		{true, "a", "a"},
		{true, "a sth", "a"},
		{true, "AN", "AN"},
		{true, "AN sth", "AN"},
		{true, "An", "An"},
		{true, "An sth", "An"},
		{true, "an", "an"},
		{true, "an sth", "an"},
		{true, "THE", "THE"},
		{true, "THE sth", "THE"},
		{true, "The", "The"},
		{true, "The sth", "The"},
		{true, "the", "the"},
		{true, "the sth", "the"},
		{true, "a an", "an"},
		{true, "a an", "a"},
		{true, "a the", "a"},
		{true, "a the", "the"},
		{true, "the a", "a"},
		{true, "the a", "the"},

		{true, "A foo", "foo"},
		{true, "A foo is a symbol", "foo"},
		{true, "A foo\n", "foo"},
		{true, "A foo\r\n", "foo"},
		{true, "A foo\nA bar", "foo"},
		{true, "A foo\r\nA bar", "foo"},
		{true, "A foo\nis a symbol", "foo"},
		{true, "A foo\r\nis a symbol", "foo"},
		{true, "a foo", "foo"},
		{true, "a foo is a symbol", "foo"},
		{true, "a foo\n", "foo"},
		{true, "a foo\r\n", "foo"},
		{true, "a foo\nA bar", "foo"},
		{true, "a foo\r\nA bar", "foo"},
		{true, "a foo\nis a symbol", "foo"},
		{true, "a foo\r\nis a symbol", "foo"},
		{true, "AN owl", "owl"},
		{true, "AN owl is a bird", "owl"},
		{true, "AN owl\n", "owl"},
		{true, "AN owl\r\n", "owl"},
		{true, "AN owl\nA bar", "owl"},
		{true, "AN owl\r\nA bar", "owl"},
		{true, "AN owl\nis a bird", "owl"},
		{true, "AN owl\r\nis a bird", "owl"},
		{true, "An owl", "owl"},
		{true, "An owl is a bird", "owl"},
		{true, "An owl\n", "owl"},
		{true, "An owl\r\n", "owl"},
		{true, "An owl\nA bar", "owl"},
		{true, "An owl\r\nA bar", "owl"},
		{true, "An owl\nis a bird", "owl"},
		{true, "An owl\r\nis a bird", "owl"},
		{true, "an owl", "owl"},
		{true, "an owl is a bird", "owl"},
		{true, "an owl\n", "owl"},
		{true, "an owl\r\n", "owl"},
		{true, "an owl\nA bar", "owl"},
		{true, "an owl\r\nA bar", "owl"},
		{true, "an owl\nis a bird", "owl"},
		{true, "an owl\r\nis a bird", "owl"},
		{true, "THE book", "book"},
		{true, "THE book is new", "book"},
		{true, "THE book\n", "book"},
		{true, "THE book\r\n", "book"},
		{true, "THE book\nA bar", "book"},
		{true, "THE book\r\nA bar", "book"},
		{true, "THE book\nis new", "book"},
		{true, "THE book\r\nis new", "book"},
		{true, "The book", "book"},
		{true, "The book is new", "book"},
		{true, "The book\n", "book"},
		{true, "The book\r\n", "book"},
		{true, "The book\nA bar", "book"},
		{true, "The book\r\nA bar", "book"},
		{true, "The book\nis new", "book"},
		{true, "The book\r\nis new", "book"},
		{true, "the book", "book"},
		{true, "the book is new", "book"},
		{true, "the book\n", "book"},
		{true, "the book\r\n", "book"},
		{true, "the book\nA bar", "book"},
		{true, "the book\r\nA bar", "book"},
		{true, "the book\nis new", "book"},
		{true, "the book\r\nis new", "book"},
	}

	for _, tt := range tests {
		got := start_with_name.MatchSymbolName(tt.doc, tt.symbol)
		assert.Equal(t, tt.want, got, "doc: %q", tt.doc)
	}
}
