package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"platformctl/internal/cfg"
	"platformctl/internal/platform"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start container runtime",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(cmd.Context(), cfg.TimeoutHeavyOperation())
		defer cancel()

		return platform.Start(ctx)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
