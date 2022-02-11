package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"platformctl/internal/cfg"
	"platformctl/internal/task"
)

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Perform a task",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(cmd.Context(), cfg.TimeoutHeavyOperation())
		defer cancel()

		if len(args) == 0 {
			return task.List(ctx)
		}

		return task.Perform(ctx, args)
	},
}

func init() {
	rootCmd.AddCommand(taskCmd)
}
