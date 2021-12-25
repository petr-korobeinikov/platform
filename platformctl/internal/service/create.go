package service

import (
	"context"
	"errors"
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

	// Create do.mod
	// Create platform.yaml
	// endregion

	return nil
}

var (
	errServiceDirectoryAlreadyExists = errors.New("service directory already exists")
)
