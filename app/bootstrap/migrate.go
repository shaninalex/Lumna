package bootstrap

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/platform/database"
)

func migrateCmd(app func() *Instance, assets Assets) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Migrate db schema",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return cmd.Help()
		},
	}

	migrator := func(cmd *cobra.Command) (*database.Migrator, error) {
		return database.NewMigrator(app().DB.From(cmd.Context()), assets.Migrations)
	}

	cmd.AddCommand(
		newMigrateApplyCmd(migrator),
		newMigrateDownCmd(migrator),
		newMigrateForceCmd(migrator),
		newMigrateVersionCmd(migrator),
		NewMigrateCreateCmd(),
	)

	return cmd
}

type migratorFactory func(cmd *cobra.Command) (*database.Migrator, error)

func newMigrateApplyCmd(newMigrator migratorFactory) *cobra.Command {
	return &cobra.Command{
		Use:     "apply",
		Aliases: []string{"up"},
		Short:   "Apply all pending migrations",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			m, err := newMigrator(cmd)
			if err != nil {
				return err
			}
			if err := m.Up(); err != nil {
				return err
			}
			return printVersion(cmd, m)
		},
	}
}

func newMigrateDownCmd(newMigrator migratorFactory) *cobra.Command {
	var all bool

	cmd := &cobra.Command{
		Use:   "down [N]",
		Short: "Roll back the last N migrations (default 1), or all of them with --all",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			steps := 1
			if len(args) == 1 {
				if all {
					return errors.New("migrate down: pass either N or --all, not both")
				}
				n, err := strconv.Atoi(args[0])
				if err != nil || n < 1 {
					return fmt.Errorf("migrate down: N must be a positive number, got %q", args[0])
				}
				steps = n
			}
			if all {
				steps = 0
			}

			m, err := newMigrator(cmd)
			if err != nil {
				return err
			}
			if err := m.Down(steps); err != nil {
				return err
			}
			return printVersion(cmd, m)
		},
	}

	cmd.Flags().BoolVar(&all, "all", false, "Roll back every applied migration")
	return cmd
}

func newMigrateForceCmd(newMigrator migratorFactory) *cobra.Command {
	return &cobra.Command{
		Use:   "force VERSION",
		Short: "Set the schema version and clear the dirty flag without running migrations",
		Long: "Use after fixing a failed migration by hand. The database is marked as being at VERSION " +
			"and clean; no SQL from the migration files is executed.",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			version, err := strconv.Atoi(args[0])
			if err != nil || version < 0 {
				return fmt.Errorf("migrate force: VERSION must be a non-negative number, got %q", args[0])
			}

			m, err := newMigrator(cmd)
			if err != nil {
				return err
			}
			if err := m.Force(version); err != nil {
				return err
			}
			return printVersion(cmd, m)
		},
	}
}

func newMigrateVersionCmd(newMigrator migratorFactory) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the current schema version",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			m, err := newMigrator(cmd)
			if err != nil {
				return err
			}
			return printVersion(cmd, m)
		},
	}
}

func printVersion(cmd *cobra.Command, m *database.Migrator) error {
	version, dirty, err := m.Version()
	if err != nil {
		return err
	}
	out := cmd.OutOrStdout()
	switch {
	case version == 0:
		_, err = fmt.Fprintln(out, "Schema version: none (no migrations applied)")
	case dirty:
		_, err = fmt.Fprintf(out, "Schema version: %d (dirty: fix the database, then run `migrate force %d`)\n", version, version)
	default:
		_, err = fmt.Fprintf(out, "Schema version: %d\n", version)
	}
	return err
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
