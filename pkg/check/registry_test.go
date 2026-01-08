package check_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/godoc-lint/godoc-lint/pkg/check"
	"github.com/godoc-lint/godoc-lint/pkg/model"
)

func TestPopulatedRegistryHasAllRules(t *testing.T) {
	rules := check.NewPopulatedRegistry().GetCoveredRules()
	assert.True(t, model.AllRules.IsSupersetOf(rules), "rule not defined in model")
	assert.True(t, rules.IsSupersetOf(model.AllRules), "checker for rule is not registered")
}
