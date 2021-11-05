package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"platformctl/internal/action/lib"
	"platformctl/internal/cfg"
)

// libGetCmd gets a new dependent library
var libGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a new dependent library",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(cmd.Context(), cfg.TimeoutMediumOperation())
		defer cancel()

		return lib.Get(ctx)
	},
}

func init() {
	libCmd.AddCommand(libGetCmd)
}
