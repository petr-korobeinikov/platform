package platform

import "context"

type bridge interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

var (
	_ bridge = (*minikubeBridge)(nil)
	_ bridge = (*unstartableBridge)(nil)
	_ bridge = (*unsupportedBridge)(nil)
)

func dispatch(vm string) bridge {
	switch vm {
	default:
		return newUnsupportedBridge()
	case "docker-desktop", "rancher-desktop":
		return newUnstartableBridge()
	case "minikube":
		return newMinikubeBridge()
	}
}
