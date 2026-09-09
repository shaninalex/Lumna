package project

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/config"
	"gitlab.com/shaninalex/lumna/app/pkg/utils"
	"gitlab.com/shaninalex/lumna/app/repositories"
	"gitlab.com/shaninalex/lumna/app/services/persistence"
)

func NewProjectsCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [project title] [workspace id] [owner id]",
		Short: "Create project",
		Long:  "Require 3 arguments - project title, workspace id, owner id",
		Run: func(cmd *cobra.Command, args []string) {
			flags := getProjectsCreateCommandFlags(cmd)
			ctx, appCancel := context.WithCancel(context.Background())
			defer appCancel()

			conf := config.ProvideConfig(flags.configPath)()
			db := persistence.ProvideDB(conf)
			projectRepository := repositories.NewGormProjectRepository(db)

			project := &models.Project{
				Title:       flags.title,
				WorkspaceID: flags.workspaceId,
				OwnerID:     utils.Pointer(flags.ownerId),
			}
			if err := projectRepository.Create(ctx, project); err != nil {
				panic(err)
			}

			fmt.Printf("Project: %s created! ID: %d\n", project.Title, project.ID)
		},
	}

	cmd.PersistentFlags().String("title", "", "Column title")
	cmd.PersistentFlags().Int("workspace_id", 0, "Workspace id")
	cmd.PersistentFlags().Int("owner_id", 0, "Owner id")

	return cmd
}

type flags struct {
	configPath  string
	title       string
	workspaceId int
	ownerId     int
}

func getProjectsCreateCommandFlags(cmd *cobra.Command) flags {
	configPath, err := cmd.Flags().GetString("config")
	if err != nil {
		panic(err)
	}

	title, err := cmd.Flags().GetString("title")
	if err != nil {
		panic(err)
	}
	workspaceId, err := cmd.Flags().GetInt("workspace_id")
	if err != nil {
		panic(err)
	}

	ownerId, err := cmd.Flags().GetInt("owner_id")
	if err != nil {
		panic(err)
	}

	return flags{
		configPath:  configPath,
		title:       title,
		workspaceId: workspaceId,
		ownerId:     ownerId,
	}
}
