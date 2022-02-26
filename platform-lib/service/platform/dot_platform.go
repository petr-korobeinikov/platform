package platform

import (
	"path"

	"github.com/pkorobeinikov/platform/platform-lib/filesystem"
)

const Directory = ".platform"

func CreateDotPlatformDirectory() error {
	dp := []string{
		Directory,
		path.Join(Directory, "docker-compose"),
		path.Join(Directory, "docker"),
		path.Join(Directory, "env"),
	}

	for _, d := range dp {
		if err := filesystem.MkDir(d); err != nil {
			return err
		}
	}

	return nil
}
