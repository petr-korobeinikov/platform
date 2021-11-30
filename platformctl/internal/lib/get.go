package lib

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

func Get(ctx context.Context, envs []string, lib, version string) error {
	var cmd *exec.Cmd

	fqln := fmt.Sprintf("%s@%s", lib, version)

	cmd = exec.CommandContext(ctx, `go`, `get`, `-x`, fqln)

	cmd.Env = append(os.Environ(), envs...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.CommandContext(ctx, `go`, `mod`, `tidy`, `-v`)

	cmd.Env = append(os.Environ(), envs...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
