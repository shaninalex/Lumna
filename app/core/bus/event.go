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
	subs map[reflect.Type][]Invoke
}

func NewEventBus() *EventBus {
	return &EventBus{
		subs: make(map[reflect.Type][]Invoke),
	}
}

func Subscribe[E Event](b *EventBus, h func(context.Context, E) error) error {
	panic("Subscribe: not implemented")
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

func safeCall(ctx context.Context, h Invoke, e Event) error {
	panic("safeCall: not implemented")
}
