package service

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/pkorobeinikov/platform/platform-lib/service/spec"
)

func Log(ctx context.Context) error {
	spec, err := spec.Read()
	if err != nil {
		return err
	}

	fmt.Println("Service name:", spec.Name)

	// We need to able support for tailing only service log.
	cmd := exec.CommandContext(ctx, `docker-compose`, `logs`, `--follow`, `--no-log-prefix`)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
