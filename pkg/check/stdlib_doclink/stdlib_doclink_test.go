package stdlib_doclink

import (
	"testing"

	"github.com/godoc-lint/godoc-lint/pkg/check/stdlib_doclink/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidPotentialDoclinkRE(t *testing.T) {
	require.NotNil(t, potentialDoclinkRE())
}

func TestFindPotentialDoclinks(t *testing.T) {
	tests := []struct {
		name  string
		texts []string
		want  []potentialDoclink
	}{
		{
			name:  "empty",
			texts: []string{""},
		}, {
			name:  "whitespace",
			texts: []string{"  "},
		}, {
			name:  "whitespace and LF",
			texts: []string{"  \t\n \n"},
		}, {
			name:  "whitespace and CRLF",
			texts: []string{"  \t\r\n \r\n"},
		}, {
			name: "ignore package-only",
			texts: []string{
				"fmt",
				"works like fmt package",
				"works like fmt and bytes packages",
			},
		},
		{
			name: "match pkg.name, once, single package",
			texts: []string{
				"fmt.Printf",
				"works like fmt.Printf",
				"works like fmt.Printf.",
				"works like fmt.Printf, etc.",
				"works like fmt.Printf\n",
				"works like fmt.Printf\nas expected",
				"works like\nfmt.Printf as expected",
				"works like\nfmt.Printf\nas expected",
			},
			want: []potentialDoclink{
				{original: "fmt.Printf", doclink: "[fmt.Printf]", kind: internal.SymbolKindFunc, count: 1},
			},
		},
		{
			name: "match pkg.name, twice, single package",
			texts: []string{
				"fmt.Printf fmt.Printf",
				"works like fmt.Printf and fmt.Printf",
				"works like fmt.Printf. fmt.Printf.",
				"works like fmt.Printf, fmt.Printf, etc.",
				"works like fmt.Printf\nfmt.Printf",
				"works like fmt.Printf\nas expected\nfmt.Printf",
				"works like\nfmt.Printf as expected fmt.Printf",
				"works like\nfmt.Printf\nas expected\nfmt.Printf\n",
			},
			want: []potentialDoclink{
				{original: "fmt.Printf", doclink: "[fmt.Printf]", kind: internal.SymbolKindFunc, count: 2},
			},
		},
		{
			name: "match pkg.name, once, two different pkgs",
			texts: []string{
				"fmt.Printf io.Reader",
				"io.Reader fmt.Printf",
				"works like fmt.Printf and io.Reader",
				"works like fmt.Printf. io.Reader.",
				"works like fmt.Printf, io.Reader, etc.",
				"works like fmt.Printf\nio.Reader",
				"works like fmt.Printf\nas expected\nio.Reader",
				"works like\nfmt.Printf as expected io.Reader",
				"works like\nfmt.Printf\nas expected\nio.Reader\n",
			},
			want: []potentialDoclink{
				{original: "fmt.Printf", doclink: "[fmt.Printf]", kind: internal.SymbolKindFunc, count: 1},
				{original: "io.Reader", doclink: "[io.Reader]", kind: internal.SymbolKindType, count: 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, text := range tt.texts {
				got := findPotentialDoclinks(text)
				assert.Equal(t, tt.want, got, "unexpected result for test %q and text %q", tt.name, text)
			}
		})
	}
}
