package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

type StateCreate struct {
	clock clock.Clock
	repo  domain.StageRepo
}

func NewCreateStage(repo domain.StageRepo, clock clock.Clock) *StateCreate {
	return &StateCreate{
		clock: clock,
		repo:  repo,
	}
}

func (s *StateCreate) Handle(ctx context.Context, cmd contract.StageCreate) (contract.StageView, error) {
	stage, err := s.repo.Create(ctx, domain.Stage{
		ScopeID:     cmd.ScopeID,
		Name:        cmd.Name,
		Description: cmd.Description,
		Category:    domain.StageCategory(cmd.Category),
		Position:    cmd.Position,
		WIPLimit:    cmd.WIPLimit,
		CreatedAt:   s.clock.Now(),
	})
	if err != nil {
		return contract.StageView{}, err
	}

	return toStageView(stage), nil
}
