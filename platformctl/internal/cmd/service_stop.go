package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"platformctl/internal/cfg"
	"platformctl/internal/service"
)

var serviceStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop current service",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(cmd.Context(), cfg.TimeoutMediumOperation())
		defer cancel()

		// Temporary disable minikube status checking
		// if _, err := minikube.IsRunning(ctx); err != nil {
		// 	return err
		// }

		return service.Stop(ctx)
	},
}

func init() {
	serviceCmd.AddCommand(serviceStopCmd)
}
