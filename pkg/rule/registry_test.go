package rule_test

import (
	"github.com/godoc-lint/godoc-lint/pkg/model"
	"github.com/godoc-lint/godoc-lint/pkg/rule"
)

var _ model.Registry = &rule.Registry{}
