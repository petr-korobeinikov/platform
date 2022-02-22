package cmd

import (
	"github.com/spf13/cobra"

	"platformctl/internal/service"
)

var serviceCmd = &cobra.Command{
	Use:               "service",
	Short:             "Operations with services",
	PersistentPreRunE: service.PersistentPreRunE,
}

func init() {
	rootCmd.AddCommand(serviceCmd)
}
