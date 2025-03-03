package config

import (
	"path/filepath"
	"regexp"

	"github.com/godoc-lint/godoc-lint/pkg/model"
)

// config represents the godoc-lint analyzer configuration.
type config struct {
	cwd string

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
	p, err := filepath.Rel(c.cwd, path)
	if err != nil {
		p = path
	}

	for _, re := range c.excludeAsRegexp {
		if re.MatchString(p) {
			return false
		}
	}
	if c.includeAsRegexp == nil {
		return true
	}
	for _, re := range c.includeAsRegexp {
		if re.MatchString(p) {
			return true
		}
	}
	return false
}

// GetRuleOptions implements the corresponding interface method.
func (c *config) GetRuleOptions() *model.RuleOptions {
	return c.options
}
