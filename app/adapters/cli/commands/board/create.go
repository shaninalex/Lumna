package board

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/config"
	"gitlab.com/shaninalex/lumna/app/repositories"
	"gitlab.com/shaninalex/lumna/app/services/persistence"
)

func NewBoardsCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [title] [project id]",
		Short: "Create board",
		Long:  "Require 2 arguments - title and project id",
		Run: func(cmd *cobra.Command, args []string) {
			flags := getBoardsCreateCmdFlags(cmd)
			ctx, appCancel := context.WithCancel(context.Background())
			defer appCancel()

			provideConfig := config.ProvideConfig(flags.configPath)
			db := persistence.ProvideDB(provideConfig())

			board := models.Board{
				Title:     flags.title,
				ProjectID: flags.projectId,
			}

			repository := repositories.NewGormBoardRepository(db)
			if err := repository.Create(ctx, &board); err != nil {
				panic(board)
			}
			fmt.Printf("Board: %s, id: %d, project Id: %d\n", board.Title, board.ID, board.ProjectID)
		},
	}

	cmd.PersistentFlags().String("title", "", "Column title")
	cmd.PersistentFlags().Int("project_id", 0, "Project id")

	return cmd
}

type flags struct {
	configPath string
	title      string
	projectId  int
}

func getBoardsCreateCmdFlags(cmd *cobra.Command) flags {
	configPath, err := cmd.Flags().GetString("config")
	if err != nil {
		panic(err)
	}

	title, err := cmd.Flags().GetString("title")
	if err != nil {
		panic(err)
	}

	projectId, err := cmd.Flags().GetInt("project_id")
	if err != nil {
		panic(err)
	}

	return flags{
		configPath: configPath,
		title:      title,
		projectId:  projectId,
	}
}
