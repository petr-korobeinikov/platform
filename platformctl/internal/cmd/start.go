package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"platformctl/internal/action/minikube"
	"platformctl/internal/cfg"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start minikube",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(cmd.Context(), cfg.TimeoutHeavyOperation())
		defer cancel()

		return minikube.Start(ctx)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
