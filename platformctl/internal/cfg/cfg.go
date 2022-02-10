package cfg

import "time"

func TimeoutDefault() time.Duration {
	return time.Minute
}

func TimeoutMediumOperation() time.Duration {
	return 3 * time.Minute
}

func TimeoutHeavyOperation() time.Duration {
	return 10 * time.Minute
}

func KuberneterVersion() string {
	return kubernetesVersion
}

func MinikubeProfile() string {
	return minikubeProfile
}

const (
	kubernetesVersion = "v1.22.2"

	minikubeProfile = "platform"
)

var (
	ServiceEnv string
)
