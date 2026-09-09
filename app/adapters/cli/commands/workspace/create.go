package workspace

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/config"
	"gitlab.com/shaninalex/lumna/app/repositories"
	"gitlab.com/shaninalex/lumna/app/services/persistence"
)

func NewWorkspaceCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [title] [owner email]",
		Short: "Create wokrspace",
		Long:  "Require 2 argument - title and owner email",
		Run: func(cmd *cobra.Command, args []string) {
			flags := getIdentitiesCreateCmd(cmd)
			ctx, appCancel := context.WithCancel(context.Background())
			defer appCancel()

			provideConfig := config.ProvideConfig(flags.configPath)
			db := persistence.ProvideDB(provideConfig())
			wpRepository := repositories.NewGormWorkspaceRepository(db)

			workspace := &models.Workspace{
				Title:      flags.title,
				OwnerEmail: flags.ownerEmail,
				Active:     true,
			}

			if err := wpRepository.Create(ctx, workspace); err != nil {
				panic(err)
			}

			fmt.Printf("Workspace: %s with owner email %s created!\n", workspace.Title, workspace.OwnerEmail)
		},
	}

	cmd.PersistentFlags().String("title", "", "Workspace title")
	cmd.PersistentFlags().String("owner_email", "", "Workspace owner email")

	return cmd
}

type flags struct {
	configPath string
	title      string
	ownerEmail string
}

func getIdentitiesCreateCmd(cmd *cobra.Command) flags {
	configPath, err := cmd.Flags().GetString("config")
	if err != nil {
		panic(err)
	}
	title, err := cmd.Flags().GetString("title")
	if err != nil {
		panic(err)
	}

	ownerEmail, err := cmd.Flags().GetString("owner_email")
	if err != nil {
		panic(err)
	}
	return flags{
		configPath: configPath,
		title:      title,
		ownerEmail: ownerEmail,
	}
}
