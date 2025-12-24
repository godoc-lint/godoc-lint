package with_import_multiple_aliases

import (
	json1 "encoding/json"
)

var _ = json1.Encoder{}

// godoc with doclink to [json1.Encoder] and [json2.Encoder].
const Alpha = 0
