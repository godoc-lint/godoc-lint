package config

import (
	"sync"

	"github.com/godoc-lint/godoc-lint/pkg/model"
)

// OnceConfigBuilder wraps a config builder and make it a one-time builder, so
// that further attempts to build will return the same result.
//
// This type is concurrent-safe.
type OnceConfigBuilder struct {
	builder model.ConfigBuilder

	mu          sync.Mutex
	isBuilt     bool
	buildResult model.Config
	buildErr    error
}

// NewOnceConfigBuilder crates a new instance of the corresponding struct.
func NewOnceConfigBuilder(builder model.ConfigBuilder) *OnceConfigBuilder {
	return &OnceConfigBuilder{
		builder: builder,
	}
}

// GetConfig implements the corresponding interface method.
func (ocb *OnceConfigBuilder) GetConfig() (model.Config, error) {
	ocb.mu.Lock()
	defer ocb.mu.Unlock()

	if !ocb.isBuilt {
		ocb.buildResult, ocb.buildErr = ocb.builder.GetConfig()
		ocb.isBuilt = true
	}
	return ocb.buildResult, ocb.buildErr
}

// SetOverride implements the corresponding interface method.
func (ocb *OnceConfigBuilder) SetOverride(override *model.ConfigOverride) {
	ocb.mu.Lock()
	defer ocb.mu.Unlock()

	if ocb.isBuilt {
		return
	}
	ocb.builder.SetOverride(override)
}
