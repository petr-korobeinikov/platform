package task

import (
	"context"
	"os"
	"os/exec"
	"strings"

	"github.com/pkorobeinikov/platform/platform-lib/service/spec"
)

func Perform(ctx context.Context, args []string) error {
	s, err := spec.Read()
	if err != nil {
		return err
	}

	taskName := strings.Join(args, " ")
	task, err := s.TaskByName(taskName)
	if err != nil {
		return err
	}

	// Pass service env vars into container
	// Pass volume into container
	containerArgs := []string{
		"docker",
		"run",
		"--pull", "always",
		"--rm",
		task.Image,
	}
	cmd := exec.CommandContext(ctx, containerArgs[0], containerArgs[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
