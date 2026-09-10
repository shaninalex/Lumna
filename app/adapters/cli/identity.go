package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/modules/identity/contract"
)

func NewIdentitiesRootCmd(app Resolve) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "identities",
		Short: "Manage identities",
	}

	cmd.AddCommand(newIdentitiesListCmd(app))

	return cmd
}

func newIdentitiesListCmd(app Resolve) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List identities",
		Run: func(cmd *cobra.Command, args []string) {
			limit, err := cmd.Flags().GetInt("limit")
			if err != nil {
				panic(err)
			}
			offset, err := cmd.Flags().GetInt("offset")
			if err != nil {
				panic(err)
			}

			page, err := contract.AskListProfiles(cmd.Context(), app(), contract.ListProfiles{
				Limit:  limit,
				Offset: offset,
			})
			if err != nil {
				panic(err)
			}

			fmt.Println(page)
		},
	}

	cmd.PersistentFlags().Int("limit", 10, "Limit identities amount in list")
	cmd.PersistentFlags().Int("offset", 0, "Pagination offset")

	return cmd
}
