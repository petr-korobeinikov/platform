package cmd

import (
	"context"

	"github.com/spf13/cobra"

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

		return lib.Get(ctx, args[0])
	},
}

func init() {
	libCmd.AddCommand(libGetCmd)
}
