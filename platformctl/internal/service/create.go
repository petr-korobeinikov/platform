package service

import (
	"context"
	"errors"
	"fmt"
	"path"

	"github.com/pkorobeinikov/platform/platform-lib/filesystem"
	"github.com/pkorobeinikov/platform/platform-lib/service/validation"
)

func Create(ctx context.Context, serviceName string) error {
	if err := validation.EnsureServiceNameValid(serviceName); err != nil {
		return err
	}

	if filesystem.IsDirectoryExists(serviceName) {
		return errServiceDirectoryAlreadyExists
	}

	if err := filesystem.MkDir(serviceName); err != nil {
		return err
	}

	// region Service Template
	if err := filesystem.MkDir(path.Join(serviceName, "cmd")); err != nil {
		return err
	}

	err := filesystem.Touch(
		path.Join(serviceName, "go.mod"),
		filesystem.WithContentsOfString(
			fmt.Sprintf(gomodFmt, serviceName),
		),
	)
	if err != nil {
		return err
	}
	// endregion

	return nil
}

const (
	gomodFmt = `module %s

go 1.17
`
)

var (
	errServiceDirectoryAlreadyExists = errors.New("service directory already exists")
)
