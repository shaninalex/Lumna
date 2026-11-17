package bus

import "context"

type QueryBus struct{ registry }

func NewQueryBus(mw ...Middleware) *QueryBus {
	return &QueryBus{newRegistry(mw)}
}

func RegisterQuery[Q any, R any](b *QueryBus, h func(context.Context, Q) (R, error)) error {
	return register(&b.registry, h)
}

func Ask[Q any, R any](ctx context.Context, b *QueryBus, q Q) (R, error) {
	return dispatch[Q, R](ctx, &b.registry, q)
}
