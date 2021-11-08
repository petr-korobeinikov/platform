package cmd

import (
	"context"

	"github.com/spf13/cobra"

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

		return lib.Sync(ctx)
	},
}

func init() {
	libCmd.AddCommand(libSyncCmd)
}
