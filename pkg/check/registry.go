package check

import (
	"github.com/godoc-lint/godoc-lint/pkg/check/max_len"
	"github.com/godoc-lint/godoc-lint/pkg/check/pkg_doc"
	"github.com/godoc-lint/godoc-lint/pkg/model"
)

// Registry implements a registry of rules.
type Registry struct {
	checkers     map[model.Checker]struct{}
	coveredRules model.RuleSet
}

// NewRegistry returns a new rule registry instance.
func NewRegistry(checkers ...model.Checker) *Registry {
	registry := Registry{
		checkers: make(map[model.Checker]struct{}, len(checkers)+10),
	}
	for _, c := range checkers {
		registry.Add(c)
	}
	return &registry
}

// NewPopulatedRegistry returns a registry with all supported rules registered.
func NewPopulatedRegistry() *Registry {
	return NewRegistry(
		max_len.NewMaxLenChecker(),
		pkg_doc.NewPkgDocChecker(),
	)
}

// Add implements the corresponding interface method.
func (r *Registry) Add(checker model.Checker) {
	if _, ok := r.checkers[checker]; ok {
		panic("checker already registered")
	}
	r.coveredRules = r.coveredRules.Merge(checker.GetCoveredRules())
	r.checkers[checker] = struct{}{}
}

// List implements the corresponding interface method.
func (r *Registry) List() []model.Checker {
	all := make([]model.Checker, 0, len(r.checkers))
	for c := range r.checkers {
		all = append(all, c)
	}
	return all
}

// GetCoveredRules implements the corresponding interface method.
func (r *Registry) GetCoveredRules() model.RuleSet {
	return r.coveredRules
}
