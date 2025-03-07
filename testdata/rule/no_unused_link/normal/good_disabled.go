// some header

// godoc with unused link
//
// [link]: https://foo.com
//
//godoclint:disable no-unused-link
package no_unused_link

// godoc with unused link
//
// [link]: https://foo.com
//
//godoclint:disable
const UnusedLinkConstDisabled = 0

// godoc with unused link
//
// [link]: https://foo.com
//
//godoclint:disable no-unused-link
const (
	// godoc with unused link
	//
	// [link]: https://foo.com
	//godoclint:disable
	MultiUnusedLinkConstDisabled = 0
)

// godoc with unused link
//
// [link]: https://foo.com
//
//godoclint:disable no-unused-link
type TUnusedLinkDisabled int

// godoc with unused link
//
// [link]: https://foo.com
//
//godoclint:disable
type (
	TMultiUnusedLinkDisabled int
)

// godoc with unused link
//
// [link]: https://foo.com
//
//godoclint:disable no-unused-link
func UnusedLinkFuncDisabled() {}
