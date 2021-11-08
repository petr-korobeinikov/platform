package service

import (
	"context"
	"fmt"

	"platformctl/internal/deployment"
	"platformctl/internal/docker"
	"platformctl/internal/service"
)

func Start(ctx context.Context) error {
	spec, err := service.ReadSpec(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("Starting service `%s`.\n", spec.Name)

	if err := docker.Build(ctx); err != nil {
		return err
	}

	if err := deployment.FetchSpec(ctx); err != nil {
		return err
	}

	if err := deployment.ApplySpec(ctx); err != nil {
		return err
	}

	return nil
}
