package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"platformctl/internal/cfg"
	"platformctl/internal/lib"
)

// libGetCmd gets a new dependent library
var libGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a new dependent library",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(cmd.Context(), cfg.TimeoutMediumOperation())
		defer cancel()

		envs := viper.GetStringSlice("go_env_vars")

		return lib.Get(ctx, args[0], envs)
	},
}

func init() {
	libCmd.AddCommand(libGetCmd)
}
