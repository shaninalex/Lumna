package project

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/config"
	"gitlab.com/shaninalex/lumna/app/services/persistence"
	"gorm.io/gorm"
)

func NewProjectsListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List projects",
		Run: func(cmd *cobra.Command, args []string) {
			configPath, err := cmd.Flags().GetString("config")
			if err != nil {
				panic(err)
			}

			conf := config.ProvideConfig(configPath)()
			db := persistence.ProvideDB(conf)
			projectsList(db)
		},
	}

	return cmd
}

func projectsList(db *gorm.DB) {
	projects := []models.Project{}

	if result := db.Find(&projects); result.Error != nil {
		panic(result.Error)
	}

	for i, p := range projects {
		fmt.Println(i, p.String())
		boards := make([]models.Board, 0)
		if err := db.Find(&boards).Where("project_id = ?", p.ID).Find(&boards); err != nil {
			continue
		}
		for _, b := range boards {
			fmt.Println("\t\t", b.String())
		}
	}

	fmt.Println("Total projects: ", len(projects))
}
