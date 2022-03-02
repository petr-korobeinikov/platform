package platform

import (
	"context"

	"platformctl/internal/cfg"
)

func IP(ctx context.Context) (string, error) {
	return dispatch(cfg.PlatformFlavorContainerRuntimeVM).IP(ctx)
}
