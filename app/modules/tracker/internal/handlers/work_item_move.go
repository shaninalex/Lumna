package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

type WorkItemMove struct {
	repo  domain.WorkingItemRepo
	clock clock.Clock
}

func NewWorkItemMove(repo domain.WorkingItemRepo, clock clock.Clock) *WorkItemMove {
	return &WorkItemMove{
		repo:  repo,
		clock: clock,
	}
}

func (s *WorkItemMove) Handle(ctx context.Context, cmd contract.WorkItemMove) (contract.WorkItemView, error) {
	r, err := s.repo.Get(ctx, cmd.WorkItemId)
	if err != nil {
		return contract.WorkItemView{}, err
	}

	r.Rank = cmd.Rank
	updTime := s.clock.Now()
	r.UpdatedAt = &updTime

	if err := s.repo.Save(ctx, r); err != nil {
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
	}, nil
}
