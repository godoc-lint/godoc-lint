package check_test

import (
	"github.com/godoc-lint/godoc-lint/pkg/check"
	"github.com/godoc-lint/godoc-lint/pkg/model"
)

var _ model.Registry = &check.Registry{}
