package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

type WorkItemUpdate struct {
	items domain.WorkingItemRepo
	clock clock.Clock
}

func NewWorkItemUpdate(items domain.WorkingItemRepo, clock clock.Clock) *WorkItemUpdate {
	return &WorkItemUpdate{
		items: items,
		clock: clock,
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

	return toWorkItemView(r), nil
}
