package config_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/godoc-lint/godoc-lint/pkg/check/pkg_doc"
	"github.com/godoc-lint/godoc-lint/pkg/config"
	"github.com/godoc-lint/godoc-lint/pkg/model"
)

func TestTransferOptions(t *testing.T) {
	uintPtr := func(x uint) *uint {
		return &x
	}

	def, err := config.FromYAML(config.DefaultConfigYAML)
	require.Nil(t, err)

	tests := []struct {
		name     string
		sources  []*config.PlainRuleOptions
		expected *model.RuleOptions
	}{{
		name:     "empty",
		sources:  []*config.PlainRuleOptions{{}},
		expected: &model.RuleOptions{},
	}, {
		name: "empty and then non-empty",
		sources: []*config.PlainRuleOptions{{}, {
			MaxLenLength: uintPtr(999),
		}},
		expected: &model.RuleOptions{
			MaxLenLength: 999,
		},
	}, {
		name: "non-empty and then empty",
		sources: []*config.PlainRuleOptions{{
			MaxLenLength: uintPtr(999),
		}, {}},
		expected: &model.RuleOptions{
			MaxLenLength: 999,
		},
	}, {
		name: "non-empty and then non-empty",
		sources: []*config.PlainRuleOptions{{
			MaxLenLength: uintPtr(888),
		}, {
			MaxLenLength: uintPtr(999),
		}},
		expected: &model.RuleOptions{
			MaxLenLength: 999,
		},
	}, {
		name:    "default",
		sources: []*config.PlainRuleOptions{def.Options},
		expected: &model.RuleOptions{
			MaxLenLength:                   77,
			MaxLenIncludeTests:             false,
			PkgDocIncludeTests:             false,
			SinglePkgDocIncludeTests:       false,
			RequirePkgDocIncludeTests:      false,
			SpecificFilePkgDocIncludeTests: false,
			SpecificFilePkgDocFilePattern:  pkg_doc.FilePatternDoc,
			RequireDocIncludeTests:         false,
			RequireDocIgnoreExported:       false,
			RequireDocIgnoreUnexported:     true,
			StartWithNameIncludeTests:      false,
			StartWithNameIncludeUnexported: false,
			NoUnusedLinkIncludeTests:       false,
		},
	},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)

			target := &model.RuleOptions{}
			for _, source := range tt.sources {
				config.TransferOptions(target, source)
			}

			require.Equal(tt.expected, target)
		})
	}
}

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
					SpecificFilePkgDocFilePattern: ptr("foo"),
				},
			},
			wantErr: []string{
				`invalid default set "foo"; must be one of ["all" "basic" "none"]`,
				`invalid rule name(s) to enable: ["foo" "bar" "baz"]`,
				`invalid rule name(s) to disable: ["foo" "bar" "baz"]`,
				`invalid inclusion pattern(s): ["(" ")"]`,
				`invalid exclusion pattern(s): ["(" ")"]`,
				fmt.Sprintf(`invalid file-pattern: "foo" (must be one of %q)`, strings.Join(pkg_doc.AllFilePatterns, ",")),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.pcfg.Validate()
			if len(tt.wantErr) == 0 {
				require.NoError(t, err)
			} else {
				for _, value := range tt.wantErr {
					require.ErrorContains(t, err, value)
				}
			}
		})
	}
}

func ptr[T any](v T) *T {
	return &v
}
