package service

import (
	"context"
	"errors"
	"os"
	"os/exec"

	"github.com/pkorobeinikov/platform/platform-lib/service/deployment"
	"github.com/pkorobeinikov/platform/platform-lib/service/spec"
	"platformctl/internal/cfg"
)

func Log(ctx context.Context) error {
	return ErrLogCommandUnimplemented

	s, err := spec.Read()
	if err != nil {
		return err
	}

	args := deployment.DockerComposeArgs(cfg.PlatformFlavorContainerRuntimeCtl, s.Name, `logs`, `--follow`, `--no-log-prefix`, `service`)

	cmd := exec.CommandContext(ctx, args[0], args[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

var ErrLogCommandUnimplemented = errors.New("log command not implemented")
