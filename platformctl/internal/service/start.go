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
	s, err := spec.Read()
	if err != nil {
		return err
	}

	environment := s.EnvironmentFor("local")
	env.Registry().RegisterMany(environment)

	generator := deployment.NewDockerComposeGenerator()
	deploymentSpec, err := generator.Generate(s)
	if err != nil {
		return err
	}

	err = deployment.WriteDockerComposeFile(deploymentSpec)
	if err != nil {
		return err
	}

	err = env.WriteEnvFile()
	if err != nil {
		return err
	}

	// if err := docker.EnsureSentinelNotRunning(ctx, s.Name); err != nil {
	// 	return err
	// }

	if err := docker.Build(ctx, s); err != nil {
		return err
	}

	args := []string{
		`docker`, `compose`,
		`--file`,
		deployment.DockerComposeFile,
		`--env-file`,
		env.File,
		`up`,
		`-d`,
		`--remove-orphans`,
	}

	cmd := exec.CommandContext(ctx, args[0], args[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
