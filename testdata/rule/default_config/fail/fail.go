// some header

// bad godoc // want `package godoc should start with "Package default_config "` `package has more than one godoc \("default_config"\)`
package default_config

// bad godoc // want `godoc should start with symbol name \("Foo"\)`
const Foo = 0

// Bar is a symbol. // want `deprecation note should be formatted as "Deprecated: "`
//
// DEPRECATED: invalid deprecation note
type Bar int
