package cmd

import "github.com/spf13/cobra"

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Operations with services",
}

func init() {
	rootCmd.AddCommand(serviceCmd)
}
