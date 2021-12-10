package spec_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pkorobeinikov/platform/platform-lib/service/spec"
)

func TestSpec_EnvironmentFor(t *testing.T) {
	tests := []struct {
		name     string
		expected map[string]string
		given    *spec.Spec
	}{
		{
			name: "global and local are specified",
			expected: map[string]string{
				"WORKER_NAP_DURATION": "1s",
				"FOO":                 "LOCAL_FOO",
			},
			given: &spec.Spec{
				Name: "wordcounter",
				Environment: map[string]map[string]string{
					"_": {
						"WORKER_NAP_DURATION": "1s",
						"FOO":                 "GLOBAL_FOO",
					},
					"local": {
						"FOO": "LOCAL_FOO",
					},
				},
			},
		},
		{
			name: "global only is specified",
			expected: map[string]string{
				"WORKER_NAP_DURATION": "1s",
				"FOO":                 "GLOBAL_FOO",
			},
			given: &spec.Spec{
				Name: "wordcounter",
				Environment: map[string]map[string]string{
					"_": {
						"WORKER_NAP_DURATION": "1s",
						"FOO":                 "GLOBAL_FOO",
					},
				},
			},
		},
		{
			name: "local only is specified",
			expected: map[string]string{
				"FOO": "EXCLUSIVE_LOCAL_FOO",
			},
			given: &spec.Spec{
				Name: "wordcounter",
				Environment: map[string]map[string]string{
					"local": {
						"FOO": "EXCLUSIVE_LOCAL_FOO",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		actual := tt.given.EnvironmentFor("local")

		assert.Equal(t, tt.expected, actual, tt.name)
	}
}

func TestRead(t *testing.T) {
	_, _ = spec.Read()
}
