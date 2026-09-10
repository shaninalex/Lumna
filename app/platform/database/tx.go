package database

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

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
