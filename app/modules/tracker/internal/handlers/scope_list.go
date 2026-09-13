package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/internal/domain"
)

type ScopeList struct {
	repo domain.ScopeRepo
}

func NewScopeList(repo domain.ScopeRepo) *ScopeList {
	return &ScopeList{
		repo: repo,
	}
}

func (s *ScopeList) Handle(ctx context.Context, cmd contract.ScopeList) ([]contract.ScopeView, error) {
	result, err := s.repo.Get(ctx, cmd.ProjectId)
	if err != nil {
		return nil, err
	}
	scopes := make([]contract.ScopeView, len(result))
	for i, scope := range result {
		scopes[i] = contract.ScopeView{
			Id:          scope.ID,
			Name:        scope.Name,
			Description: scope.Description,
			ProjectId:   scope.ProjectID,
			CreatedAt:   scope.CreatedAt,
			UpdatedAt:   scope.UpdatedAt,
		}
	}

	return scopes, nil
}
