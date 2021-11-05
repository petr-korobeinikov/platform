package minikube

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"platformctl/internal/cfg"
)

func Start(ctx context.Context) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	args := []string{
		"minikube",
		"--profile=cloudmts",
		"start",
		"--memory=4g",
		"--cpus=4",
		"--disk-size=50g",
		`--addons="dashboard default-storageclass storage-provisioner metrics-server"`,
		"--driver=docker",
		fmt.Sprintf("--kubernetes-version=%s", cfg.KuberneterVersion()),
		"--mount=true",
		fmt.Sprintf(`--mount-string=%[1]s:%[1]s`, homeDir),
		"--delete-on-failure=true",
	}

	cmd := exec.CommandContext(ctx, args[0], args[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
