package config

import (
	"regexp"

	"github.com/godoc-lint/godoc-lint/pkg/model"
)

// config represents the godoc-lint analyzer configuration.
type config struct {
	includeAsRegexp []*regexp.Regexp
	excludeAsRegexp []*regexp.Regexp
	enabledRules    *model.RuleSet
	disabledRules   *model.RuleSet
	options         *model.RuleOptions
}

// IsAnyRuleApplicable implements the corresponding interface method.
func (c *config) IsAnyRuleApplicable(rs model.RuleSet) bool {
	if c.disabledRules != nil && c.disabledRules.IsSupersetOf(rs) {
		return false
	}
	return c.enabledRules == nil || c.enabledRules.HasCommonsWith(rs)
}

// IsPathApplicable implements the corresponding interface method.
func (c *config) IsPathApplicable(path string) bool {
	for _, re := range c.excludeAsRegexp {
		if re.MatchString(path) {
			return false
		}
	}
	if c.includeAsRegexp == nil {
		return true
	}
	for _, re := range c.includeAsRegexp {
		if re.MatchString(path) {
			return true
		}
	}
	return false
}

// GetRuleOptions implements the corresponding interface method.
func (c *config) GetRuleOptions() *model.RuleOptions {
	return c.options
}
