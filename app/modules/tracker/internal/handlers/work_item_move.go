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
	w, err := s.repo.Get(ctx, cmd.WorkItemId)
	if err != nil {
		return contract.WorkItemView{}, err
	}

	w.ChangeRank(cmd.Rank, s.clock.Now())

	if err := s.repo.Update(ctx, w); err != nil {
		return contract.WorkItemView{}, err
	}
	return toWorkItemView(w), nil
}
