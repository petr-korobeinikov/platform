package validation_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/pkorobeinikov/platform/platform-lib/service/validation"
)

func TestEnsureServiceNameValid(t *testing.T) {
	tests := []struct {
		name     string
		expected error
		given    string
	}{
		{
			name:     "service name too long",
			expected: ErrServiceNameTooLong,
			given:    strings.Repeat("too-long-service-name", 10),
		},
		{
			name:     "service name contains underscore",
			expected: ErrServiceNameContainsUnderscore,
			given:    "service_name",
		},
		{
			name:     "service name contains dot",
			expected: ErrServiceNameContainsDot,
			given:    "service.name",
		},
	}

	for _, tt := range tests {
		actual := EnsureServiceNameValid(tt.given)
		assert.ErrorIs(t, actual, tt.expected)
	}
}
