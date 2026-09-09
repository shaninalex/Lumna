package column

import "github.com/spf13/cobra"

func NewColumnRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "column",
		Short: "Manage columns",
	}

	cmd.AddCommand(NewColumnCreateCmd())

	return cmd
}
