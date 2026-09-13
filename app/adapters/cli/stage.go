package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/actor"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
)

func NewStageRootCmd(app core.Resolve) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stage",
		Short: "Manage stages",
	}

	cmd.AddCommand(newCreateStageCmd(app))
	return cmd
}

func newCreateStageCmd(app core.Resolve) *cobra.Command {
	var scopeID int
	var name string
	var description string
	var category string
	var position float64
	var wipLimit int

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create stage",
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := actor.With(cmd.Context(), actor.System())
			params := contract.CreateStage{
				ScopeID:     scopeID,
				Name:        name,
				Description: description,
				Category:    category,
				Position:    position,
			}
			if wipLimit > 0 {
				params.WIPLimit = &wipLimit
			}

			stage, err := contract.ExecCreateStage(ctx, app(), params)
			if err != nil {
				return err
			}
			fmt.Printf("Stage \"%s\" [id:%d] created.\n", stage.Name, stage.Id)
			return nil
		},
	}

	cmd.Flags().StringVar(&name, "name", "", "Scope name")
	cmd.Flags().StringVar(&description, "description", "", "Scope description")
	cmd.Flags().IntVar(&scopeID, "scope-id", 0, "Scope project id")
	cmd.Flags().StringVar(&category, "category", "", "Stage category")
	cmd.Flags().Float64Var(&position, "position", 0, "Stage position")
	cmd.Flags().IntVar(&wipLimit, "wip-limit", 0, "Stage wip limit")

	return cmd
}
