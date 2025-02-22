package config_test

import (
	"testing"

	"github.com/godoc-lint/godoc-lint/pkg/config"
)

func TestDefaultConfigYAMLIsValid(t *testing.T) {
	_, err := config.FromYAML(config.DefaultConfigYAML)
	if err != nil {
		t.Fatalf("cannot parse default configuration")
	}
}
