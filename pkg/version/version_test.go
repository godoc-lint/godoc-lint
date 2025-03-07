package version_test

import (
	"fmt"
	"testing"

	"github.com/godoc-lint/godoc-lint/pkg/version"
	"github.com/stretchr/testify/assert"
)

var _ fmt.Stringer = version.Version{}

func TestString(t *testing.T) {
	tests := []struct {
		name     string
		value    version.Version
		expected string
	}{{
		name:     "zero",
		expected: "0.0.0",
	}, {
		name:     "no suffix",
		value:    version.Version{1, 2, 3, ""},
		expected: "1.2.3",
	}, {
		name:     "all",
		value:    version.Version{1, 2, 3, "foo"},
		expected: "1.2.3-foo",
	},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, tt.value.String(), "case: %q", tt.name)
	}
}
