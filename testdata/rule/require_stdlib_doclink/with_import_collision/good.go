package with_import_collision

// In these files, the same alias is used for two different import paths. In such
// cases, potential doclinks should be ignored. This is actually the doc tool
// behaviour.

import (
	blah "encoding/json"
)

var _ = blah.Encoder{}

// godoc with inapplicable potential doclink to blah.Encoder, due to import alias collision.
const Alpha = 0
