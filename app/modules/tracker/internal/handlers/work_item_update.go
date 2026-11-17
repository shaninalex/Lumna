package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/core/bus"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

type WorkItemUpdate struct {
	items domain.WorkingItemRepo
	clock clock.Clock
	bus   *bus.EventBus
}

func NewWorkItemUpdate(items domain.WorkingItemRepo, clock clock.Clock, bus *bus.EventBus) *WorkItemUpdate {
	return &WorkItemUpdate{
		items: items,
		clock: clock,
		bus:   bus,
	}
}

func (s *WorkItemUpdate) Handle(ctx context.Context, cmd contract.WorkItemUpdate) (contract.WorkItemView, error) {
	r, err := s.items.Get(ctx, cmd.WorkItemId)
	if err != nil {
		return contract.WorkItemView{}, err
	}

	r.Title = cmd.Title
	r.Description = cmd.Description
	r.UpdatedAt = s.clock.Now()

	if err := s.items.Update(ctx, r); err != nil {
		return contract.WorkItemView{}, err
	}

	_ = s.bus.Publish(ctx, contract.WorkItemChanged{
		WorkItemId: cmd.WorkItemId,
		ActorId:    actorID(ctx),
	})

	return toWorkItemView(r), nil
}
