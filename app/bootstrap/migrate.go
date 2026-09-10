package bootstrap

import (
	"fmt"

	"github.com/spf13/cobra"
)

func migrateCmd(app *App) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Migrate db schema",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("migrate")
		},
	}

	return cmd
}
