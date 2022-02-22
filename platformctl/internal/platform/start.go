package platform

import (
	"context"

	"github.com/spf13/viper"
)

func Start(ctx context.Context) error {
	return dispatch(viper.GetString("platform.flavor.container-runtime-vm")).Start(ctx)
}
