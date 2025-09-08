// some header

// godoc with unused link // want `godoc has unused link \("link"\)`
//
// [link]: https://foo.com
package no_unused_link

// godoc with unused link // want `godoc has unused link \("link"\)`
//
// [link]: https://foo.com
const UnusedLinkConstTest = 0

// godoc with unused link // want `godoc has unused link \("link"\)`
//
// [link]: https://foo.com
const (
	// godoc with unused link // want `godoc has unused link \("link"\)`
	//
	// [link]: https://foo.com
	MultiUnusedLinkConstTest = 0
)

// godoc with unused link // want `godoc has unused link \("link"\)`
//
// [link]: https://foo.com
type TUnusedLinkTest int

// godoc with unused link // want `godoc has unused link \("link"\)`
//
// [link]: https://foo.com
type (
	TMultiUnusedLinkTest int
)

// godoc with unused link // want `godoc has unused link \("link"\)`
//
// [link]: https://foo.com
func UnusedLinkFuncTest() {}
