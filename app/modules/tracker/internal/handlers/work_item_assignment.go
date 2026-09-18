package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

type WorkItemAssignment struct {
	items domain.WorkingItemRepo
	clock clock.Clock
}

func NewWorkItemAssignment(items domain.WorkingItemRepo, clock clock.Clock) *WorkItemAssignment {
	return &WorkItemAssignment{
		items: items,
		clock: clock,
	}
}

func (s *WorkItemAssignment) Handle(ctx context.Context, cmd contract.WorkItemAssign) (bool, error) {
	if err := s.items.Assignment(ctx, cmd.IdentityId, cmd.WorkItemId); err != nil {
		return false, err
	}
	return true, nil
}
