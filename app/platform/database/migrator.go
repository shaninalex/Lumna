package database

import (
	"errors"
	"fmt"
	"io/fs"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"gorm.io/gorm"
)

type Migrator struct {
	m *migrate.Migrate
}

func NewMigrator(db *gorm.DB, migrations fs.FS) (*Migrator, error) {
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	source, err := iofs.New(migrations, ".")
	if err != nil {
		return nil, fmt.Errorf("migrations source: %w", err)
	}
	driver, err := sqlite3.WithInstance(sqlDB, &sqlite3.Config{})
	if err != nil {
		return nil, fmt.Errorf("migrations database driver: %w", err)
	}
	m, err := migrate.NewWithInstance("iofs", source, "sqlite3", driver)
	if err != nil {
		return nil, err
	}
	return &Migrator{m: m}, nil
}

func (s *Migrator) Up() error {
	return ignoreNoChange(s.m.Up())
}

func (s *Migrator) Down(steps int) error {
	if steps <= 0 {
		return ignoreNoChange(s.m.Down())
	}
	version, _, err := s.Version()
	if err != nil {
		return err
	}
	if version == 0 {
		return nil
	}
	err = s.m.Steps(-steps)
	if _, short := errors.AsType[migrate.ErrShortLimit](err); short {
		return nil
	}
	return ignoreNoChange(err)
}

func (s *Migrator) Force(version int) error {
	return s.m.Force(version)
}

func (s *Migrator) Version() (version uint, dirty bool, err error) {
	version, dirty, err = s.m.Version()
	if errors.Is(err, migrate.ErrNilVersion) {
		return 0, false, nil
	}
	return version, dirty, err
}

func ignoreNoChange(err error) error {
	if errors.Is(err, migrate.ErrNoChange) {
		return nil
	}
	return err
}
