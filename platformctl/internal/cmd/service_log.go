package cmd

import (
	"github.com/spf13/cobra"

	"platformctl/internal/service"
)

var serviceLogCmd = &cobra.Command{
	Use:   "log",
	Short: "Log from current service",
	RunE: func(cmd *cobra.Command, args []string) error {
		return service.Log(cmd.Context())
	},
}

func init() {
	serviceCmd.AddCommand(serviceLogCmd)
}
