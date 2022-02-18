package docker

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/pkorobeinikov/platform/platform-lib/service/platform"
	"github.com/pkorobeinikov/platform/platform-lib/service/spec"
	"platformctl/internal/cfg"
)

func Build(ctx context.Context, s *spec.Spec) error {
	err := os.WriteFile(path.Join(platform.Directory, "docker", "Dockerfile"), []byte(dockerfile), 0644)
	if err != nil {
		return err
	}

	args := []string{
		cfg.PlatformFlavorContainerRuntimeCtl,
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

// Move dockerfiles into lib?
//go:embed dockerfile/go/Dockerfile
var dockerfile string
