package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/actor"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/contract"
)

func NewProjectRootCmd(app core.Resolve) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "project",
		Short: "Manage projects",
	}

	cmd.AddCommand(newCreateProjectCmd(app))

	return cmd
}

func newCreateProjectCmd(app core.Resolve) *cobra.Command {
	var title string
	var ownerId, workspaceId int

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create project",
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := actor.With(cmd.Context(), actor.System())
			project, err := contract.ExecCreateProject(ctx, app(), contract.CreateProject{
				Title:       title,
				OwnerId:     ownerId,
				WorkspaceId: workspaceId,
			})
			if err != nil {
				return err
			}

			fmt.Printf("Project \"%s\" [id:%d] created.\n", project.Title, project.Id)

			return nil
		},
	}

	cmd.Flags().StringVar(&title, "title", "", "Workspace title")
	cmd.Flags().IntVar(&ownerId, "owner-id", 0, "Workspace Id belongs to")
	cmd.Flags().IntVar(&workspaceId, "workspace-id", 0, "Workspace Id belongs to")

	return cmd
}
