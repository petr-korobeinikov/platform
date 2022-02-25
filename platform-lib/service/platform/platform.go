package platform

import (
	"os"
	"path"
)

const Directory = ".platform"

func CreateDirectory() error {
	dp := []string{
		Directory,
		path.Join(Directory, "docker-compose"),
		path.Join(Directory, "docker"),
		path.Join(Directory, "env"),
	}

	for _, d := range dp {
		if err := os.MkdirAll(d, os.ModePerm); err != nil {
			return err
		}
	}

	return nil
}
