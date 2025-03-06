package strict

// (BG: bad godoc)

// godoc // want `godoc should start with symbol name \(pattern "\(\(A\|a\|An\|an\|THE\|The\|the\) \)\?%"\)`
const FooBG = 0

// godoc // want `godoc should start with symbol name \(pattern "\(\(A\|a\|An\|an\|THE\|The\|the\) \)\?%"\)`
var OwlBG = 0

// godoc // want `godoc should start with symbol name \(pattern "\(\(A\|a\|An\|an\|THE\|The\|the\) \)\?%"\)`
type CatBG int

// godoc // want `godoc should start with symbol name \(pattern "\(\(A\|a\|An\|an\|THE\|The\|the\) \)\?%"\)`
func YoloBG() {}

// godoc // want `godoc should start with symbol name \(pattern "\(\(A\|a\|An\|an\|THE\|The\|the\) \)\?%"\)`
const fooBG = 0

// godoc // want `godoc should start with symbol name \(pattern "\(\(A\|a\|An\|an\|THE\|The\|the\) \)\?%"\)`
var owlBG = 0

// godoc // want `godoc should start with symbol name \(pattern "\(\(A\|a\|An\|an\|THE\|The\|the\) \)\?%"\)`
type catBG int

// godoc // want `godoc should start with symbol name \(pattern "\(\(A\|a\|An\|an\|THE\|The\|the\) \)\?%"\)`
func yoloBG() {}
