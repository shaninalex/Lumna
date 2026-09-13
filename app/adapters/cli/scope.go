package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/actor"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
)

func NewScopeRootCmd(app core.Resolve) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "scope",
		Short: "Manage scopes",
	}

	cmd.AddCommand(newCreateScopeCmd(app))

	return cmd
}

func newCreateScopeCmd(app core.Resolve) *cobra.Command {
	var name string
	var projectId int

	cmd := &cobra.Command{
		Use:   "scope",
		Short: "Create scope",
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := actor.With(cmd.Context(), actor.System())
			scope, err := contract.ExecCreateScope(ctx, app(), contract.CreateScope{
				Name:      name,
				ProjectId: projectId,
			})

			if err != nil {
				return err
			}

			fmt.Printf(
				"Scope \"%s\" [id:%d]created.\n",
				scope.Name,
				scope.Id,
			)

			return nil
		},
	}

	cmd.Flags().StringVar(&name, "name", "", "Scope name")
	cmd.Flags().IntVar(&projectId, "project-id", 0, "Scope project id")

	return cmd
}
