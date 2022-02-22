package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"platformctl/internal/bootstrap"
	"platformctl/internal/cfg"
)

var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Bootstrap platform configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(cmd.Context(), cfg.TimeoutDefault())
		defer cancel()

		return bootstrap.Bootstrap(ctx, bootstrapURL)
	},
}

func init() {
	bootstrapCmd.Flags().StringVarP(&bootstrapURL, "bootstrap-url", "", "", "A bootstrap URL with config contents")
	bootstrapCmd.MarkFlagRequired("bootstrap-url")

	rootCmd.AddCommand(bootstrapCmd)
}

var (
	bootstrapURL string
)
