package config_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/godoc-lint/godoc-lint/pkg/config"
)

func TestDefaultConfigYAMLIsValid(t *testing.T) {
	require := require.New(t)

	def, err := config.FromYAML(config.DefaultConfigYAML)
	require.NoError(err)

	require.NoError(def.Validate())

	// The default rule options must be non-nil for the default.
	require.NotNil(def.Options, "default rule options must be non-nil")

	visitedOptions := map[string]struct{}{}

	v := reflect.ValueOf(*def.Options)
	vt := reflect.TypeFor[config.PlainRuleOptions]()
	for i := range vt.NumField() {
		ft := vt.Field(i)

		tagYAML := ft.Tag.Get("yaml")
		require.NotEmpty(tagYAML, `"yaml" tag is required for field %q`, ft.Name)
		tagMapstructure := ft.Tag.Get("mapstructure")
		require.NotEmpty(tagMapstructure, `"mapstructure" tag is required for field %q`, ft.Name)

		require.Equal(tagMapstructure, tagYAML, `"mapstructure" and "yaml" tag values must be equal`)

		require.NotContains(visitedOptions, tagMapstructure, "duplicate option tag values: %q", tagMapstructure)
		visitedOptions[tagMapstructure] = struct{}{}

		// All fields must be assigned by not nil.
		f := v.Field(i)
		require.False(f.IsNil(), "value of %q must be non-nil", ft.Name)
	}
}

func TestDefaultConfigYAMLEqualsTheExample(t *testing.T) {
	require := require.New(t)

	def, err := config.FromYAML(config.DefaultConfigYAML)
	require.NoError(err)

	example, err := config.FromYAMLFile("../../.godoc-lint.default.yaml")
	require.NoError(err)

	require.Equal(def, example, "default config does not match the example file")
}
