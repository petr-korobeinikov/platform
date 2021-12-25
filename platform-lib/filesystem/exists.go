package filesystem

import "os"

func IsDirectoryExists(path string) bool {
	fi, err := os.Stat(path)

	return !os.IsNotExist(err) && fi.IsDir()
}
