package service

import (
	"context"
	"errors"

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

	// Create entrypoint cmd/service/main.go

	// Create do.mod

	return nil
}

var (
	errServiceDirectoryAlreadyExists = errors.New("service directory already exists")
)
