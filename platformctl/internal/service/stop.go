package service

import (
	"context"
	"os"
	"os/exec"

	"github.com/pkorobeinikov/platform/platform-lib/service/deployment"
	"github.com/pkorobeinikov/platform/platform-lib/service/env"
)

func Stop(ctx context.Context) error {
	args := []string{
		`docker-compose`,
		`--file`,
		deployment.DockerComposeFile,
		`--env-file`,
		env.File,
		`down`,
		`--remove-orphans`,
	}

	cmd := exec.CommandContext(ctx, args[0], args[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}