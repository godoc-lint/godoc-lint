package with_import_shadow

// Note that the aliased imports are in the other file, and they should affect
// the way we detect doclinks in this file as well.
//
//   import (
// 	    fmt "encoding/json"
//   )

// godoc with inapplicable potential doclink to fmt.Println because "fmt" is an alias for "encoding/json".
const Alpha = 0

// godoc with inapplicable potential doclink to json.Encoder because "encoding/json" is not known as "json".
const Beta = 0
