package platform

import (
	"context"
	"os"
	"os/exec"

	"github.com/spf13/viper"
)

func (b *minikubeBridge) Start(ctx context.Context) error {
	profile := viper.GetString("platform.minikube.profile")
	memory := viper.GetString("platform.minikube.memory")
	cpus := viper.GetString("platform.minikube.cpus")
	diskSize := viper.GetString("platform.minikube.disk-size")
	driver := viper.GetString("platform.minikube.driver")

	kubernetesVersion := viper.GetString("platform.kubernetes.version")

	args := []string{
		"minikube",
		"--profile", profile,
		"start",
		"--memory", memory,
		"--cpus", cpus,
		"--disk-size", diskSize,
		"--driver", driver,
		"--kubernetes-version", kubernetesVersion,
		"--delete-on-failure=true",
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

func newMinikubeBridge() *minikubeBridge {
	return &minikubeBridge{}
}

type minikubeBridge struct {
}
