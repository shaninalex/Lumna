package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/actor"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/contract"
)

func NewWorkspaceRootCmd(app core.Resolve) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "workspace",
		Short: "Manage workspaces",
	}

	cmd.AddCommand(newCreateWorkspaceCmd(app))

	return cmd
}

func newCreateWorkspaceCmd(app core.Resolve) *cobra.Command {
	var title, ownerEmail string
	var active bool

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create workspace",
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := actor.With(cmd.Context(), actor.System())
			workspace, err := contract.ExecCreateWorkspace(ctx, app(), contract.CreateWorkspace{
				Title:      title,
				OwnerEmail: ownerEmail,
				Active:     active,
			})

			if err != nil {
				return err
			}

			fmt.Printf(
				"Workspace \"%s\" [id:%d] with owner email %s was successfully created.\n",
				workspace.Title,
				workspace.Id,
				workspace.OwnerEmail,
			)

			return nil
		},
	}

	cmd.Flags().StringVar(&title, "title", "", "Workspace title")
	cmd.Flags().StringVar(&ownerEmail, "owner-email", "", "Workspace owner email")
	cmd.Flags().BoolVar(&active, "active", false, "Workspace is active")

	return cmd
}
