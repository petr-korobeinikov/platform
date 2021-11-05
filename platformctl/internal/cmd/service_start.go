package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"platformctl/internal/action/service"
	"platformctl/internal/cfg"
)

var serviceStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a service",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(cmd.Context(), cfg.TimeoutMediumOperation())
		defer cancel()

		return service.Start(ctx)
	},
}

func init() {
	serviceCmd.AddCommand(serviceStartCmd)
}
