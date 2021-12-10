package platform

import "os"

const Directory = ".platform"

func CreateDirectory() error {
	// Should be better.
	return os.Mkdir(Directory, os.ModePerm)
}
