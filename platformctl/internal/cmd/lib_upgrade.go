package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"platformctl/internal/cfg"
	"platformctl/internal/lib"
)

var libUpgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade project libraries",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(cmd.Context(), cfg.TimeoutMediumOperation())
		defer cancel()

		envs := viper.GetStringSlice("go_env_vars")

		return lib.Upgrade(ctx, envs)
	},
}

func init() {
	libCmd.AddCommand(libUpgradeCmd)
}
