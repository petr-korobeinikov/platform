package platform

import (
	"context"

	"platformctl/internal/cfg"
)

func Start(ctx context.Context) error {
	return dispatch(cfg.PlatformFlavorContainerRuntimeVM).Start(ctx)
}
