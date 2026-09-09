package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/pkg/config"
	"gitlab.com/shaninalex/lumna/app/services/persistence"
)

func NewMigrateApplyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "apply",
		Short: "Apply migrations",
		Args:  cobra.ArbitraryArgs,
		Run: func(cmd *cobra.Command, args []string) {
			configPath, err := cmd.Flags().GetString("config")
			if err != nil {
				panic(err)
			}

			c := config.ProvideConfig(configPath)()
			if err := persistence.ApplyMigrations(c); err != nil {
				fmt.Println(err)
			}
		},
	}

	return cmd
}
