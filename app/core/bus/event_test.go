package bus_test

import (
	"context"
	"errors"
	"testing"

	"gitlab.com/shaninalex/lumna/app/core/bus"
)

type TestEvent struct {
	Message string
}

func (e TestEvent) EventName() string {
	return "test.event"
}

func TestEventBus_SubscribeAndPublish(t *testing.T) {
	b := bus.NewEventBus()
	ctx := context.Background()

	var received []string
	err := bus.Subscribe(b, func(ctx context.Context, e TestEvent) error {
		received = append(received, e.Message)
		return nil
	})
	if err != nil {
		t.Fatalf("unexpected subscribe error: %v", err)
	}

	err = bus.Subscribe(b, func(ctx context.Context, e TestEvent) error {
		received = append(received, e.Message+"_2")
		return nil
	})
	if err != nil {
		t.Fatalf("unexpected subscribe error: %v", err)
	}

	b.Seal()

	err = b.Publish(ctx, TestEvent{Message: "hello"})
	if err != nil {
		t.Fatalf("unexpected publish error: %v", err)
	}

	if len(received) != 2 || received[0] != "hello" || received[1] != "hello_2" {
		t.Fatalf("unexpected received events: %+v", received)
	}
}

func TestEventBus_PanicRecovery(t *testing.T) {
	b := bus.NewEventBus()
	ctx := context.Background()

	var executedAfterPanic bool
	_ = bus.Subscribe(b, func(ctx context.Context, e TestEvent) error {
		panic("something went wrong")
	})

	_ = bus.Subscribe(b, func(ctx context.Context, e TestEvent) error {
		executedAfterPanic = true
		return nil
	})

	err := b.Publish(ctx, TestEvent{Message: "boom"})
	if err == nil {
		t.Fatal("expected error from panic, got nil")
	}

	if !executedAfterPanic {
		t.Fatal("expected second subscriber to execute even if first panicked")
	}
}

func TestEventBus_SubscribeAfterSeal(t *testing.T) {
	b := bus.NewEventBus()
	b.Seal()

	err := bus.Subscribe(b, func(ctx context.Context, e TestEvent) error {
		return nil
	})
	if err == nil {
		t.Fatal("expected error when subscribing after Seal(), got nil")
	}
}

func TestEventBus_HandlerError(t *testing.T) {
	b := bus.NewEventBus()
	ctx := context.Background()

	expectedErr := errors.New("handler failed")
	_ = bus.Subscribe(b, func(ctx context.Context, e TestEvent) error {
		return expectedErr
	})

	err := b.Publish(ctx, TestEvent{Message: "fail"})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, expectedErr) {
		t.Fatalf("expected error to wrap handler error, got: %v", err)
	}
}

type namedEvent interface {
	bus.Event
	Name() string
}

type OtherEvent struct{}

func (OtherEvent) EventName() string { return "other.event" }

func (e TestEvent) Name() string { return e.Message }

func TestEventBus_SubscribeTo(t *testing.T) {
	b := bus.NewEventBus()
	ctx := context.Background()

	var order []string
	_ = bus.Subscribe(b, func(ctx context.Context, e TestEvent) error {
		order = append(order, "concrete")
		return nil
	})
	err := bus.SubscribeTo(b, func(ctx context.Context, e namedEvent) error {
		order = append(order, "iface:"+e.Name())
		return nil
	})
	if err != nil {
		t.Fatalf("unexpected subscribe error: %v", err)
	}
	b.Seal()

	if err := b.Publish(ctx, TestEvent{Message: "hi"}); err != nil {
		t.Fatalf("unexpected publish error: %v", err)
	}
	// OtherEvent does not implement namedEvent and must not reach the handler.
	if err := b.Publish(ctx, OtherEvent{}); err != nil {
		t.Fatalf("unexpected publish error: %v", err)
	}

	if len(order) != 2 || order[0] != "concrete" || order[1] != "iface:hi" {
		t.Fatalf("unexpected calls: %+v", order)
	}
}

func TestEventBus_SubscribeToRejectsConcreteType(t *testing.T) {
	b := bus.NewEventBus()
	err := bus.SubscribeTo(b, func(ctx context.Context, e TestEvent) error { return nil })
	if err == nil {
		t.Fatal("expected error for non-interface type, got nil")
	}
}
