package lib

import (
	"context"
	"os"
	"os/exec"
)

func Get(ctx context.Context, lib string, envs []string) error {
	cmd := exec.CommandContext(ctx, `go`, `get`, `-x`, lib)

	cmd.Env = append(os.Environ(), envs...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
