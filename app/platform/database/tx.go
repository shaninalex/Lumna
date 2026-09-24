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

type hooksKey struct{}

func AfterCommit(ctx context.Context, fn func()) {
	if h, ok := ctx.Value(hooksKey{}).(*[]func()); ok {
		*h = append(*h, fn)
		return
	}
	fn()
}

func (d *DB) InTx(ctx context.Context, fn func(context.Context) (any, error)) (any, error) {
	if _, nested := ctx.Value(txKey{}).(*gorm.DB); nested {
		return nil, errors.New("database: nested transaction")
	}
	var out any
	var hooks []func()
	ctx = context.WithValue(ctx, hooksKey{}, &hooks)
	err := d.gorm.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var e error
		out, e = fn(context.WithValue(ctx, txKey{}, tx))
		return e
	})

	if err == nil {
		for _, h := range hooks {
			h()
		}
	}
	return out, err
}
