package cfg

import "time"

func TimeoutDefault() time.Duration {
	return timeout
}

func TimeoutMediumOperation() time.Duration {
	return timeout * 6
}

func TimeoutHeavyOperation() time.Duration {
	return timeout * 15
}

func KuberneterVersion() string {
	return kubernetesVersion
}

func MinikubeProfile() string {
	return minikubeProfile
}

const (
	timeout = 10 * time.Second

	kubernetesVersion = "v1.22.2"

	minikubeProfile = "platform"
)
