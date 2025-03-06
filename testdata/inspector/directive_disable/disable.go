// some header

//godoclint:disable
//godoclint:disable foo bar
package directive_disable

//godoclint:disable

//godoclint:disable foo bar baz

//godoclint:disable yolo
var Foo string

// some godoc
//
//godoclint:disable yolo
var Bar string

// some godoc
//
//godoclint:disable
//godoclint:disable yolo
var Baz string

// parent godoc
//
//godoclint:disable
var (
	// first godoc
	//
	//godoclint:disable foo
	MultiFoo = 0
	// second godoc
	MultiBar = 0
)

//godoclint:disable foo-at-end
