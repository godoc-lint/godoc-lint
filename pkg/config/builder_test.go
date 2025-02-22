package config_test

import (
	"github.com/godoc-lint/godoc-lint/pkg/config"
	"github.com/godoc-lint/godoc-lint/pkg/model"
)

var _ model.ConfigBuilder = &config.ConfigBuilder{}
