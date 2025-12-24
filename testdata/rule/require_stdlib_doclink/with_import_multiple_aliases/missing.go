package with_import_multiple_aliases

import (
	json2 "encoding/json"
)

// (BG: bad godoc)

var _ = json2.Encoder{}

// godoc with potential doclink to json1.Encoder. // want `text "json1\.Encoder" should be replaced with "\[json1\.Encoder\]" to link to stdlib type`
const AlphaBG = 0

// godoc with potential doclink to json2.Encoder. // want `text "json2\.Encoder" should be replaced with "\[json2\.Encoder\]" to link to stdlib type`
const BravoBG = 0
