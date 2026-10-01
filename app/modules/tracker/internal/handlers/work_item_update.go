package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/core/actor"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

type WorkItemUpdate struct {
	items    domain.WorkingItemRepo
	clock    clock.Clock
	activity domain.ActivityRepo
}

func NewWorkItemUpdate(items domain.WorkingItemRepo, activity domain.ActivityRepo, clock clock.Clock) *WorkItemUpdate {
	return &WorkItemUpdate{
		items:    items,
		clock:    clock,
		activity: activity,
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

	activity := domain.Activity{
		EntityID:   r.ID,
		EntityType: "work_item",
		EventType:  "work_item/update",
		Content:    "Item was updated",
	}
	if a, ok := actor.From(ctx); ok {
		activity.IdentityID = a.IdentityID
	}

	if _, err = s.activity.Create(ctx, activity); err != nil {
		return contract.WorkItemView{}, err
	}

	return toWorkItemView(r), nil
}
