package service

import (
	"context"
	"os"
	"os/exec"

	"github.com/pkorobeinikov/platform/platform-lib/service/deployment"
	"github.com/pkorobeinikov/platform/platform-lib/service/spec"
)

func Stop(ctx context.Context) error {
	s, err := spec.Read()
	if err != nil {
		return err
	}

	args := deployment.DockerComposeArgs(s.Name, `down`, `--remove-orphans`)

	cmd := exec.CommandContext(ctx, args[0], args[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
