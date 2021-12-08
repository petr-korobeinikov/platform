package service

import (
	"context"
	"os"
	"os/exec"

	"github.com/pkorobeinikov/platform/platform-lib/service/deployment"
	"github.com/pkorobeinikov/platform/platform-lib/service/env"
	"github.com/pkorobeinikov/platform/platform-lib/service/spec"
	"platformctl/internal/docker"
)

func Start(ctx context.Context) error {
	spec, err := spec.Read()
	if err != nil {
		return err
	}

	// Here goes the generation of the `.env` file.

	if err := docker.Build(ctx, spec); err != nil {
		return err
	}

	generator := deployment.NewDockerComposeGenerator()
	deploymentSpec, err := generator.Generate(spec)
	if err != nil {
		return err
	}

	err = os.WriteFile(deployment.DockerComposeFile, deploymentSpec, 0644)
	if err != nil {
		return err
	}

	args := []string{
		`docker-compose`,
		`--file`,
		deployment.DockerComposeFile,
		`--env-file`,
		env.File,
		`up`,
		`-d`,
	}

	cmd := exec.CommandContext(ctx, args[0], args[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
