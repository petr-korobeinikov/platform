package service

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/pkorobeinikov/platform/platform-lib/service/deployment"
	"github.com/pkorobeinikov/platform/platform-lib/service/env"
	"github.com/pkorobeinikov/platform/platform-lib/service/platform"
	"github.com/pkorobeinikov/platform/platform-lib/service/spec"
	"platformctl/internal/docker"
)

func Start(ctx context.Context) error {
	s, err := spec.Read()
	if err != nil {
		return err
	}

	_ = platform.CreateDirectory()

	if err := docker.Build(ctx, s); err != nil {
		return err
	}

	generator := deployment.NewDockerComposeGenerator()
	deploymentSpec, err := generator.Generate(s)
	if err != nil {
		return err
	}

	// Rework mkdir
	_ = os.Mkdir(path.Join(platform.Directory, "docker-compose"), os.ModePerm)

	// Extract into func
	err = os.WriteFile(deployment.DockerComposeFile, deploymentSpec, 0644)
	if err != nil {
		return err
	}

	environment := s.EnvironmentFor("local")
	env.Registry().RegisterMany(environment)

	// Rework mkdir
	_ = os.Mkdir(path.Join(platform.Directory, "env"), os.ModePerm)

	// Extract into func
	envFile, err := os.Create(env.File)
	if err != nil {
		return err
	}
	w := bufio.NewWriter(envFile)
	for k, v := range env.Registry().All() {
		_, err = w.WriteString(fmt.Sprintf("%s=%s\n", k, v))
		if err != nil {
			return err
		}
	}
	w.Flush()

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
