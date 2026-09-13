package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/internal/domain"
)

type CreateScope struct {
	scopeRepo domain.ScopeRepo
}

func NewCreateScope(scopeRepo domain.ScopeRepo) *CreateScope {
	return &CreateScope{
		scopeRepo: scopeRepo,
	}
}

func (s *CreateScope) Handle(ctx context.Context, cmd contract.CreateScope) (contract.ScopeView, error) {
	var result contract.ScopeView
	scope := domain.Scope{
		Name:        cmd.Name,
		Description: cmd.Description,
		ProjectID:   cmd.ProjectId,
		UpdatedAt:   nil,
	}
	if err := s.scopeRepo.Save(ctx, &scope); err != nil {
		return result, err
	}

	result = contract.ScopeView{
		Id:          scope.ID,
		Name:        scope.Name,
		Description: scope.Description,
	}

	return result, nil
}
