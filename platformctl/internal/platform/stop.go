package platform

import (
	"context"

	"github.com/spf13/viper"
)

func Stop(ctx context.Context) error {
	return dispatch(viper.GetString("platform.flavor.container-runtime-vm")).Stop(ctx)
}
