package cfg

import "time"

func TimeoutDefault() time.Duration {
	return timeout
}

func TimeoutHeavyOperation() time.Duration {
	return timeout * 15
}

func KuberneterVersion() string {
	return kubernetesVersion
}

const (
	timeout = 10 * time.Second

	kubernetesVersion = "v1.22.2"
)
