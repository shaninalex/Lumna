package database

import (
	"context"
	"errors"

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

type txKey struct{}

func (d *DB) From(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(txKey{}).(*gorm.DB); ok {
		return tx
	}
	return d.gorm.WithContext(ctx)
}

func (d *DB) InTx(ctx context.Context, fn func(context.Context) (any, error)) (any, error) {
	if _, nested := ctx.Value(txKey{}).(*gorm.DB); nested {
		return nil, errors.New("database: nested transaction")
	}
	var out any
	err := d.gorm.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var e error
		out, e = fn(context.WithValue(ctx, txKey{}, tx))
		return e
	})
	return out, err
}
