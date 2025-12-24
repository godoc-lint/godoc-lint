package config_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/godoc-lint/godoc-lint/pkg/config"
	"github.com/godoc-lint/godoc-lint/pkg/model"
)

var _ model.ConfigBuilder = &config.ConfigBuilder{}

func TestTransferPrimitiveOptions(t *testing.T) {
	uintPtr := func(x uint) *uint {
		return &x
	}

	def, err := config.FromYAML(config.DefaultConfigYAML)
	require.NoError(t, err)

	tests := []struct {
		name     string
		sources  []*config.PlainRuleOptions
		expected *model.RuleOptions
	}{
		{
			name:     "empty",
			sources:  []*config.PlainRuleOptions{{}},
			expected: &model.RuleOptions{},
		},
		{
			name: "empty and then non-empty",
			sources: []*config.PlainRuleOptions{{}, {
				MaxLenLength: uintPtr(999),
			}},
			expected: &model.RuleOptions{
				MaxLenLength: 999,
			},
		},
		{
			name: "non-empty and then empty",
			sources: []*config.PlainRuleOptions{{
				MaxLenLength: uintPtr(999),
			}, {}},
			expected: &model.RuleOptions{
				MaxLenLength: 999,
			},
		},
		{
			name: "non-empty and then non-empty",
			sources: []*config.PlainRuleOptions{{
				MaxLenLength: uintPtr(888),
			}, {
				MaxLenLength: uintPtr(999),
			}},
			expected: &model.RuleOptions{
				MaxLenLength: 999,
			},
		},
		{
			name:    "default",
			sources: []*config.PlainRuleOptions{def.Options},
			expected: &model.RuleOptions{
				MaxLenLength:                     77,
				MaxLenIncludeTests:               false,
				PkgDocIncludeTests:               false,
				SinglePkgDocIncludeTests:         false,
				RequirePkgDocIncludeTests:        false,
				RequireDocIncludeTests:           false,
				RequireDocIgnoreExported:         false,
				RequireDocIgnoreUnexported:       true,
				StartWithNameIncludeTests:        false,
				StartWithNameIncludeUnexported:   false,
				RequireStdlibDoclinkIncludeTests: false,
				NoUnusedLinkIncludeTests:         false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)

			target := &model.RuleOptions{}
			for _, source := range tt.sources {
				config.TransferPrimitiveOptions(target, source)
			}

			require.Equal(tt.expected, target)
		})
	}
}

func TestConfigResolution(t *testing.T) {
	tests := []struct {
		name                     string
		fs                       map[string]string
		overrideConfigPath       string
		populatePlainConfig      bool
		expectedConfigCWDAt      map[string]string
		expectedConfigFilePathAt map[string]string
	}{{
		name: "no config file at base",
		fs:   map[string]string{},
		expectedConfigCWDAt: map[string]string{
			".":                   ".", // "." is a placeholder for base dir in this test
			"./foo/bar":           ".",
			"../path/out/of/base": ".",
		},
		expectedConfigFilePathAt: map[string]string{
			".":                   "", // empty means default config is loaded
			"./foo/bar":           "",
			"../path/out/of/base": "",
		},
	}, {
		name: "no config file at base, with config in subdirs",
		fs: map[string]string{
			"./foo/.godoc-lint.yaml":         "",
			"./foo/bar/baz/.godoc-lint.yaml": "",
		},
		expectedConfigCWDAt: map[string]string{
			".":                   ".", // "." is a placeholder for base dir in this test
			"./foo":               "./foo",
			"./foo/bar":           "./foo",
			"./foo/bar/baz":       "./foo/bar/baz",
			"./foo/bar/baz/foo":   "./foo/bar/baz",
			"./bar":               ".",
			"../path/out/of/base": ".",
		},
		expectedConfigFilePathAt: map[string]string{
			".":                   "", // empty means default config is loaded
			"./foo":               "./foo/.godoc-lint.yaml",
			"./foo/bar":           "./foo/.godoc-lint.yaml",
			"./foo/bar/baz":       "./foo/bar/baz/.godoc-lint.yaml",
			"./foo/bar/baz/foo":   "./foo/bar/baz/.godoc-lint.yaml",
			"./bar":               "",
			"../path/out/of/base": "",
		},
	}, {
		name:                "no config file as base, with plain config populated",
		fs:                  map[string]string{},
		populatePlainConfig: true,
		expectedConfigCWDAt: map[string]string{
			".":                   ".", // "." is a placeholder for base dir in this test
			"./foo/bar":           ".",
			"../path/out/of/base": ".",
		},
		expectedConfigFilePathAt: map[string]string{
			".":                   "", // empty means default config is loaded
			"./foo/bar":           "",
			"../path/out/of/base": "",
		},
	}, {
		name: "no config file at base, with config in subdirs, with plain config populated",
		fs: map[string]string{
			"./foo/.godoc-lint.yaml":         "",
			"./foo/bar/baz/.godoc-lint.yaml": "",
		},
		populatePlainConfig: true,
		expectedConfigCWDAt: map[string]string{
			".":                   ".", // "." is a placeholder for base dir in this test
			"./foo":               "./foo",
			"./foo/bar":           "./foo",
			"./foo/bar/baz":       "./foo/bar/baz",
			"./foo/bar/baz/foo":   "./foo/bar/baz",
			"./bar":               ".",
			"../path/out/of/base": ".",
		},
		expectedConfigFilePathAt: map[string]string{
			".":                   "", // empty means populated plain config
			"./foo":               "./foo/.godoc-lint.yaml",
			"./foo/bar":           "./foo/.godoc-lint.yaml",
			"./foo/bar/baz":       "./foo/bar/baz/.godoc-lint.yaml",
			"./foo/bar/baz/foo":   "./foo/bar/baz/.godoc-lint.yaml",
			"./bar":               "",
			"../path/out/of/base": "",
		},
	}, {
		name: "config file at base, with config in subdirs, with plain config populated",
		fs: map[string]string{
			"./.godoc-lint.yaml":             "",
			"./foo/.godoc-lint.yaml":         "",
			"./foo/bar/baz/.godoc-lint.yaml": "",
		},
		populatePlainConfig: true,
		expectedConfigCWDAt: map[string]string{
			".":                   ".", // "." is a placeholder for base dir in this test
			"./foo":               "./foo",
			"./foo/bar":           "./foo",
			"./foo/bar/baz":       "./foo/bar/baz",
			"./foo/bar/baz/foo":   "./foo/bar/baz",
			"./bar":               ".",
			"../path/out/of/base": ".",
		},
		expectedConfigFilePathAt: map[string]string{
			".":                   "./.godoc-lint.yaml",
			"./foo":               "./foo/.godoc-lint.yaml",
			"./foo/bar":           "./foo/.godoc-lint.yaml",
			"./foo/bar/baz":       "./foo/bar/baz/.godoc-lint.yaml",
			"./foo/bar/baz/foo":   "./foo/bar/baz/.godoc-lint.yaml",
			"./bar":               "./.godoc-lint.yaml",
			"../path/out/of/base": "./.godoc-lint.yaml",
		},
	}, {
		name: "pick up conventional config file at base dir: .godoc-lint.yaml",
		fs: map[string]string{
			"./.godoc-lint.yaml": "",
		},
		expectedConfigCWDAt: map[string]string{
			".":                   ".", // "." is a placeholder for base dir in this test
			"./foo/bar":           ".",
			"../path/out/of/base": ".",
		},
		expectedConfigFilePathAt: map[string]string{
			".":                   "./.godoc-lint.yaml",
			"./foo/bar":           "./.godoc-lint.yaml",
			"../path/out/of/base": "./.godoc-lint.yaml",
		},
	}, {
		name: "pick up conventional config file at base dir: .godoc-lint.yml",
		fs: map[string]string{
			"./.godoc-lint.yml": "",
		},
		expectedConfigCWDAt: map[string]string{
			".":                   ".", // "." is a placeholder for base dir in this test
			"./foo/bar":           ".",
			"../path/out/of/base": ".",
		},
		expectedConfigFilePathAt: map[string]string{
			".":                   "./.godoc-lint.yml",
			"./foo/bar":           "./.godoc-lint.yml",
			"../path/out/of/base": "./.godoc-lint.yml",
		},
	}, {
		name: "pick up conventional config file at base dir: .godoc-lint.json",
		fs: map[string]string{
			"./.godoc-lint.json": "",
		},
		expectedConfigCWDAt: map[string]string{
			".":                   ".", // "." is a placeholder for base dir in this test
			"./foo/bar":           ".",
			"../path/out/of/base": ".",
		},
		expectedConfigFilePathAt: map[string]string{
			".":                   "./.godoc-lint.json",
			"./foo/bar":           "./.godoc-lint.json",
			"../path/out/of/base": "./.godoc-lint.json",
		},
	}, {
		name: "pick up conventional config file at base dir: .godoclint.yaml",
		fs: map[string]string{
			"./.godoclint.yaml": "",
		},
		expectedConfigCWDAt: map[string]string{
			".":                   ".", // "." is a placeholder for base dir in this test
			"./foo/bar":           ".",
			"../path/out/of/base": ".",
		},
		expectedConfigFilePathAt: map[string]string{
			".":                   "./.godoclint.yaml",
			"./foo/bar":           "./.godoclint.yaml",
			"../path/out/of/base": "./.godoclint.yaml",
		},
	}, {
		name: "pick up conventional config file at base dir: .godoclint.yml",
		fs: map[string]string{
			"./.godoclint.yml": "",
		},
		expectedConfigCWDAt: map[string]string{
			".":                   ".", // "." is a placeholder for base dir in this test
			"./foo/bar":           ".",
			"../path/out/of/base": ".",
		},
		expectedConfigFilePathAt: map[string]string{
			".":                   "./.godoclint.yml",
			"./foo/bar":           "./.godoclint.yml",
			"../path/out/of/base": "./.godoclint.yml",
		},
	}, {
		name: "pick up conventional config file at base dir: .godoclint.json",
		fs: map[string]string{
			"./.godoclint.json": "",
		},
		expectedConfigCWDAt: map[string]string{
			".":                   ".", // "." is a placeholder for base dir in this test
			"./foo/bar":           ".",
			"../path/out/of/base": ".",
		},
		expectedConfigFilePathAt: map[string]string{
			".":                   "./.godoclint.json",
			"./foo/bar":           "./.godoclint.json",
			"../path/out/of/base": "./.godoclint.json",
		},
	}, {
		name: "override config file at base dir",
		fs: map[string]string{
			"./config": "",
		},
		overrideConfigPath: "./config",
		expectedConfigCWDAt: map[string]string{
			".":                   ".", // "." is a placeholder for base dir in this test
			"./foo/bar":           ".",
			"../path/out/of/base": ".",
		},
		expectedConfigFilePathAt: map[string]string{
			".":                   "./config",
			"./foo/bar":           "./config",
			"../path/out/of/base": "./config",
		},
	}, {
		name: "override config file at subdir",
		fs: map[string]string{
			"./subdir/config": "",
		},
		overrideConfigPath: "./subdir/config",
		expectedConfigCWDAt: map[string]string{
			".":                   ".", // "." is a placeholder for base dir in this test
			"./foo/bar":           ".",
			"../path/out/of/base": ".",
		},
		expectedConfigFilePathAt: map[string]string{
			".":                   "./subdir/config",
			"./foo/bar":           "./subdir/config",
			"../path/out/of/base": "./subdir/config",
		},
	}, {
		name: "override config file at base dir, with config in subdirs",
		fs: map[string]string{
			"./config":                       "",
			"./foo/.godoc-lint.yaml":         "",
			"./foo/bar/baz/.godoc-lint.yaml": "",
		},
		overrideConfigPath: "./config",
		expectedConfigCWDAt: map[string]string{
			".":                   ".", // "." is a placeholder for base dir in this test
			"./foo":               "./foo",
			"./foo/bar":           "./foo",
			"./foo/bar/baz":       "./foo/bar/baz",
			"./foo/bar/baz/foo":   "./foo/bar/baz",
			"./bar":               ".",
			"../path/out/of/base": ".",
		},
		expectedConfigFilePathAt: map[string]string{
			".":                   "./config",
			"./foo":               "./foo/.godoc-lint.yaml",
			"./foo/bar":           "./foo/.godoc-lint.yaml",
			"./foo/bar/baz":       "./foo/bar/baz/.godoc-lint.yaml",
			"./foo/bar/baz/foo":   "./foo/bar/baz/.godoc-lint.yaml",
			"./bar":               "./config",
			"../path/out/of/base": "./config",
		},
	}, {
		name: "override config file at subdir, with config in subdirs",
		fs: map[string]string{
			"./subdir/config":                "",
			"./foo/.godoc-lint.yaml":         "",
			"./foo/bar/baz/.godoc-lint.yaml": "",
		},
		overrideConfigPath: "./subdir/config",
		expectedConfigCWDAt: map[string]string{
			".":                   ".", // "." is a placeholder for base dir in this test
			"./foo":               "./foo",
			"./foo/bar":           "./foo",
			"./foo/bar/baz":       "./foo/bar/baz",
			"./foo/bar/baz/foo":   "./foo/bar/baz",
			"./bar":               ".",
			"../path/out/of/base": ".",
		},
		expectedConfigFilePathAt: map[string]string{
			".":                   "./subdir/config",
			"./foo":               "./foo/.godoc-lint.yaml",
			"./foo/bar":           "./foo/.godoc-lint.yaml",
			"./foo/bar/baz":       "./foo/bar/baz/.godoc-lint.yaml",
			"./foo/bar/baz/foo":   "./foo/bar/baz/.godoc-lint.yaml",
			"./bar":               "./subdir/config",
			"../path/out/of/base": "./subdir/config",
		},
	},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseDir := t.TempDir()

			err := setupFS(tt.fs, baseDir)
			require.NoError(t, err)

			cb := config.NewConfigBuilder(baseDir)
			if tt.populatePlainConfig {
				cb = cb.WithBaseDirPlainConfig(&config.PlainConfig{})
			}

			if tt.overrideConfigPath != "" {
				configPath := filepath.Join(baseDir, filepath.FromSlash(tt.overrideConfigPath))
				cb.SetOverride(&model.ConfigOverride{
					ConfigFilePath: &configPath,
				})
			}

			for relCWDWithSlash, expectedRelConfigCWD := range tt.expectedConfigCWDAt {
				cwd := filepath.FromSlash(relCWDWithSlash)
				cfg, err := cb.GetConfig(filepath.Join(baseDir, cwd))
				require.NoError(t, err)
				configCWD := cfg.GetCWD()
				expectedCWD := filepath.Join(baseDir, filepath.FromSlash(expectedRelConfigCWD))
				assert.Equal(t, expectedCWD, configCWD)
			}

			for relCWDWithSlash, expectedRelConfigFilePath := range tt.expectedConfigFilePathAt {
				cwd := filepath.FromSlash(relCWDWithSlash)
				cfg, err := cb.GetConfig(filepath.Join(baseDir, cwd))
				require.NoError(t, err)
				configFilePath := cfg.GetConfigFilePath()
				if expectedRelConfigFilePath == "" {
					assert.Empty(t, configFilePath)
				} else {
					wantConfigFilePath := filepath.Join(baseDir, filepath.FromSlash(expectedRelConfigFilePath))
					assert.Equal(t, wantConfigFilePath, configFilePath)
				}
			}
		})
	}
}

func setupFS(fs map[string]string, baseDir string) error {
	for pathWithSlash, content := range fs {
		fullPath := filepath.Join(baseDir, filepath.FromSlash(pathWithSlash))
		if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
			return fmt.Errorf("failed to create dir %q: %v", filepath.Dir(fullPath), err)
		}
		if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
			return fmt.Errorf("failed to write file %q: %v", fullPath, err)
		}
	}
	return nil
}
