package lib

import (
	"context"
	"os"
	"os/exec"
)

func Upgrade(ctx context.Context, envs []string) error {
	cmd := exec.CommandContext(ctx, `go`, `mod`, `tidy`, `-v`)

	cmd.Env = append(os.Environ(), envs...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
