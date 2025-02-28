package package_doc_test

import (
	"github.com/godoc-lint/godoc-lint/pkg/check/package_doc"
	"github.com/godoc-lint/godoc-lint/pkg/model"
)

var _ model.Checker = &package_doc.PackageDocChecker{}
