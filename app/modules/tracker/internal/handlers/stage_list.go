package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/internal/domain"
)

type StageList struct {
	repo domain.StageRepo
}

func NewStageList(repo domain.StageRepo) *StageList {
	return &StageList{
		repo: repo,
	}
}

func (s *StageList) Handle(ctx context.Context, cmd contract.StageList) ([]contract.StageView, error) {
	result, err := s.repo.ListByScope(ctx, cmd.ScopeId)
	if err != nil {
		return nil, err
	}
	return toStageViews(result), nil
}
