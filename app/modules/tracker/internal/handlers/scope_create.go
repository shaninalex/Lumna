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
	scope, err := s.scopeRepo.Create(ctx, domain.Scope{
		Name:        cmd.Name,
		Description: cmd.Description,
		ProjectID:   cmd.ProjectId,
		CreatedAt:   s.clock.Now(),
	})
	if err != nil {
		return contract.ScopeView{}, err
	}

	return toScopeView(scope), nil
}
