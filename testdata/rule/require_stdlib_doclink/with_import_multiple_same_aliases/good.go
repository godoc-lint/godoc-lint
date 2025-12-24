package with_import_multiple_same_aliases

import (
	blah "encoding/json"
)

var _ = blah.Encoder{}

// godoc with doclink to [blah.Encoder].
const Alpha = 0
