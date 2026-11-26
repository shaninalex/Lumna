package database

import (
	"gitlab.com/shaninalex/lumna/app/platform/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	gorm *gorm.DB
}

func (d *DB) Close() error {
	sqlDB, err := d.gorm.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func New(cfg config.Interface) *DB {
	gdb, err := gorm.Open(
		sqlite.Open(cfg.String("database.url")),
		connectionOptions(cfg),
	)

	if err != nil {
		panic("failed to connect sqlite database: " + err.Error())
	}

	return &DB{
		gorm: gdb,
	}
}

func connectionOptions(conf config.Interface) *gorm.Config {
	return &gorm.Config{}
}
