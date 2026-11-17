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
	ifaces []ifaceSub
	chain  []Middleware
	sealed bool
}

// ifaceSub is a subscription to every event whose type implements iface.
type ifaceSub struct {
	iface reflect.Type
	call  Invoke
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
	call := wrap(b.chain, func(ctx context.Context, msg any) (any, error) {
		return nil, h(ctx, msg.(E))
	})
	b.subs[key] = append(b.subs[key], call)
	return nil
}

func SubscribeTo[I any](b *EventBus, h func(context.Context, I) error) error {
	if b.sealed {
		return errors.New("bus: subscription after Seal()")
	}
	iface := reflect.TypeFor[I]()
	if iface.Kind() != reflect.Interface {
		return fmt.Errorf("bus: SubscribeTo requires an interface type, got %s", iface)
	}
	call := wrap(b.chain, func(ctx context.Context, msg any) (any, error) {
		return nil, h(ctx, msg.(I))
	})
	b.ifaces = append(b.ifaces, ifaceSub{iface: iface, call: call})
	return nil
}

func (b *EventBus) Publish(ctx context.Context, e Event) error {
	var errs []error
	t := reflect.TypeOf(e)
	for _, h := range b.subs[t] {
		if err := safeCall(ctx, h, e); err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", e.EventName(), err))
		}
	}
	for _, s := range b.ifaces {
		if !t.Implements(s.iface) {
			continue
		}
		if err := safeCall(ctx, s.call, e); err != nil {
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
