package bus

import (
	"context"
	"errors"
	"fmt"
	"reflect"
)

type QueryBus struct {
	handlers map[reflect.Type]Invoke
	chain    []Middleware
	sealed   bool
}

func NewQueryBus(mw ...Middleware) *QueryBus {
	return &QueryBus{
		handlers: map[reflect.Type]Invoke{},
		chain:    mw,
	}
}

func RegisterQuery[C any, R any](b *QueryBus, h func(context.Context, C) (R, error)) error {
	if b.sealed {
		return errors.New("bus: registration after Seal()")
	}
	key := reflect.TypeFor[C]()
	if _, dup := b.handlers[key]; dup {
		return fmt.Errorf("bus: handler for %s already registered", key)
	}
	call := Invoke(func(ctx context.Context, msg any) (any, error) {
		return h(ctx, msg.(C))
	})
	for i := len(b.chain) - 1; i >= 0; i-- {
		call = b.chain[i](call)
	}
	b.handlers[key] = call
	return nil
}

// Seal - close registration
func (b *QueryBus) Seal() { b.sealed = true }

func Ask[C any, R any](ctx context.Context, b *QueryBus, cmd C) (R, error) {
	var zero R
	call, ok := b.handlers[reflect.TypeFor[C]()]
	if !ok {
		return zero, fmt.Errorf("bus: no handler for %T", cmd)
	}
	res, err := call(ctx, cmd)
	if err != nil {
		return zero, err
	}
	typed, ok := res.(R)
	if !ok {
		return zero, fmt.Errorf("bus: %T returned %T, want %T", cmd, res, zero)
	}
	return typed, nil
}
