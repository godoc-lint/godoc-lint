// some header

//godoclint:disable
//godoclint:disable pkg-doc require-pkg-doc
package directive_disable

//godoclint:disable

//godoclint:disable pkg-doc require-pkg-doc single-pkg-doc

//godoclint:disable start-with-name
var Foo string

// some godoc
//
//godoclint:disable start-with-name
var Bar string

// some godoc
//
//godoclint:disable
//godoclint:disable start-with-name
var Baz string

// parent godoc
//
//godoclint:disable
var (
	// first godoc
	//
	//godoclint:disable pkg-doc
	MultiFoo = 0
	// second godoc
	MultiBar = 0
)

//godoclint:disable max-len
