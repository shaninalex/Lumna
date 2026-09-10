package bootstrap

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna"
	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/platform/config"
	"gitlab.com/shaninalex/lumna/app/platform/database"
	"gorm.io/gorm"
)

func migrateCmd(app core.Resolve) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Migrate db schema",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("migrate")
		},
	}

	cmd.AddCommand(NewMigrateApplyCmd())
	cmd.AddCommand(NewMigrateCreateCmd())

	return cmd
}

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

			cfg := config.ProvideConfig(configPath)
			db := database.New(cfg)

			if err := MigrateSQLite(db.From(cmd.Context())); err != nil {
				panic(err)
			}
		},
	}

	return cmd
}

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

func MigrateSQLite(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sourceDriver, err := iofs.New(lumna.StaticFS("resources/migrations/sqlite"), ".")
	if err != nil {
		return err
	}

	dbDriver, err := sqlite3.WithInstance(sqlDB, &sqlite3.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance(
		"iofs",
		sourceDriver,
		"sqlite3",
		dbDriver,
	)
	if err != nil {
		return err
	}

	err = m.Up()

	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	fmt.Println("SQLite database migrated")
	return nil
}
