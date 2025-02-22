package config

import (
	_ "embed"
)

// defaultConfigFiles is the list of default configuration file names.
var defaultConfigFiles = []string{
	".godoc-lint.yaml",
	".godoc-lint.yml",
	".godoc-lint.json",
	".godoclint.yaml",
	".godoclint.yml",
	".godoclint.json",
}

// defaultConfigYAML is the default configuration (as YAML).
//
//go:embed default.yaml
var defaultConfigYAML []byte
