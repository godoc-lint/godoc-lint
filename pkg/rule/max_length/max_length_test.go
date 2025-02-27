package max_length_test

import (
	"github.com/godoc-lint/godoc-lint/pkg/model"
	"github.com/godoc-lint/godoc-lint/pkg/rule/max_length"
)

var _ model.Checker = &max_length.MaxLengthChecker{}
