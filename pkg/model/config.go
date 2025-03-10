package model

// ConfigBuilder defines a configuration builder.
type ConfigBuilder interface {
	// SetOverride sets the configuration override.
	SetOverride(override *ConfigOverride)

	// GetConfig builds and returns the configuration object for the given path.
	GetConfig(cwd string) (Config, error)
}

// ConfigOverride represents a configuration override.
//
// Non-nil values (including empty slices) indicate that the corresponding field
// is overridden.
type ConfigOverride struct {
	// ConfigFilePath is the path to config file.
	ConfigFilePath *string

	// Include is the overridden list of regexp patterns matching the files that
	// the linter should include.
	Include []string

	// Exclude is the overridden list of regexp patterns matching the files that
	// the linter should exclude.
	Exclude []string

	// Enable is the overridden list of rules to enable.
	Enable []string

	// Disable is the overridden list of rules to disable.
	Disable []string
}

// NewConfigOverride returns a new config override instance.
func NewConfigOverride() *ConfigOverride {
	return &ConfigOverride{}
}

// Config defines an analyzer configuration.
type Config interface {
	// IsAnyRuleEnabled determines if any of the given rule names is among
	// enabled rules, or not among disabled rules.
	IsAnyRuleApplicable(RuleSet) bool

	// IsPathApplicable determines if the given path matches the included path
	// patterns, or does not match the excluded path patterns.
	IsPathApplicable(path string) bool

	// Returns the rule-specific options.
	//
	// It never returns a nil pointer.
	GetRuleOptions() *RuleOptions
}

// RuleOptions represents individual linter rule configurations.
type RuleOptions struct {
	MaxLenLength               uint   `option:"max-len/length"`
	MaxLenIncludeTests         bool   `option:"max-len/include-tests"`
	PkgDocStartWith            string `option:"pkg-doc/start-with"`
	PkgDocIncludeTests         bool   `option:"pkg-doc/include-tests"`
	SinglePkgDocIncludeTests   bool   `option:"single-pkg-doc/include-tests"`
	RequirePkgDocIncludeTests  bool   `option:"require-pkg-doc/include-tests"`
	RequireDocIncludeTests     bool   `option:"require-doc/include-tests"`
	RequireDocIgnoreExported   bool   `option:"require-doc/ignore-exported"`
	RequireDocIgnoreUnexported bool   `option:"require-doc/ignore-unexported"`
	StartWithNamePattern       string `option:"start-with-name/pattern"`
	StartWithNameIncludeTests  bool   `option:"start-with-name/include-tests"`
	NoUnusedLinkIncludeTests   bool   `option:"no-unused-link/include-tests"`
}
