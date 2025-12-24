package with_import_collision

import (
	blah "fmt"
)

var _ = blah.Println

// godoc with inapplicable potential doclink to blah.Println, due to import alias collision.
const Beta = 0
