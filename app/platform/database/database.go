package database

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/platform/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	gorm *gorm.DB
}

func New(cfg *config.Config) *DB {
	gdb, err := gorm.Open(
		sqlite.Open(cfg.String("database.url")),
		connectionOptions(cfg),
	)

	if err != nil {
		panic("failed to connect sqlite database: " + err.Error())
	}

	// ============ Enable foreign keys for SQLite ============
	result := gorm.WithResult()
	if err := gorm.G[any](gdb, result).Exec(context.Background(), "PRAGMA foreign_keys = ON"); err != nil {
		panic(err)
	}
	// ============= =============  ============= =============

	return &DB{
		gorm: gdb,
	}
}

func connectionOptions(conf *config.Config) *gorm.Config {
	opt := &gorm.Config{}
	//if conf.Env() != config.EnvironmentDev && conf.Env() != config.EnvironmentTest {
	//	opt.Logger = silentLogger()
	//}
	return opt
}
