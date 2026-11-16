package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/core/bus"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

type WorkItemAssignment struct {
	items domain.WorkingItemRepo
	clock clock.Clock
	bus   *bus.EventBus
}

func NewWorkItemAssignment(items domain.WorkingItemRepo, clock clock.Clock, bus *bus.EventBus) *WorkItemAssignment {
	return &WorkItemAssignment{
		items: items,
		clock: clock,
		bus:   bus,
	}
}

func (s *WorkItemAssignment) Handle(ctx context.Context, cmd contract.WorkItemAssign) (bool, error) {
	assigned, err := s.items.Assignment(ctx, cmd.IdentityId, cmd.WorkItemId)
	if err != nil {
		return false, err
	}

	var e bus.Event = contract.WorkItemUnassigned{WorkItemId: cmd.WorkItemId, AssigneeId: cmd.IdentityId, ActorId: actorID(ctx)}
	if assigned {
		e = contract.WorkItemAssigned{WorkItemId: cmd.WorkItemId, AssigneeId: cmd.IdentityId, ActorId: actorID(ctx)}
	}
	if err := s.bus.Publish(ctx, e); err != nil {
		return false, err
	}

	return assigned, nil
}
