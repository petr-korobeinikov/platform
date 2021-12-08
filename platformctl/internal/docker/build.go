package docker

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/pkorobeinikov/platform/platform-lib/service/spec"
)

func Build(ctx context.Context, s *spec.Spec) error {
	args := []string{
		"docker",
		"build",
		"--file",
		".platform/docker/Dockerfile",
		"--tag",
		fmt.Sprintf("%s:latest", s.Name),
		".",
	}

	cmd := exec.CommandContext(ctx, args[0], args[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
