package cmd

import (
	"context"
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"platformctl/internal/cfg"
	"platformctl/internal/lib"
)

// libGetCmd gets a new dependent library, upgrades, or downgrades the existing one
var libGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a new dependent library",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return ErrTooFewArguments
		}

		if len(args) > 2 {
			return ErrTooLongArgumentList
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(cmd.Context(), cfg.TimeoutMediumOperation())
		defer cancel()

		envs := viper.GetStringSlice("go_env_vars")

		version := "latest"
		if len(args) == 2 {
			version = args[1]
		}

		return lib.Get(ctx, envs, args[0], version)
	},
}

func init() {
	libCmd.AddCommand(libGetCmd)
}

var (
	ErrTooFewArguments     = errors.New(`too few arguments`)
	ErrTooLongArgumentList = errors.New(`too long argument list`)
)
