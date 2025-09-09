package shared

import (
	"go/doc/comment"
	"strings"
)

func HasDeprecatedParagraph(blocks []comment.Block) bool {
	for _, block := range blocks {
		par, ok := block.(*comment.Paragraph)
		if !ok || len(par.Text) == 0 {
			continue
		}
		text, ok := (par.Text[0]).(comment.Plain)
		if !ok {
			continue
		}
		if strings.HasPrefix(string(text), "Deprecated:") {
			return true
		}
	}
	return false
}
