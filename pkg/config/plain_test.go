package config_test

import (
	"testing"

	"github.com/stretchr/testify/require"

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
			MaxLen: uintPtr(999),
		}},
		expected: &model.RuleOptions{
			MaxLen: 999,
		},
	}, {
		name: "non-empty and then empty",
		sources: []*config.PlainRuleOptions{{
			MaxLen: uintPtr(999),
		}, {}},
		expected: &model.RuleOptions{
			MaxLen: 999,
		},
	}, {
		name: "non-empty and then non-empty",
		sources: []*config.PlainRuleOptions{{
			MaxLen: uintPtr(888),
		}, {
			MaxLen: uintPtr(999),
		}},
		expected: &model.RuleOptions{
			MaxLen: 999,
		},
	}, {
		name:    "default",
		sources: []*config.PlainRuleOptions{def.Options},
		expected: &model.RuleOptions{
			MaxLen:                     77,
			PkgDocStartWith:            "Package",
			RequirePkgDocIncludeTests:  false,
			RequireDocIncludeTests:     false,
			RequireDocIgnoreExported:   false,
			RequireDocIgnoreUnexported: true,
			StartWithNameIncludeTests:  false,
			StartWithNamePattern:       "((A|a|An|an|THE|The|the) )?%",
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
