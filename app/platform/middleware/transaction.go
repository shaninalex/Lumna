package middleware

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/core/bus"
	"gitlab.com/shaninalex/lumna/app/platform/database"
)

func Transaction(db *database.DB) bus.Middleware {
	return func(next bus.Invoke) bus.Invoke {
		return func(ctx context.Context, msg any) (any, error) {
			return db.InTx(ctx, func(txCtx context.Context) (any, error) {
				return next(txCtx, msg)
			})
		}
	}
}
