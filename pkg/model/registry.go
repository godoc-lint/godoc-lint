package model

// Registry defines a registry of rules.
type Registry interface {
	// Add registers a new rule.
	Add(rule Rule)

	// Get returns the rule with the given name.
	Get(ruleName string) (Rule, error)

	// Rules returns a slice of the registered rules.
	Rules() []Rule

	// Names returns a slice of the registered rules' names.
	Names() []string
}
