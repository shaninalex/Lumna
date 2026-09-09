package commands

import "github.com/spf13/cobra"

func NewStatusRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Manage lists",
	}

	cmd.AddCommand(NewStatusCreateCmd())

	return cmd
}
