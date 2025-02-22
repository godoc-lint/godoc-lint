package model

import (
	"golang.org/x/tools/go/analysis"
)

// AnalysisContext provides contextual information about the running analysis.
type AnalysisContext struct {
	// Config provides analyzer configuration.
	Config Config

	// InspectorResult is the analysis result of the pre-run inspector.
	InspectorResult *InspectorResult
}

// Rule defines a linter rule.
type Rule interface {
	// GetName returns the name of the rule.
	GetName() string

	// Apply checks for the rule.
	Apply(actx *AnalysisContext, pass *analysis.Pass) error
}
