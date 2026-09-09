package commands

import "github.com/spf13/cobra"

func NewEmailRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "email",
		Short: "Manage emails",
	}

	cmd.AddCommand(NewEmailCreateCmd())

	return cmd
}
