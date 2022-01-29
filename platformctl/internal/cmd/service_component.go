package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"platformctl/internal/cfg"
	"platformctl/internal/service"
)

var serviceComponentCmd = &cobra.Command{
	Use:   "component",
	Short: "A list of enabled service components",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(cmd.Context(), cfg.TimeoutDefault())
		defer cancel()

		return service.Component(ctx)
	},
}

func init() {
	serviceCmd.AddCommand(serviceComponentCmd)
}
