package cfg

import "time"

func TimeoutDefault() time.Duration {
	return time.Minute
}

func TimeoutMediumOperation() time.Duration {
	return 3 * time.Minute
}

func TimeoutHeavyOperation() time.Duration {
	return 30 * time.Minute
}

var (
	ServiceEnv string

	PlatformFlavorContainerRuntime    string
	PlatformFlavorContainerRuntimeCtl string
	PlatformFlavorContainerRuntimeVM  string

	PlatformMinikubeProfile string
)
