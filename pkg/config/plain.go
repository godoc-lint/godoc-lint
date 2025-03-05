package config

import (
	"reflect"

	"github.com/godoc-lint/godoc-lint/pkg/model"
)

// PlainConfig represents the plain configuration type as users would provide
// via a config file (e.g., a YAML file).
type PlainConfig struct {
	Version *string           `yaml:"version" mapstructure:"version"`
	Exclude []string          `yaml:"exclude" mapstructure:"exclude"`
	Include []string          `yaml:"include" mapstructure:"include"`
	Enable  []string          `yaml:"enable" mapstructure:"enable"`
	Disable []string          `yaml:"disable" mapstructure:"disable"`
	Options *PlainRuleOptions `yaml:"options" mapstructure:"options"`
}

type PlainRuleOptions struct {
	MaxLen                     *uint   `option:"max-len" yaml:"max-len" mapstructure:"max-len"`
	PkgDocStartWith            *string `option:"pkg-doc/start-with" yaml:"pkg-doc/start-with" mapstructure:"pkg-doc/start-with"`
	RequirePkgDocSkipTests     *bool   `option:"require-pkg-doc/skip-tests" yaml:"require-pkg-doc/skip-tests" mapstructure:"require-pkg-doc/skip-tests"`
	RequireDocSkipTests        *bool   `option:"require-doc/skip-tests" yaml:"require-doc/skip-tests" mapstructure:"require-doc/skip-tests"`
	RequireDocIgnoreExported   *bool   `option:"require-doc/ignore-exported" yaml:"require-doc/ignore-exported" mapstructure:"require-doc/ignore-exported"`
	RequireDocIgnoreUnexported *bool   `option:"require-doc/ignore-unexported" yaml:"require-doc/ignore-unexported" mapstructure:"require-doc/ignore-unexported"`
	StartWithNamePattern       *string `option:"start-with-name/pattern" yaml:"start-with-name/pattern" mapstructure:"start-with-name/pattern"`
}

func transferOptions(target *model.RuleOptions, source *PlainRuleOptions) {
	resV := reflect.ValueOf(target).Elem()
	resVT := resV.Type()

	resOptionMap := make(map[string]string, resVT.NumField())
	for i := 0; i < resVT.NumField(); i++ {
		ft := resVT.Field(i)
		key, ok := ft.Tag.Lookup("option")
		if !ok {
			continue
		}
		resOptionMap[key] = ft.Name
	}

	v := reflect.ValueOf(source).Elem()
	vt := v.Type()
	for i := 0; i < vt.NumField(); i++ {
		ft := vt.Field(i)
		key, ok := ft.Tag.Lookup("option")
		if !ok {
			continue
		}
		if ft.Type.Kind() != reflect.Pointer {
			continue
		}
		f := v.Field(i)
		if f.IsNil() {
			continue
		}
		resFieldName, ok := resOptionMap[key]
		if !ok {
			continue
		}
		resV.FieldByName(resFieldName).Set(f.Elem())
	}
}
