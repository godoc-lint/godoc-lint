package config_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/godoc-lint/godoc-lint/pkg/config"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		pcfg    *config.PlainConfig
		wantErr []string
	}{
		{
			name: "zero",
			pcfg: &config.PlainConfig{},
		},
		{
			name: "default",
			pcfg: config.GetDefaultPlainConfig(),
		},
		{
			name: "invalid",
			pcfg: &config.PlainConfig{
				Default: ptr("foo"),
				Enable:  []string{"foo", "bar", "baz"},
				Disable: []string{"foo", "bar", "baz"},
				Include: []string{"(", ")"},
				Exclude: []string{"(", ")"},
				Options: &config.PlainRuleOptions{
					MaxLenIgnorePatterns: []string{"(", "^foo$", ")"},
				},
			},
			wantErr: []string{
				`invalid default set "foo"; must be one of ["all" "basic" "none"]`,
				`invalid rule name(s) to enable: ["foo" "bar" "baz"]`,
				`invalid rule name(s) to disable: ["foo" "bar" "baz"]`,
				`invalid inclusion pattern(s): ["(" ")"]`,
				`invalid exclusion pattern(s): ["(" ")"]`,
				`invalid max-len ignore pattern(s): ["(" ")"]`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.pcfg.Validate()
			if tt.wantErr == nil {
				require.NoError(t, err)
			} else {
				require.ErrorContains(t, err, strings.Join(tt.wantErr, "\n"))
			}
		})
	}
}

func ptr[T any](v T) *T {
	return &v
}
