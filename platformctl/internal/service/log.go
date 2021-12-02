package service

import (
	"context"
	"os"
	"os/exec"

	"github.com/pkorobeinikov/platform/platform-lib/service/deployment"
)

func Log(ctx context.Context) error {
	// We need to support ability for tailing only service log.
	args := []string{
		`docker-compose`,
		`--file`,
		deployment.DockerComposeFile,
		`logs`,
		`--follow`,
		`--no-log-prefix`,
	}

	cmd := exec.CommandContext(ctx, args[0], args[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
