package platform

import (
	"os"
	"path"
)

const Directory = ".platform"

func CreateDirectory() (err error) {
	// Should be better.

	err = os.Mkdir(Directory, os.ModePerm)
	if err != nil {
		return
	}

	err = os.Mkdir(path.Join(Directory, "docker-compose"), os.ModePerm)
	if err != nil {
		return
	}

	err = os.Mkdir(path.Join(Directory, "docker"), os.ModePerm)
	if err != nil {
		return
	}

	err = os.Mkdir(path.Join(Directory, "env"), os.ModePerm)
	if err != nil {
		return
	}

	return
}
