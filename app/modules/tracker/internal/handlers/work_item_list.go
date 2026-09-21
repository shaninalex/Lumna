package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

type WorkItemList struct {
	repo  domain.WorkingItemRepo
	clock clock.Clock
}

func NewWorkItemList(repo domain.WorkingItemRepo, clock clock.Clock) *WorkItemList {
	return &WorkItemList{
		repo:  repo,
		clock: clock,
	}
}

func (s *WorkItemList) Handle(ctx context.Context, cmd contract.WorkItemList) ([]contract.WorkItemView, error) {
	results, err := s.repo.ListByScope(ctx, cmd.ScopeId)
	if err != nil {
		return nil, err
	}

	return toWorkItemViews(results), nil
}
