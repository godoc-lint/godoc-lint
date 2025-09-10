package model

// Rule represents a rule.
type Rule string

const (
	PkgDocRule        Rule = "pkg-doc"
	SinglePkgDocRule  Rule = "single-pkg-doc"
	RequirePkgDocRule Rule = "require-pkg-doc"
	StartWithNameRule Rule = "start-with-name"
	RequireDocRule    Rule = "require-doc"
	MaxLenRule        Rule = "max-len"
	NoUnusedLinkRule  Rule = "no-unused-link"
)

// AllRules is the set of all supported rules.
var AllRules = func() RuleSet {
	return RuleSet{}.Add(
		PkgDocRule,
		SinglePkgDocRule,
		RequirePkgDocRule,
		StartWithNameRule,
		RequireDocRule,
		MaxLenRule,
		NoUnusedLinkRule,
	)
}()
