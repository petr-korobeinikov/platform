package service

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/pkorobeinikov/platform/platform-lib/service/env"
	"github.com/pkorobeinikov/platform/platform-lib/service/spec"
)

func Debug(ctx context.Context) error {
	fmt.Println("Debug service")

	spec, err := spec.Read()
	if err != nil {
		return err
	}

	fmt.Println("Service name:", spec.Name)

	// Here goes the generation of the `.env.gen` file.

	cmd := exec.CommandContext(ctx, `docker-compose`, `--env-file`, env.File, `up`, `-d`)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
