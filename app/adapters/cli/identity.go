package cli

import (
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/actor"
	"gitlab.com/shaninalex/lumna/app/modules/identity/contract"
)

func NewIdentitiesRootCmd(app core.Resolve) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "identities",
		Short: "Manage identities",
	}

	cmd.AddCommand(newIdentitiesListCmd(app))

	return cmd
}

func newIdentitiesListCmd(app core.Resolve) *cobra.Command {
	var limit, offset int

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List identities",
		RunE: func(cmd *cobra.Command, _ []string) error {
			// The CLI acts on behalf of the system — there are no user permissions here.
			ctx := actor.With(cmd.Context(), actor.System())

			page, err := contract.AskListProfiles(ctx, app(), contract.ListProfiles{
				Limit:  limit,
				Offset: offset,
			})
			if err != nil {
				return err
			}

			return printProfiles(cmd.OutOrStdout(), page.Profiles)
		},
	}

	cmd.Flags().IntVar(&limit, "limit", 10, "Limit identities amount in list")
	cmd.Flags().IntVar(&offset, "offset", 0, "Pagination offset")

	return cmd
}

func printProfiles(out io.Writer, profiles []contract.ProfileView) error {
	w := tabwriter.NewWriter(out, 0, 0, 2, ' ', 0)
	if _, err := fmt.Fprintln(w, "ID\tEMAIL\tNAME\tACTIVE"); err != nil {
		return err
	}
	for _, p := range profiles {
		if _, err := fmt.Fprintf(w, "%d\t%s\t%s\t%t\n", p.ID, p.Email, p.FullName, p.Active); err != nil {
			return err
		}
	}
	return w.Flush()
}
