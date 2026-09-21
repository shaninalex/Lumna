package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

type WorkItemCreate struct {
	repo  domain.WorkingItemRepo
	clock clock.Clock
}

func NewWorkItemCreate(repo domain.WorkingItemRepo, clock clock.Clock) *WorkItemCreate {
	return &WorkItemCreate{
		repo:  repo,
		clock: clock,
	}
}

func (s *WorkItemCreate) Handle(ctx context.Context, cmd contract.WorkItemCreate) (contract.WorkItemView, error) {
	w, err := s.repo.Create(ctx, domain.WorkItem{
		Title:       cmd.Title,
		Description: cmd.Description,
		ProjectID:   cmd.ProjectId,
		Rank:        cmd.Position,
		StageID:     cmd.StageId,
		ScopeID:     cmd.ScopeId,
		DueTo:       cmd.DueTo,
	})
	if err != nil {
		return contract.WorkItemView{}, err
	}

	return toWorkItemView(w), nil
}
