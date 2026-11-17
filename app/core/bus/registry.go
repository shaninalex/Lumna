package bus

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"slices"
)

type registry struct {
	handlers map[reflect.Type]Invoke
	chain    []Middleware
	sealed   bool
}

func newRegistry(mw []Middleware) registry {
	return registry{handlers: map[reflect.Type]Invoke{}, chain: mw}
}

func (r *registry) Seal() { r.sealed = true }

func register[M any, R any](r *registry, h func(context.Context, M) (R, error)) error {
	if r.sealed {
		return errors.New("bus: registration after Seal()")
	}
	key := reflect.TypeFor[M]()
	if _, dup := r.handlers[key]; dup {
		return fmt.Errorf("bus: handler for %s already registered", key)
	}
	r.handlers[key] = wrap(r.chain, func(ctx context.Context, msg any) (any, error) {
		return h(ctx, msg.(M))
	})
	return nil
}

func dispatch[M any, R any](ctx context.Context, r *registry, msg M) (R, error) {
	var zero R
	call, ok := r.handlers[reflect.TypeFor[M]()]
	if !ok {
		return zero, fmt.Errorf("bus: no handler for %T", msg)
	}
	res, err := call(ctx, msg)
	if err != nil {
		return zero, err
	}
	typed, ok := res.(R)
	if !ok {
		return zero, fmt.Errorf("bus: %T returned %T, want %T", msg, res, zero)
	}
	return typed, nil
}

func wrap(chain []Middleware, call Invoke) Invoke {
	for _, mw := range slices.Backward(chain) {
		call = mw(call)
	}
	return call
}
