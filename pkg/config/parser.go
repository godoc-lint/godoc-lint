package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/godoc-lint/godoc-lint/pkg/model"
	"gopkg.in/yaml.v3"
)

// PlainConfig represents the plain configuration type as users would provide
// via a config file (e.g., a YAML file).
type PlainConfig struct {
	Version string   `yaml:"version"`
	Exclude []string `yaml:"exclude"`
	Include []string `yaml:"include"`
	Enable  []string `yaml:"enable"`
	Disable []string `yaml:"disable"`

	Options struct {
		MaxLength *struct {
			Length *uint `yaml:"length"`
		} `yaml:"max-length"`
	} `yaml:"options"`
}

func (c *PlainConfig) extractRuleOptions() *model.RuleOptions {
	ro := model.RuleOptions{}
	if c.Options.MaxLength != nil {
		ro.MaxLength = &model.MaxLengthRuleOptions{}
		if c.Options.MaxLength.Length != nil {
			ro.MaxLength.Length = uint(*c.Options.MaxLength.Length)
		}
	}
	return &ro
}

// FromYAML parses configuration from given YAML content.
func FromYAML(in []byte) (*PlainConfig, error) {
	raw := PlainConfig{}
	if err := yaml.Unmarshal(in, &raw); err != nil {
		return nil, fmt.Errorf("cannot parse config from YAML file: %w", err)
	}

	if raw.Version != "" && !strings.HasPrefix(raw.Version, "1.") {
		return nil, fmt.Errorf("unsupported config version: %s", raw.Version)
	}

	return &raw, nil
}

// FromYAML parses configuration from given file path.
func FromYAMLFile(path string) (*PlainConfig, error) {
	in, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot read file (%s): %w", path, err)
	}

	raw := PlainConfig{}
	if err := yaml.Unmarshal(in, &raw); err != nil {
		return nil, fmt.Errorf("cannot parse config from YAML file: %w", err)
	}

	if raw.Version != "" && !strings.HasPrefix(raw.Version, "1.") {
		return nil, fmt.Errorf("unsupported config version: %s", raw.Version)
	}

	return &raw, nil
}
