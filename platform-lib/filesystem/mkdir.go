package filesystem

import "os"

func MkDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}
