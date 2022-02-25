package service

import (
	"github.com/spf13/cobra"
)

func PersistentPreRunE(cmd *cobra.Command, args []string) error {
	if parent := cmd.Parent(); parent != nil {
		if parent.PersistentPreRunE != nil {
			if err := parent.PersistentPreRunE(parent, args); err != nil {
				return err
			}
		}
	}

	// Ensure platform started
	// Parse service spec
	// Create .env

	return nil
}
