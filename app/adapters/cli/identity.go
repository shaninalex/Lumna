package cli

import (
	"errors"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/actor"
	"gitlab.com/shaninalex/lumna/app/core/bus"
	"gitlab.com/shaninalex/lumna/app/core/errs"
	"gitlab.com/shaninalex/lumna/app/modules/identity/contract"
)

func NewIdentitiesRootCmd(app core.Resolve) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "identity",
		Short: "Manage identities",
	}

	cmd.AddCommand(newIdentitiesListCmd(app))
	cmd.AddCommand(newCreateIdentityCmd(app))

	return cmd
}

var (
	FailedToAddIdentityToWorkspaceError = errors.New("failed to add identity to workspace")
)

func newCreateIdentityCmd(app core.Resolve) *cobra.Command {
	var email, fullName, password string
	var active bool

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create identity",
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := actor.With(cmd.Context(), actor.System())
			profile, err := contract.ExecRegister(ctx, app(), contract.Register{
				Email:    email,
				FullName: fullName,
				Password: bus.Secret(password),
			})
			if err != nil {
				return errs.Conflict("WSP", "unable to add identity to workspace")
			}

			fmt.Printf("Identity created: %s [%d]\n", profile.Email, profile.ID)

			return nil
		},
	}

	cmd.Flags().StringVar(&email, "email", "", "Identity email")
	cmd.Flags().StringVar(&fullName, "full-name", "", "Identity full name")
	cmd.Flags().StringVar(&password, "password", "", "Identity raw password")
	cmd.Flags().BoolVar(&active, "active", false, "Identity is active")

	return cmd
}

func newIdentitiesListCmd(app core.Resolve) *cobra.Command {
	var limit, offset int

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List identities",
		RunE: func(cmd *cobra.Command, _ []string) error {
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
