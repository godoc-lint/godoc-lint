package stdlib_doclink

import (
	"testing"

	"github.com/godoc-lint/godoc-lint/pkg/check/stdlib_doclink/internal"
	"github.com/stretchr/testify/assert"
)

func TestFindPotentialDoclinks(t *testing.T) {
	tests := []struct {
		name  string
		texts []string
		pi    packageImports
		want  []potentialDoclink
	}{
		{
			name:  "empty",
			texts: []string{""},
		},
		{
			name:  "whitespace",
			texts: []string{"  "},
		},
		{
			name:  "whitespace and LF",
			texts: []string{"  \t\n \n"},
		},
		{
			name:  "whitespace and CRLF",
			texts: []string{"  \t\r\n \r\n"},
		},
		{
			name: "ignore package-only",
			texts: []string{
				"fmt",
				"bytes",
				"encoding/json",
				"net/http",
				"works like fmt package",
				"works like fmt and bytes packages",
			},
		},
		{
			name: "ignore package-only even if aliased",
			pi: packageImports{
				importAsMap: map[string]string{
					"foo": "fmt",
				},
			},
			texts: []string{
				"foo",
				"fmt",
				"bytes",
				"encoding/json",
				"net/http",
				"works like fmt package",
				"works like fmt and bytes packages",
			},
		},
		{
			name: "do not match existing doclinks",
			texts: []string{
				"[encoding/json.Encoder]",
				"[*encoding/json.Encoder]",
				"[encoding/json.Encoder.Encode]",

				"works like [encoding/json.Encoder]",
				"works like [*encoding/json.Encoder]",
				"works like [encoding/json.Encoder.Encode]",

				"works like [encoding/json.Encoder].",
				"works like [*encoding/json.Encoder].",
				"works like [encoding/json.Encoder.Encode].",

				"works like [encoding/json.Encoder], etc.",
				"works like [*encoding/json.Encoder], etc.",
				"works like [encoding/json.Encoder.Encode], etc.",

				"works like [encoding/json.Encoder]\n",
				"works like [*encoding/json.Encoder]\n",
				"works like [encoding/json.Encoder.Encode]\n",

				"works like [encoding/json.Encoder]\nas expected",
				"works like [*encoding/json.Encoder]\nas expected",
				"works like [encoding/json.Encoder.Encode]\nas expected",

				"works like\n[encoding/json.Encoder] as expected",
				"works like\n[*encoding/json.Encoder] as expected",
				"works like\n[encoding/json.Encoder.Encode] as expected",

				"works like\n[encoding/json.Encoder]\nas expected",
				"works like\n[*encoding/json.Encoder]\nas expected",
				"works like\n[encoding/json.Encoder.Encode]\nas expected",
			},
		},
		{
			name: "match pkg.name, once, single package",
			texts: []string{
				"encoding/json.Encoder",
				"*encoding/json.Encoder",

				"works like encoding/json.Encoder",
				"works like *encoding/json.Encoder",

				"works like encoding/json.Encoder.",
				"works like *encoding/json.Encoder.",

				"works like encoding/json.Encoder, etc.",
				"works like *encoding/json.Encoder, etc.",

				"works like encoding/json.Encoder\n",
				"works like *encoding/json.Encoder\n",

				"works like encoding/json.Encoder\nas expected",
				"works like *encoding/json.Encoder\nas expected",

				"works like\nencoding/json.Encoder as expected",
				"works like\n*encoding/json.Encoder as expected",

				"works like\nencoding/json.Encoder\nas expected",
				"works like\n*encoding/json.Encoder\nas expected",
			},
			want: []potentialDoclink{
				{
					originalNoStar: "encoding/json.Encoder",
					doclink:        "[encoding/json.Encoder]",
					kind:           internal.SymbolKindType,
					count:          1,
				},
			},
		},
		{
			name: "match pkg.name, twice, single package",
			texts: []string{
				"encoding/json.Encoder encoding/json.Encoder",
				"encoding/json.Encoder *encoding/json.Encoder",
				"*encoding/json.Encoder *encoding/json.Encoder",
				"*encoding/json.Encoder *encoding/json.Encoder",

				"works like encoding/json.Encoder and encoding/json.Encoder",
				"works like encoding/json.Encoder and *encoding/json.Encoder",
				"works like *encoding/json.Encoder and encoding/json.Encoder",
				"works like *encoding/json.Encoder and *encoding/json.Encoder",

				"works like encoding/json.Encoder. encoding/json.Encoder.",
				"works like encoding/json.Encoder. *encoding/json.Encoder.",
				"works like *encoding/json.Encoder. encoding/json.Encoder.",
				"works like *encoding/json.Encoder. *encoding/json.Encoder.",

				"works like encoding/json.Encoder, encoding/json.Encoder, etc.",
				"works like encoding/json.Encoder, *encoding/json.Encoder, etc.",
				"works like *encoding/json.Encoder, encoding/json.Encoder, etc.",
				"works like *encoding/json.Encoder, *encoding/json.Encoder, etc.",

				"works like encoding/json.Encoder\nencoding/json.Encoder",
				"works like encoding/json.Encoder\n*encoding/json.Encoder",
				"works like *encoding/json.Encoder\nencoding/json.Encoder",
				"works like *encoding/json.Encoder\n*encoding/json.Encoder",

				"works like encoding/json.Encoder\nas expected\nencoding/json.Encoder",
				"works like encoding/json.Encoder\nas expected\n*encoding/json.Encoder",
				"works like *encoding/json.Encoder\nas expected\nencoding/json.Encoder",
				"works like *encoding/json.Encoder\nas expected\n*encoding/json.Encoder",

				"works like\nencoding/json.Encoder as expected encoding/json.Encoder",
				"works like\nencoding/json.Encoder as expected *encoding/json.Encoder",
				"works like\n*encoding/json.Encoder as expected encoding/json.Encoder",
				"works like\n*encoding/json.Encoder as expected *encoding/json.Encoder",

				"works like\nencoding/json.Encoder\nas expected\nencoding/json.Encoder\n",
				"works like\nencoding/json.Encoder\nas expected\n*encoding/json.Encoder\n",
				"works like\n*encoding/json.Encoder\nas expected\nencoding/json.Encoder\n",
				"works like\n*encoding/json.Encoder\nas expected\n*encoding/json.Encoder\n",
			},
			want: []potentialDoclink{
				{
					originalNoStar: "encoding/json.Encoder",
					doclink:        "[encoding/json.Encoder]",
					kind:           internal.SymbolKindType,
					count:          2,
				},
			},
		},
		{
			name: "match pkg.name, once, two different pkgs",
			texts: []string{
				"encoding/json.Encoder bytes.Buffer",
				"encoding/json.Encoder *bytes.Buffer",

				"bytes.Buffer encoding/json.Encoder",
				"bytes.Buffer *encoding/json.Encoder",

				"works like encoding/json.Encoder and bytes.Buffer",
				"works like encoding/json.Encoder and *bytes.Buffer",

				"works like encoding/json.Encoder. bytes.Buffer.",
				"works like encoding/json.Encoder. *bytes.Buffer.",

				"works like encoding/json.Encoder, bytes.Buffer, etc.",
				"works like encoding/json.Encoder, *bytes.Buffer, etc.",

				"works like encoding/json.Encoder\nbytes.Buffer",
				"works like encoding/json.Encoder\n*bytes.Buffer",

				"works like encoding/json.Encoder\nas expected\nbytes.Buffer",
				"works like encoding/json.Encoder\nas expected\n*bytes.Buffer",

				"works like\nencoding/json.Encoder as expected bytes.Buffer",
				"works like\nencoding/json.Encoder as expected *bytes.Buffer",

				"works like\nencoding/json.Encoder\nas expected\nbytes.Buffer\n",
				"works like\nencoding/json.Encoder\nas expected\n*bytes.Buffer\n",
			},
			want: []potentialDoclink{
				{
					originalNoStar: "bytes.Buffer",
					doclink:        "[bytes.Buffer]",
					kind:           internal.SymbolKindType,
					count:          1,
				},
				{
					originalNoStar: "encoding/json.Encoder",
					doclink:        "[encoding/json.Encoder]",
					kind:           internal.SymbolKindType,
					count:          1,
				},
			},
		},
		{
			name: "match pkg.recv.name, once, single package",
			texts: []string{
				"encoding/json.Encoder.Encode",
				"works like encoding/json.Encoder.Encode",
				"works like encoding/json.Encoder.Encode.",
				"works like encoding/json.Encoder.Encode, etc.",
				"works like encoding/json.Encoder.Encode\n",
				"works like encoding/json.Encoder.Encode\nas expected",
				"works like\nencoding/json.Encoder.Encode as expected",
				"works like\nencoding/json.Encoder.Encode\nas expected",
			},
			want: []potentialDoclink{
				{
					originalNoStar: "encoding/json.Encoder.Encode",
					doclink:        "[encoding/json.Encoder.Encode]",
					kind:           internal.SymbolKindMethod,
					count:          1,
				},
			},
		},
		{
			name: "match pkg.recv.name, twice, single package",
			texts: []string{
				"encoding/json.Encoder.Encode encoding/json.Encoder.Encode",
				"works like encoding/json.Encoder.Encode and encoding/json.Encoder.Encode",
				"works like encoding/json.Encoder.Encode. encoding/json.Encoder.Encode.",
				"works like encoding/json.Encoder.Encode, encoding/json.Encoder.Encode, etc.",
				"works like encoding/json.Encoder.Encode\nencoding/json.Encoder.Encode",
				"works like encoding/json.Encoder.Encode\nas expected\nencoding/json.Encoder.Encode",
				"works like\nencoding/json.Encoder.Encode as expected encoding/json.Encoder.Encode",
				"works like\nencoding/json.Encoder.Encode\nas expected\nencoding/json.Encoder.Encode\n",
			},
			want: []potentialDoclink{
				{
					originalNoStar: "encoding/json.Encoder.Encode",
					doclink:        "[encoding/json.Encoder.Encode]",
					kind:           internal.SymbolKindMethod,
					count:          2,
				},
			},
		},
		{
			name: "match pkg.recv.name, once, two different pkgs",
			texts: []string{
				"encoding/json.Encoder.Encode bytes.Buffer",
				"bytes.Buffer encoding/json.Encoder.Encode",
				"works like encoding/json.Encoder.Encode and bytes.Buffer",
				"works like encoding/json.Encoder.Encode. bytes.Buffer.",
				"works like encoding/json.Encoder.Encode, bytes.Buffer, etc.",
				"works like encoding/json.Encoder.Encode\nbytes.Buffer",
				"works like encoding/json.Encoder.Encode\nas expected\nbytes.Buffer",
				"works like\nencoding/json.Encoder.Encode as expected bytes.Buffer",
				"works like\nencoding/json.Encoder.Encode\nas expected\nbytes.Buffer\n",
			},
			want: []potentialDoclink{
				{
					originalNoStar: "bytes.Buffer",
					doclink:        "[bytes.Buffer]",
					kind:           internal.SymbolKindType,
					count:          1,
				},
				{
					originalNoStar: "encoding/json.Encoder.Encode",
					doclink:        "[encoding/json.Encoder.Encode]",
					kind:           internal.SymbolKindMethod,
					count:          1,
				},
			},
		},
		{
			name: "normal imports",
			pi: packageImports{
				importAsMap: map[string]string{
					"bytes": "bytes",
					"json":  "encoding/json",
				},
			},
			texts: []string{
				"json.Encoder bytes.Buffer",
				"bytes.Buffer json.Encoder",
				"works like json.Encoder and bytes.Buffer",
				"works like json.Encoder. bytes.Buffer.",
				"works like json.Encoder, bytes.Buffer, etc.",
				"works like json.Encoder\nbytes.Buffer",
				"works like json.Encoder\nas expected\nbytes.Buffer",
				"works like\njson.Encoder as expected bytes.Buffer",
				"works like\njson.Encoder\nas expected\nbytes.Buffer\n",
			},
			want: []potentialDoclink{
				{
					originalNoStar: "bytes.Buffer",
					doclink:        "[bytes.Buffer]",
					kind:           internal.SymbolKindType,
					count:          1,
				},
				{
					originalNoStar: "json.Encoder",
					doclink:        "[json.Encoder]",
					kind:           internal.SymbolKindType,
					count:          1,
				},
			},
		},
		{
			name: "imports with confusing alias",
			pi: packageImports{
				importAsMap: map[string]string{
					"bytes": "bytes",
					"fmt":   "encoding/json",
				},
			},
			texts: []string{
				"fmt.Println fmt.Encoder bytes.Buffer", // fmt.Println should not be picked since "fmt" is alias for "encoding/json"
				"bytes.Buffer fmt.Encoder",
				"works like fmt.Encoder and bytes.Buffer",
				"works like fmt.Encoder. bytes.Buffer.",
				"works like fmt.Encoder, bytes.Buffer, etc.",
				"works like fmt.Encoder\nbytes.Buffer",
				"works like fmt.Encoder\nas expected\nbytes.Buffer",
				"works like\nfmt.Encoder as expected bytes.Buffer",
				"works like\nfmt.Encoder\nas expected\nbytes.Buffer\n",
			},
			want: []potentialDoclink{
				{
					originalNoStar: "bytes.Buffer",
					doclink:        "[bytes.Buffer]",
					kind:           internal.SymbolKindType,
					count:          1,
				},
				{
					originalNoStar: "fmt.Encoder",
					doclink:        "[fmt.Encoder]",
					kind:           internal.SymbolKindType,
					count:          1,
				},
			},
		},
		{
			name: "imports with multiple aliases",
			pi: packageImports{
				importAsMap: map[string]string{
					"bytes": "bytes",
					"json1": "encoding/json",
					"json2": "encoding/json",
				},
			},
			texts: []string{
				"json1.Encoder bytes.Buffer",
				"bytes.Buffer json1.Encoder",
				"works like json1.Encoder and bytes.Buffer",
				"works like json1.Encoder. bytes.Buffer.",
				"works like json1.Encoder, bytes.Buffer, etc.",
				"works like json1.Encoder\nbytes.Buffer",
				"works like json1.Encoder\nas expected\nbytes.Buffer",
				"works like\njson1.Encoder as expected bytes.Buffer",
				"works like\njson1.Encoder\nas expected\nbytes.Buffer\n",
			},
			want: []potentialDoclink{
				{
					originalNoStar: "bytes.Buffer",
					doclink:        "[bytes.Buffer]",
					kind:           internal.SymbolKindType,
					count:          1,
				},
				{
					originalNoStar: "json1.Encoder",
					doclink:        "[json1.Encoder]",
					kind:           internal.SymbolKindType,
					count:          1,
				},
			},
		},
		{
			name: "imports with bad/colliding names",
			pi: packageImports{
				importAsMap: map[string]string{
					"bytes": "bytes",
					"blah":  "encoding/json",
				},
				badImportAs: map[string]struct{}{
					"blah": {}, // For example, "blah" has also appeared as an alias for "fmt" elsewhere
				},
			},
			texts: []string{
				"blah.Encoder bytes.Buffer",
				"bytes.Buffer blah.Encoder",
				"works like blah.Encoder and bytes.Buffer",
				"works like blah.Encoder. bytes.Buffer.",
				"works like blah.Encoder, bytes.Buffer, etc.",
				"works like blah.Encoder\nbytes.Buffer",
				"works like blah.Encoder\nas expected\nbytes.Buffer",
				"works like\nblah.Encoder as expected bytes.Buffer",
				"works like\nblah.Encoder\nas expected\nbytes.Buffer\n",
			},
			want: []potentialDoclink{
				{
					originalNoStar: "bytes.Buffer",
					doclink:        "[bytes.Buffer]",
					kind:           internal.SymbolKindType,
					count:          1,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, text := range tt.texts {
				got := findPotentialDoclinks(&tt.pi, text)
				assert.Equal(t, tt.want, got, "unexpected result for test %q and text %q", tt.name, text)
			}
		})
	}
}
