package cmd

import "github.com/spf13/cobra"

var libCmd = &cobra.Command{
	Use:   "lib",
	Short: "Operations with project libraries",
}

func init() {
	rootCmd.AddCommand(libCmd)
}
