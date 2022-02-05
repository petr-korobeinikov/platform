package cmd

import (
	"github.com/spf13/cobra"

	"platformctl/internal/service"
)

var serviceDocCmd = &cobra.Command{
	Use:   "doc",
	Short: "Service documentation",
	RunE: func(cmd *cobra.Command, args []string) error {
		return service.Doc(cmd.Context())
	},
}

func init() {
	serviceCmd.AddCommand(serviceDocCmd)
}
