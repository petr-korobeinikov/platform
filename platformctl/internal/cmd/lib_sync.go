package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"platformctl/internal/cfg"
	"platformctl/internal/lib"
)

// libSyncCmd runs sync for dependent libraries
var libSyncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Synchronize project libraries",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(cmd.Context(), cfg.TimeoutMediumOperation())
		defer cancel()

		envs := viper.GetStringSlice(cfgKeyGoEnvVars)

		return lib.Sync(ctx, envs)
	},
}

func init() {
	libCmd.AddCommand(libSyncCmd)
}
