package model

// ConfigBuilder defines a configuration builder.
//
// A ConfigBuilder is a single-use object. That is, once either of the GetConfig
// or MustGetConfig methods is called, the object is sealed to further changes
// and will panic if one attempts to alter its state.
type ConfigBuilder interface {
	// SetOverride sets the configuration override.
	SetOverride(override *ConfigOverride)

	// GetConfig builds and returns the configuration object.
	//
	// Further calls to this method will return the same result.
	GetConfig() (Config, error)

	// MustGetConfig builds and returns the configuration object. It terminates
	// the process if there is an error.
	//
	// Further calls to this method will return the same result.
	MustGetConfig() Config
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
	// IsRuleEnabled determines if the given rule name is among enabled rules
	// or not among disabled rules.
	IsRuleApplicable(ruleName string) bool

	// IsPathApplicable determines if the given path matches the included path
	// patterns, or does not match the excluded path patterns.
	IsPathApplicable(path string) bool

	// Returns the rule-specific options
	GetRuleOptions() *RuleOptions
}

// RuleOptions represents individual linter rule configurations.
//
// Nil value indicate that the corresponding options are not assigned.
type RuleOptions struct {
	// MaxLength is the options for the `max-length` rule.
	MaxLength *MaxLengthRuleOptions
}

// MaxLengthRuleOptions represents options for the `max-length` rule.
type MaxLengthRuleOptions struct {
	// Length is the maximum line length.
	Length uint
}
