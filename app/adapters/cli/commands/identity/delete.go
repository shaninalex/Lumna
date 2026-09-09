package identity

import (
	"log"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/config"
	"gitlab.com/shaninalex/lumna/app/services/persistence"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

func NewIdentitiesDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete [id]",
		Short: "Delete identity",
		Long:  "Require 1 arguments - Identity Id to delete. For example:\nlumna identities delete <uuid>",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			c := dig.New()
			configPath, err := cmd.Flags().GetString("config")
			if err != nil {
				panic(err)
			}
			id := uuid.MustParse(args[0])
			_ = c.Provide(config.ProvideConfig(configPath))
			_ = c.Provide(persistence.ProvideDB)

			if err := c.Invoke(deleteIdentity(id)); err != nil {
				panic(err)
			}
		},
	}

	return cmd
}

func deleteIdentity(id uuid.UUID) func(database *gorm.DB) {
	return func(database *gorm.DB) {
		if result := database.Where("id = ?", id.String()).Delete(&models.Identity{}); result.Error != nil {
			log.Fatal(result.Error)
		}

		log.Println("User with id", id.String(), "deleted.")
	}
}
