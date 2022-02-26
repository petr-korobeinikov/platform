package platform

import (
	"context"

	"platformctl/internal/cfg"
)

func Stop(ctx context.Context) error {
	return dispatch(cfg.PlatformFlavorContainerRuntimeVM).Stop(ctx)
}
