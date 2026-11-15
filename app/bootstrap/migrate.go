package bootstrap

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

func migrateCmd(app func() *Instance, assets Assets) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Migrate db schema",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("migrate")
		},
	}

	cmd.AddCommand(NewMigrateApplyCmd(app, assets))
	cmd.AddCommand(NewMigrateCreateCmd())

	return cmd
}

func NewMigrateApplyCmd(app func() *Instance, assets Assets) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "apply",
		Short: "Apply migrations",
		Args:  cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return MigrateSQLite(app().DB.From(cmd.Context()), assets.Migrations)
		},
	}

	return cmd
}

const (
	migrationDateFormat  = "20060102150405"
	migrationsSqlitePath = "./resources/migrations"
)

func NewMigrateCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create new migration",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			filename := fmt.Sprintf("%s_%s", time.Now().Format(migrationDateFormat), args[0])
			migrationFiles := []string{}
			for _, p := range []string{migrationsSqlitePath} {
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

func MigrateSQLite(db *gorm.DB, migrateFs fs.FS) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sourceDriver, err := iofs.New(migrateFs, ".")
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
