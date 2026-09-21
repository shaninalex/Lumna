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
	stage, err := s.repo.Get(ctx, cmd.StageId)
	if err != nil {
		return contract.StageView{}, err
	}

	if stage.ScopeID != cmd.ScopeId {
		return contract.StageView{}, errs.Validation("stage_scope_mismatch", "stage does not belong to this scope")
	}

	stage.UpdatePosition(cmd.Position, s.clock.Now())
	if err := s.repo.Update(ctx, stage); err != nil {
		return contract.StageView{}, err
	}

	return toStageView(stage), nil
}
