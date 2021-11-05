package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"platformctl/internal/action/service"
	"platformctl/internal/cfg"
)

var serviceCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a service",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(cmd.Context(), cfg.TimeoutMediumOperation())
		defer cancel()

		return service.Create(ctx, args[0])
	},
}

func init() {
	serviceCmd.AddCommand(serviceCreateCmd)
}
