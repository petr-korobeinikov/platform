package filesystem_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/pkorobeinikov/platform/platform-lib/filesystem"
)

func TestDirectoryExists(t *testing.T) {
	tests := []struct {
		name     string
		given    string
		expected bool
	}{
		{
			name:     "existent directory",
			given:    "testdata/is_directory_exists/existent_dir",
			expected: true,
		},
		{
			name:     "not existent directory",
			given:    "testdata/is_directory_exists/not_existent_dir",
			expected: false,
		},
	}

	for _, tt := range tests {
		actual := IsDirectoryExists(tt.given)
		assert.Equal(t, tt.expected, actual, tt.name)
	}
}
