package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/core/actor"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

type WorkItemAssignment struct {
	items    domain.WorkingItemRepo
	clock    clock.Clock
	activity domain.ActivityRepo
}

func NewWorkItemAssignment(items domain.WorkingItemRepo, activity domain.ActivityRepo, clock clock.Clock) *WorkItemAssignment {
	return &WorkItemAssignment{
		items:    items,
		clock:    clock,
		activity: activity,
	}
}

func (s *WorkItemAssignment) Handle(ctx context.Context, cmd contract.WorkItemAssign) (bool, error) {
	assigned, err := s.items.Assignment(ctx, cmd.IdentityId, cmd.WorkItemId)
	if err != nil {
		return false, err
	}

	activity := domain.Activity{
		EntityID:   cmd.WorkItemId,
		EntityType: "work_item",
		EventType:  "work_item/assignment",
	}
	if assigned {
		activity.Content = "Item was assigned user"
	} else {
		activity.Content = "User was unassigned from this work item"
	}
	if a, ok := actor.From(ctx); ok {
		activity.IdentityID = a.IdentityID
	}

	if _, err := s.activity.Create(ctx, activity); err != nil {
		return false, err
	}

	return true, nil
}
