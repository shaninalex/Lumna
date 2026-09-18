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
	updTime := s.clock.Now()
	r.UpdatedAt = &updTime

	if err := s.items.Save(ctx, r); err != nil {
		return contract.WorkItemView{}, err
	}

	return contract.WorkItemView{
		Id:          r.ID,
		Title:       r.Title,
		Description: r.Description,
		ProjectId:   r.ProjectID,
		StageId:     r.StageID,
		ScopeId:     r.ScopeID,
		Rank:        r.Rank,
		DueTo:       r.DueTo,
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
		Assignees:   r.AssigneeIDs,
	}, nil
}
