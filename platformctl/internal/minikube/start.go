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
		fmt.Sprintf("--profile=%s", cfg.MinikubeProfile()),
		"start",
		"--memory=4g",
		"--cpus=4",
		"--disk-size=50g",
		"--driver=docker",
		fmt.Sprintf("--kubernetes-version=%s", cfg.KuberneterVersion()),
		"--mount=true",
		fmt.Sprintf(`--mount-string=%[1]s:%[1]s`, homeDir),
		"--delete-on-failure=true",
	}

	cmd := exec.CommandContext(ctx, args[0], args[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	for _, addon := range []string{"metrics-server", "dashboard", "ingress", "default-storageclass", "storage-provisioner"} {
		if err := EnableAddon(ctx, addon); err != nil {
			return err
		}
	}

	// todo automatic /etc/hosts updater
	ip, err := IP(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("Update /etc/hosts with the following:\n\n%s	%s\n", cfg.MinikubeProfile(), ip)

	return nil
}
