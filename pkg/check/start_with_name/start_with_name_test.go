package start_with_name_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/godoc-lint/godoc-lint/pkg/check/start_with_name"
)

func TestGetStartMatcher(t *testing.T) {
	tests := []struct {
		name            string
		inputPattern    string
		expectedPattern string
		expectedMatch   map[string]string
		expectedError   string
	}{{
		name:            "empty",
		inputPattern:    "",
		expectedPattern: `^(?P<symbol_name>.+?)\b`,
		expectedMatch: map[string]string{
			"":                         "",
			"  ":                       "",
			"foo":                      "foo",
			"foo is a symbol":          "foo",
			"foo is\na symbol":         "foo",
			"foo\nis\na\nsymbol":       "foo",
			"foo\r\nis\r\na\r\nsymbol": "foo",
		},
	}, {
		name:            "valid",
		inputPattern:    "^(A|An) %",
		expectedPattern: `^(A|An) (?P<symbol_name>.+?)\b`,
		expectedMatch: map[string]string{
			"":                           "",
			"  ":                         "",
			"foo":                        "",
			"A foo":                      "foo",
			"A foo is a symbol":          "foo",
			"A foo\nis a symbol":         "foo",
			"A foo\r\nis\r\na\r\nsymbol": "foo",
			"\nA foo\nis\na\nsymbol":     "",
			"An owl":                     "owl",
			"An owl is a bird":           "owl",
			"An owl\nis a bird":          "owl",
			"An owl\r\nis\r\na\r\nbird":  "owl",
			"\nAn owl\nis\na\nbird":      "",
		},
	}, {
		name:            "valid, no start-anchor, no placeholder",
		inputPattern:    "(A|An) ",
		expectedPattern: `^(A|An) (?P<symbol_name>.+?)\b`,
	}, {
		name:            "valid, no placeholder",
		inputPattern:    "^(A|An) ",
		expectedPattern: `^(A|An) (?P<symbol_name>.+?)\b`,
	}, {
		name:            "valid, no start-anchor",
		inputPattern:    "(A|An) %",
		expectedPattern: `^(A|An) (?P<symbol_name>.+?)\b`,
	}, {
		name:          "invalid",
		inputPattern:  "(",
		expectedError: "invalid start pattern",
	},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			require := require.New(t)

			pattern, matcher, err := start_with_name.GetStartMatcher(tt.inputPattern)

			if tt.expectedError != "" {
				assert.ErrorContains(err, tt.expectedError)
				assert.Empty(pattern)
				assert.Nil(matcher)
				return
			}

			assert.Nil(err)
			assert.Equal(tt.expectedPattern, pattern)
			require.NotNil(matcher)

			for k, expected := range tt.expectedMatch {
				actual := matcher(k)
				assert.Equal(expected, actual)
			}
		})
	}
}
