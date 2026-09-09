package commands

import (
	"github.com/spf13/cobra"
)

func NewMigrateRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Migrate database",
	}

	cmd.AddCommand(NewMigrateApplyCmd())
	cmd.AddCommand(NewMigrateCreateCmd())

	return cmd
}
