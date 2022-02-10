package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"platformctl/internal/cfg"
	"platformctl/internal/service"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Explain and export service environment",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(cmd.Context(), cfg.TimeoutDefault())
		defer cancel()

		return service.Env(ctx, cfg.ServiceEnv)
	},
}

func init() {
	serviceCmd.AddCommand(envCmd)
}

var (
	ServiceEnv string
)
