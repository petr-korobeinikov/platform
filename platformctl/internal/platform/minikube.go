package platform

import (
	"bytes"
	"context"
	"os"
	"os/exec"

	"github.com/spf13/viper"

	"github.com/pkorobeinikov/platform/platform-lib/str"
	"platformctl/internal/cfg"
)

func (b *minikubeBridge) Start(ctx context.Context) error {
	memory := viper.GetString("platform.minikube.memory")
	cpus := viper.GetString("platform.minikube.cpus")
	diskSize := viper.GetString("platform.minikube.disk-size")
	driver := viper.GetString("platform.minikube.driver")

	containerRuntime := viper.GetString("platform.flavor.container-runtime")

	kubernetesVersion := viper.GetString("platform.kubernetes.version")

	args := []string{
		"minikube",
		"--profile", cfg.PlatformMinikubeProfile,
		"start",
		"--memory", memory,
		"--cpus", cpus,
		"--disk-size", diskSize,
		"--driver", driver,
		"--container-runtime", containerRuntime,
		"--kubernetes-version", kubernetesVersion,
		"--delete-on-failure=true",
		"--addons=ingress",
	}

	cmd := exec.CommandContext(ctx, args[0], args[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func (b *minikubeBridge) Stop(ctx context.Context) error {
	profile := viper.GetString("platform.minikube.profile")

	args := []string{
		"minikube",
		"--profile", profile,
		"stop",
	}

	cmd := exec.CommandContext(ctx, args[0], args[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func (b *minikubeBridge) IP(ctx context.Context) (string, error) {
	var buf bytes.Buffer

	args := []string{
		"minikube",
		"--profile", cfg.PlatformMinikubeProfile,
		"ip",
	}

	cmd := exec.CommandContext(ctx, args[0], args[1:]...)

	cmd.Stdout = &buf
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return str.Trim(buf.String()), nil
}

func newMinikubeBridge() *minikubeBridge {
	return &minikubeBridge{}
}

type minikubeBridge struct {
}
