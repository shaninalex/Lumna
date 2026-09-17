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
	w := domain.WorkItem{
		Title:       cmd.Title,
		Description: cmd.Description,
		ProjectID:   cmd.ProjectId,
		Rank:        cmd.Position,
		StageID:     cmd.StageId,
		ScopeID:     cmd.ScopeId,
		DueTo:       cmd.DueTo,
	}
	if err := s.repo.Save(ctx, &w); err != nil {
		return contract.WorkItemView{}, err
	}

	return contract.WorkItemView{
		Id:          w.ID,
		Title:       w.Title,
		Description: w.Description,
		ProjectId:   w.ProjectID,
		StageId:     w.StageID,
		ScopeId:     w.ScopeID,
		Rank:        w.Rank,
		DueTo:       w.DueTo,
		CreatedAt:   w.CreatedAt,
		UpdatedAt:   w.UpdatedAt,
	}, nil
}
