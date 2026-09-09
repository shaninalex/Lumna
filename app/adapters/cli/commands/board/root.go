package board

import "github.com/spf13/cobra"

func NewBoardsRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "boards",
		Short: "Manage boards",
	}

	cmd.AddCommand(NewBoardsCreateCmd())

	return cmd
}
