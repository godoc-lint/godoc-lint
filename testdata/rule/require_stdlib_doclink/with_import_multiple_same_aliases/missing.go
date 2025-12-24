package with_import_multiple_same_aliases

import (
	blah "encoding/json"
)

// (BG: bad godoc)

var _ = blah.Encoder{}

// godoc with potential doclink to blah.Encoder. // want `text "blah\.Encoder" should be replaced with "\[blah\.Encoder\]" to link to stdlib type`
const AlphaBG = 0
