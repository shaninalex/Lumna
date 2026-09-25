package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/internal/domain"
)

type ScopeList struct {
	repo      domain.ScopeRepo
	stageRepo domain.StageRepo
	itemsRepo domain.WorkingItemRepo
}

func NewScopeList(repo domain.ScopeRepo, stageRepo domain.StageRepo, itemsRepo domain.WorkingItemRepo) *ScopeList {
	return &ScopeList{
		repo:      repo,
		stageRepo: stageRepo,
		itemsRepo: itemsRepo,
	}
}

func (s *ScopeList) Handle(ctx context.Context, cmd contract.ScopeList) ([]contract.ScopeView, error) {
	result, err := s.repo.ListByProject(ctx, cmd.ProjectId)
	if err != nil {
		return nil, err
	}
	return toScopeViews(result), nil
}
