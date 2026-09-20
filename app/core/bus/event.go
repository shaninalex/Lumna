package bus

import (
	"context"
	"errors"
	"fmt"
	"reflect"
)

type Event interface {
	EventName() string
}

type EventBus struct {
	subs   map[reflect.Type][]Invoke
	chain  []Middleware
	sealed bool
}

func NewEventBus(mw ...Middleware) *EventBus {
	return &EventBus{
		subs:  make(map[reflect.Type][]Invoke),
		chain: mw,
	}
}

func (b *EventBus) Seal() { b.sealed = true }

func Subscribe[E Event](b *EventBus, h func(context.Context, E) error) error {
	if b.sealed {
		return errors.New("bus: subscription after Seal()")
	}
	key := reflect.TypeFor[E]()
	call := Invoke(func(ctx context.Context, msg any) (any, error) {
		return nil, h(ctx, msg.(E))
	})
	for i := len(b.chain) - 1; i >= 0; i-- {
		call = b.chain[i](call)
	}
	b.subs[key] = append(b.subs[key], call)
	return nil
}

func (b *EventBus) Publish(ctx context.Context, e Event) error {
	var errs []error
	for _, h := range b.subs[reflect.TypeOf(e)] {
		if err := safeCall(ctx, h, e); err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", e.EventName(), err))
		}
	}
	return errors.Join(errs...)
}

func safeCall(ctx context.Context, h Invoke, e Event) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if pe, ok := r.(error); ok {
				err = fmt.Errorf("panic: %w", pe)
			} else {
				err = fmt.Errorf("panic: %v", r)
			}
		}
	}()

	_, err = h(ctx, e)
	return err
}
