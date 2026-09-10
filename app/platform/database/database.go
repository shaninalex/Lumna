package database

import (
	"gitlab.com/shaninalex/lumna/app/platform/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
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
	sqlDB, err := gdb.DB()
	if err != nil {
		panic(err)
	}
	if _, err := sqlDB.Exec("PRAGMA foreign_keys = ON"); err != nil {
		panic(err)
	}
	// ============= =============  ============= =============

	return &DB{
		db: gdb,
	}
}

func connectionOptions(conf *config.Config) *gorm.Config {
	opt := &gorm.Config{}
	//if conf.Env() != config.EnvironmentDev && conf.Env() != config.EnvironmentTest {
	//	opt.Logger = silentLogger()
	//}
	return opt
}
