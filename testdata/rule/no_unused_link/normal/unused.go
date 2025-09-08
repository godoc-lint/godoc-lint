// some header

// godoc with unused link // want `godoc has unused link \("link"\)`
//
// [link]: https://foo.com
package no_unused_link

// godoc with unused link // want `godoc has unused link \("link"\)`
//
// [link]: https://foo.com
const UnusedLinkConst = 0

// godoc with unused link // want `godoc has unused link \("link"\)`
//
// [link]: https://foo.com
const (
	// godoc with unused link // want `godoc has unused link \("link"\)`
	//
	// [link]: https://foo.com
	MultiUnusedLinkConst = 0
)

// godoc with unused link // want `godoc has unused link \("link"\)`
//
// [link]: https://foo.com
type TUnusedLink int

// godoc with unused link // want `godoc has unused link \("link"\)`
//
// [link]: https://foo.com
type (
	TMultiUnusedLink int
)

// godoc with unused link // want `godoc has unused link \("link"\)`
//
// [link]: https://foo.com
func UnusedLinkFunc() {}
