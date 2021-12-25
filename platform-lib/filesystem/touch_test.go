package filesystem_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/pkorobeinikov/platform/platform-lib/filesystem"
)

func TestTouch(t *testing.T) {
	tests := []struct {
		name      string
		givenPath string
		givenOpts []TouchOption
	}{
		{
			name:      "touch empty file",
			givenPath: "testdata/touch/.emptyfile",
		},
		{
			name:      "touch go_mod",
			givenPath: "testdata/touch/go_mod",
			givenOpts: []TouchOption{
				WithContentsOfString("go 1.17"),
			},
		},
	}

	for _, tt := range tests {
		// Provide more assertions
		err := Touch(tt.givenPath, tt.givenOpts...)
		assert.NoError(t, err)

		_ = os.Remove(tt.givenPath)
	}
}
