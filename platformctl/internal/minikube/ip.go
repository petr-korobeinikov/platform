package minikube

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"platformctl/internal/cfg"
)

func IP(ctx context.Context) (string, error) {
	var b bytes.Buffer

	args := []string{
		"minikube",
		fmt.Sprintf("--profile=%s", cfg.MinikubeProfile()),
		"ip",
	}

	cmd := exec.CommandContext(ctx, args[0], args[1:]...)

	cmd.Stdout = &b
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return strings.TrimSpace(b.String()), nil
}
