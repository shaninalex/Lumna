package identity

import (
	"github.com/spf13/cobra"
)

func NewIdentitiesRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "identities",
		Short: "Manage identities",
	}

	cmd.AddCommand(NewIdentitiesCreateCmd())
	cmd.AddCommand(NewIdentitiesDeleteCmd())
	cmd.AddCommand(NewIdentitiesListCmd())

	return cmd
}
