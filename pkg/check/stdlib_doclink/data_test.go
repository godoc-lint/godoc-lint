package stdlib_doclink

import (
	"testing"

	"github.com/godoc-lint/godoc-lint/pkg/check/stdlib_doclink/internal"
	"github.com/stretchr/testify/require"
)

func TestParseStdlib(t *testing.T) {
	std, err := parseStdlib()
	require.NoError(t, err)

	var keys, symbols int
	for _, pkg := range std {
		keys++
		symbols += len(pkg.Symbols)
	}
	t.Logf("stdlib: %d packages, %d symbols", keys, symbols)

	require.Equal(t, "encoding/json", std["encoding/json"].Path)
	require.Equal(t, "json", std["encoding/json"].Name)
	require.Equal(t, internal.SymbolKindFunc, std["encoding/json"].Symbols["Marshal"])
	require.Equal(t, internal.SymbolKindMethod, std["encoding/json"].Symbols["Encoder.Encode"])
	require.Equal(t, internal.SymbolKindType, std["encoding/json"].Symbols["Encoder"])
	require.Equal(t, internal.SymbolKindConst, std["math"].Symbols["E"])
	require.Equal(t, internal.SymbolKindVar, std["io"].Symbols["EOF"])
}
