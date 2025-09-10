package check_test

import (
	"testing"

	"github.com/godoc-lint/godoc-lint/pkg/check"
	"github.com/godoc-lint/godoc-lint/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestPopulatedRegistryHasAllRules(t *testing.T) {
	rules := check.NewPopulatedRegistry().GetCoveredRules()
	assert.True(t, model.AllRules.IsSupersetOf(rules), "checker for rule is not registered")
	assert.True(t, rules.IsSupersetOf(model.AllRules), "rule not defined in model")
}
