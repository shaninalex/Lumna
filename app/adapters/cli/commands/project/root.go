package project

import "github.com/spf13/cobra"

func NewProjectsRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "projects",
		Short: "Manage projects",
	}

	cmd.AddCommand(NewProjectsCreateCommand())
	cmd.AddCommand(NewProjectsListCommand())

	return cmd
}
