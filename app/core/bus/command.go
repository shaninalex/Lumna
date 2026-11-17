package bus

import "context"

type CommandBus struct{ registry }

func NewCommandBus(mw ...Middleware) *CommandBus {
	return &CommandBus{newRegistry(mw)}
}

func RegisterCommand[C any, R any](b *CommandBus, h func(context.Context, C) (R, error)) error {
	return register(&b.registry, h)
}

func Execute[C any, R any](ctx context.Context, b *CommandBus, cmd C) (R, error) {
	return dispatch[C, R](ctx, &b.registry, cmd)
}
