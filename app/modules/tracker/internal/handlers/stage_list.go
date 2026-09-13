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
	result, err := s.repo.Get(ctx, cmd.ScopeId)
	if err != nil {
		return nil, err
	}
	stages := make([]contract.StageView, len(result))
	for i, scope := range result {
		stages[i] = contract.StageView{
			Id:          scope.ID,
			ScopeId:     scope.ScopeID,
			Name:        scope.Name,
			Description: scope.Description,
			Category:    string(scope.Category),
			Position:    scope.Position,
			WIPLimit:    scope.WIPLimit,
			CreatedAt:   scope.CreatedAt,
			UpdatedAt:   scope.UpdatedAt,
		}
	}

	return stages, nil
}
