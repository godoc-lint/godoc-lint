package strict

// (BG: bad godoc)

// godoc // want `godoc should start with symbol name \("FooBG"\)`
const FooBG = 0

// godoc // want `godoc should start with symbol name \("OwlBG"\)`
var OwlBG = 0

// godoc // want `godoc should start with symbol name \("CatBG"\)`
type CatBG int

// godoc // want `godoc should start with symbol name \("YoloBG"\)`
func YoloBG() {}

// godoc // want `godoc should start with symbol name \("fooBG"\)`
const fooBG = 0

// godoc // want `godoc should start with symbol name \("owlBG"\)`
var owlBG = 0

// godoc // want `godoc should start with symbol name \("catBG"\)`
type catBG int

// godoc // want `godoc should start with symbol name \("yoloBG"\)`
func yoloBG() {}
