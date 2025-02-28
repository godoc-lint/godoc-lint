package max_length_test

import (
	"github.com/godoc-lint/godoc-lint/pkg/check/max_length"
	"github.com/godoc-lint/godoc-lint/pkg/model"
)

var _ model.Checker = &max_length.MaxLengthChecker{}
