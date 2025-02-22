package config

import (
	"regexp"

	"github.com/godoc-lint/godoc-lint/pkg/model"
)

// config represents the godoc-lint analyzer configuration.
type config struct {
	includeAsRegexp []*regexp.Regexp
	excludeAsRegexp []*regexp.Regexp
	enableAsMap     map[string]struct{}
	disableAsMap    map[string]struct{}
	options         *model.RuleOptions
}

// IsRuleApplicable implements the corresponding interface method.
func (c *config) IsRuleApplicable(ruleName string) bool {
	if _, ok := c.disableAsMap[ruleName]; ok {
		return false
	}
	if c.enableAsMap == nil {
		return true
	}
	_, ok := c.enableAsMap[ruleName]
	return ok
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
