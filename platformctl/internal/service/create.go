package service

import (
	"context"

	"github.com/pkorobeinikov/platform/platform-lib/service/validation"
)

func Create(ctx context.Context, serviceName string) error {
	if err := validation.EnsureServiceNameValid(serviceName); err != nil {
		return err
	}

	// Validate directory doesn't exists

	// Create directory

	// Create entrypoint cmd/service/main.go

	// Create do.mod

	return nil
}
