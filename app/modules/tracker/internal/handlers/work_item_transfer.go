package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/core/bus"
	"gitlab.com/shaninalex/lumna/app/core/errs"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

type WorkItemTransfer struct {
	items  domain.WorkingItemRepo
	stages domain.StageRepo
	clock  clock.Clock
	bus    *bus.EventBus
}

func NewWorkItemTransfer(items domain.WorkingItemRepo, stages domain.StageRepo, clock clock.Clock, bus *bus.EventBus) *WorkItemTransfer {
	return &WorkItemTransfer{
		items:  items,
		stages: stages,
		clock:  clock,
		bus:    bus,
	}
}

func (s *WorkItemTransfer) Handle(ctx context.Context, cmd contract.WorkItemTransfer) (contract.WorkItemView, error) {
	stage, err := s.stages.GetById(ctx, cmd.StageId)
	if err != nil {
		return contract.WorkItemView{}, err
	}

	if stage.ScopeID != cmd.ScopeId {
		return contract.WorkItemView{}, errs.Validation("stage_scope_mismatch", "stage does not belong to this scope")
	}

	r, err := s.items.Get(ctx, cmd.WorkItemId)
	if err != nil {
		return contract.WorkItemView{}, err
	}

	r.ChangeStage(cmd.ScopeId, cmd.StageId, cmd.Rank, s.clock.Now())

	if err := s.items.Save(ctx, r); err != nil {
		return contract.WorkItemView{}, err
	}

	err = s.bus.Publish(ctx, contract.WorkItemStageChanged{WorkItemId: cmd.WorkItemId, StageId: cmd.StageId})
	if err != nil {
		return contract.WorkItemView{}, err
	}

	return toWorkItemView(r), nil
}
