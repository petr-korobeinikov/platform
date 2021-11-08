package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"platformctl/internal/cfg"
	"platformctl/internal/docker"
	"platformctl/internal/minikube"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start minikube",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(cmd.Context(), cfg.TimeoutHeavyOperation())
		defer cancel()

		if _, err := docker.IsConformsToMinimalRequirements(ctx); err != nil {
			return err
		}

		return minikube.Start(ctx)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
