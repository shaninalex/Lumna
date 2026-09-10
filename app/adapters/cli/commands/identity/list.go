package identity

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewIdentitiesListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List identities",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("List identities command")
		},
	}

	return cmd
}
