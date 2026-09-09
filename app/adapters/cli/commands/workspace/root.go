package workspace

import "github.com/spf13/cobra"

func NewWorkspaceRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "workspace",
		Short: "Manage workspaces",
	}

	cmd.AddCommand(NewWorkspaceCreateCmd())

	return cmd
}
