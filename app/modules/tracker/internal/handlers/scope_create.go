package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

type ScopeCreate struct {
	clock     clock.Clock
	scopeRepo domain.ScopeRepo
}

func NewCreateScope(scopeRepo domain.ScopeRepo, clock clock.Clock) *ScopeCreate {
	return &ScopeCreate{
		clock:     clock,
		scopeRepo: scopeRepo,
	}
}

func (s *ScopeCreate) Handle(ctx context.Context, cmd contract.CreateScope) (contract.ScopeView, error) {
	var result contract.ScopeView
	scope := domain.Scope{
		Name:        cmd.Name,
		Description: cmd.Description,
		ProjectID:   cmd.ProjectId,
		CreatedAt:   s.clock.Now(),
		UpdatedAt:   nil,
	}
	if err := s.scopeRepo.Save(ctx, &scope); err != nil {
		return result, err
	}

	result = contract.ScopeView{
		Id:          scope.ID,
		Name:        scope.Name,
		Description: scope.Description,
		ProjectId:   scope.ProjectID,
		CreatedAt:   scope.CreatedAt,
		UpdatedAt:   scope.UpdatedAt,
	}

	return result, nil
}
