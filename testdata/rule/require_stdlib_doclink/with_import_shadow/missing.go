package with_import_shadow

import (
	fmt "encoding/json"
)

// (BG: bad godoc)

var _ = fmt.Encoder{}

// godoc with potential doclink to fmt.Encoder. // want `text "fmt\.Encoder" should be replaced with "\[fmt\.Encoder\]" to link to stdlib type`
const AlphaBG = 0

// godoc with potential doclink to encoding/json.Encoder. // want `text "encoding/json\.Encoder" should be replaced with "\[encoding/json\.Encoder\]" to link to stdlib type`
const BetaBG = 0
