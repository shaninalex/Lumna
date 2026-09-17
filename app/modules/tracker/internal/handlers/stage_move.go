package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/core/errs"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

type StageMove struct {
	repo  domain.StageRepo
	clock clock.Clock
}

func NewStageMove(repo domain.StageRepo, clock clock.Clock) *StageMove {
	return &StageMove{
		repo:  repo,
		clock: clock,
	}
}

func (s *StageMove) Handle(ctx context.Context, cmd contract.StageMove) (contract.StageView, error) {
	stage, err := s.repo.GetById(ctx, cmd.StageId)
	if err != nil {
		return contract.StageView{}, err
	}

	if stage.ScopeID != cmd.ScopeId {
		return contract.StageView{}, errs.Validation("stage_scope_mismatch", "stage does not belong to this scope")
	}

	stage.Position = cmd.Position
	stage.UpdatedAt = s.clock.Now()

	if err := s.repo.Save(ctx, stage); err != nil {
		return contract.StageView{}, err
	}

	return contract.StageView{
		Id:          stage.ID,
		ScopeId:     stage.ScopeID,
		Name:        stage.Name,
		Description: stage.Description,
		Category:    string(stage.Category),
		Position:    stage.Position,
		WIPLimit:    stage.WIPLimit,
		CreatedAt:   stage.CreatedAt,
		UpdatedAt:   stage.UpdatedAt,
	}, nil
}
