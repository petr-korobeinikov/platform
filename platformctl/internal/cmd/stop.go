package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"platformctl/internal/cfg"
	"platformctl/internal/platform"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop container runtime",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(cmd.Context(), cfg.TimeoutMediumOperation())
		defer cancel()

		return platform.Stop(ctx)
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
