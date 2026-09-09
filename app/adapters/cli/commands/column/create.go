package column

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/config"
	"gitlab.com/shaninalex/lumna/app/repositories"
	"gitlab.com/shaninalex/lumna/app/services/persistence"
)

func NewColumnCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [title] [board id]",
		Short: "Create column",
		Long:  "Require 2 argument - title and board email",
		Run: func(cmd *cobra.Command, args []string) {
			configPath, err := cmd.Flags().GetString("config")
			if err != nil {
				panic(err)
			}

			ctx, appCancel := context.WithCancel(context.Background())
			defer appCancel()

			provideConfig := config.ProvideConfig(configPath)
			db := persistence.ProvideDB(provideConfig())

			title, err := cmd.Flags().GetString("title")
			if err != nil {
				panic(err)
			}

			boardID, err := cmd.Flags().GetInt("board_id")
			if err != nil {
				panic(err)
			}

			wpRepository := repositories.NewGormColumnRepository(db)

			column := &models.Column{
				Title:   title,
				BoardId: boardID,
			}

			if err := wpRepository.Save(ctx, column); err != nil {
				panic(err)
			}

			fmt.Printf("Column: %s, id: %d, board Id: %d\n", column.Title, column.ID, column.BoardId)
		},
	}

	cmd.PersistentFlags().String("title", "", "Column title")
	cmd.PersistentFlags().Int("board_id", 0, "Board id")

	return cmd
}
