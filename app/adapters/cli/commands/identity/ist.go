package identity

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/config"
	"gitlab.com/shaninalex/lumna/app/services/persistence"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

func NewIdentitiesListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Identities list",
		Run: func(cmd *cobra.Command, args []string) {
			c := dig.New()
			configPath, err := cmd.Flags().GetString("config")
			if err != nil {
				panic(err)
			}

			_ = c.Provide(config.ProvideConfig(configPath))
			_ = c.Provide(persistence.ProvideDB)

			if err := c.Invoke(identitiesList); err != nil {
				panic(err)
			}
		},
	}

	return cmd
}

func identitiesList(db *gorm.DB) {
	identities := []models.Identity{}
	if result := db.Find(&identities); result.Error != nil {
		panic(result.Error)
	}

	for i, identity := range identities {
		fmt.Printf("%d: %s\n", i+1, identity.String())
	}
}
