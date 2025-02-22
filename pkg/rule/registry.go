package rule

import (
	"errors"
	"fmt"

	"github.com/godoc-lint/godoc-lint/pkg/model"
)

// Registry implements a registry of rules.
type Registry struct {
	m map[string]model.Rule
}

// NewRegistry returns a new rule registry instance.
func NewRegistry(rule ...model.Rule) *Registry {
	registry := Registry{
		m: make(map[string]model.Rule),
	}
	for _, r := range rule {
		registry.Add(r)
	}
	return &registry
}

// NewPopulatedRegistry returns a registry with all supported rules registered.
func NewPopulatedRegistry() *Registry {
	return NewRegistry(
		NewMaxLengthRule(),
	)
}

// Add registers a new rule.
func (r *Registry) Add(rule model.Rule) {
	name := rule.GetName()
	if _, ok := r.m[name]; ok {
		panic(fmt.Sprintf("rule already registered: %s", name))
	}
	r.m[name] = rule
}

// Get returns the rule with the given name.
func (r *Registry) Get(ruleName string) (model.Rule, error) {
	if rule, ok := r.m[ruleName]; !ok {
		return nil, errors.New("rule not found")
	} else {
		return rule, nil
	}
}

// Rules returns a slice of the registered rules.
func (r *Registry) Rules() []model.Rule {
	all := make([]model.Rule, 0, len(r.m))
	for _, rule := range r.m {
		all = append(all, rule)
	}
	return all
}

// Names returns a slice of the registered rules' names.
func (r *Registry) Names() []string {
	all := make([]string, 0, len(r.m))
	for _, rule := range r.m {
		all = append(all, rule.GetName())
	}
	return all
}
