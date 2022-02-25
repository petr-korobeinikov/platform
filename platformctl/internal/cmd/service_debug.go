package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"platformctl/internal/cfg"
	"platformctl/internal/service"
)

var serviceDebugCmd = &cobra.Command{
	Use:   "debug",
	Short: "Debug a service",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(cmd.Context(), cfg.TimeoutMediumOperation())
		defer cancel()

		return service.Debug(ctx)
	},
}

func init() {
	serviceCmd.AddCommand(serviceDebugCmd)
}
