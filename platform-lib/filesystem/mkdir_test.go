package filesystem_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/pkorobeinikov/platform/platform-lib/filesystem"
)

func TestMkDir(t *testing.T) {
	given := "testdata/mkdir/test_mkdir"
	err := MkDir(given)
	assert.NoError(t, err)

	_ = os.RemoveAll("testdata/mkdir")
}
