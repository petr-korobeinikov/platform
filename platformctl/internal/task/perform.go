package task

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/drone/envsubst"

	"github.com/pkorobeinikov/platform/platform-lib/service/env"
	"github.com/pkorobeinikov/platform/platform-lib/service/spec"

	"platformctl/internal/cfg"
)

func Perform(ctx context.Context, args []string) error {
	s, err := spec.Read()
	if err != nil {
		return err
	}

	if err := s.SetEnvironmentFor(cfg.ServiceEnv); err != nil {
		return err
	}

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	taskName := strings.Join(args, " ")
	task, err := s.TaskByName(taskName)
	if err != nil {
		return err
	}

	envs := s.ShellEnvironmentFor(cfg.ServiceEnv)

	containerArgs := []string{
		cfg.PlatformFlavorContainerRuntimeCtl,
		"run",
		"--network", "host",
		"--rm",
		"--env-file", env.File,
		"-v", fmt.Sprintf("%s:/service", pwd),
	}

	if workdir := trim(task.Argument.Workdir); workdir != "" {
		containerArgs = append(containerArgs, "--workdir", workdir)
	}

	containerArgs = append(containerArgs, task.Image)

	if command := trim(task.Argument.Command); command != "" {
		subst, err := envsubst.EvalEnv(command)
		if err != nil {
			return err
		}

		// Escape shell args?
		containerArgs = append(containerArgs, strings.Split(subst, " ")...)
	}

	cmd := exec.CommandContext(ctx, containerArgs[0], containerArgs[1:]...)

	cmd.Env = append(os.Environ(), envs...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func trim(s string) string {
	return strings.TrimSpace(
		strings.ReplaceAll(s, "\n", ""),
	)
}
