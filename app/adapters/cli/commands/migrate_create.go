package commands

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

const (
	migrationDateFormat    = "20060102150405"
	migrationsSqlitePath   = "./resources/migrations/sqlite"
	migrationsPostgresPath = "./resources/migrations/postgres"
)

func NewMigrateCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create new migration",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			filename := fmt.Sprintf("%s_%s", time.Now().Format(migrationDateFormat), args[0])
			migrationFiles := []string{}
			for _, p := range []string{migrationsSqlitePath, migrationsPostgresPath} {
				for _, d := range []string{"up", "down"} {
					mf := fmt.Sprintf("%s/%s.%s.sql", p, filename, d)
					_, err := os.OpenFile(mf, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
					if err != nil {
						log.Fatal(err)
					}
					migrationFiles = append(migrationFiles, mf)
				}
			}

			fmt.Printf("Migration %s created\n", args[0])
			for _, f := range migrationFiles {
				fmt.Println(f)
			}
		},
	}

	return cmd
}
