package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/actor"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
)

func NewBoardRootCmd(app core.Resolve) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "board",
		Short: "Manage boards",
	}

	cmd.AddCommand(newCreateBoardCmd(app))

	return cmd
}

func newCreateBoardCmd(app core.Resolve) *cobra.Command {
	var title string
	var projectId int

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create board",
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := actor.With(cmd.Context(), actor.System())

			board, err := contract.ExecCreateBoard(ctx, app(), contract.CreateBoard{
				Title:     title,
				ProjectId: projectId,
			})

			if err != nil {
				return err
			}

			fmt.Printf(
				"Board \"%s\" [id:%d]created.\n",
				board.Title,
				board.Id,
			)

			return nil
		},
	}

	cmd.Flags().StringVar(&title, "title", "", "Board title")
	cmd.Flags().IntVar(&projectId, "project-id", 0, "Board project id")

	return cmd
}
