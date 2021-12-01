package service

import (
	"context"
	"fmt"

	"github.com/pkorobeinikov/platform/platform-lib/service/spec"
)

func Debug(ctx context.Context) error {
	fmt.Println("Debug service")

	spec, err := spec.Read()
	if err != nil {
		return err
	}

	fmt.Println("Service name:", spec.Name)

	return nil
}
