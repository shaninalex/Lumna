package commands

import (
	"log"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/config"
	"gitlab.com/shaninalex/lumna/app/services/persistence"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

func NewStatusCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [list_id] [title]",
		Short: "Create status",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			c := dig.New()
			configPath, err := cmd.Flags().GetString("config")
			if err != nil {
				panic(err)
			}
			_ = c.Provide(config.ProvideConfig(configPath))
			_ = c.Provide(persistence.ProvideDB)

			//listId, err := strconv.Atoi(args[0])
			if err != nil {
				panic(err)
			}
			title := args[1]
			column := models.Column{
				Title: title,
				//ListID: int(listId),
			}
			if err := c.Invoke(createColumn(column)); err != nil {
				panic(err)
			}
		},
	}

	return cmd
}

func createColumn(column models.Column) func(db *gorm.DB) {
	return func(db *gorm.DB) {

		if result := db.Create(&column); result.Error != nil {
			panic(result.Error)
		}

		log.Println("Column created with ID:", column.ID)
	}
}
