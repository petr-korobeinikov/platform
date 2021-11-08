package minikube

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"platformctl/internal/cfg"
)

func Stop(ctx context.Context) error {
	args := []string{
		"minikube",
		fmt.Sprintf("--profile=%s", cfg.MinikubeProfile()),
		"stop",
	}

	cmd := exec.CommandContext(ctx, args[0], args[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
