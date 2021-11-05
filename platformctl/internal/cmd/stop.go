package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"platformctl/internal/action/minikube"
	"platformctl/internal/cfg"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop minikube",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(cmd.Context(), cfg.TimeoutMediumOperation())
		defer cancel()

		return minikube.Stop(ctx)
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
