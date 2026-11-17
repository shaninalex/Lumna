package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/notifications/internal/domain"
	trackerContract "gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

// WorkItemStageChanged notifies the work item's assignees, except the one
// who moved it.
type WorkItemStageChanged struct {
	c       clock.Clock
	deliver *Deliverer
}

func NewWorkItemStageChanged(c clock.Clock, deliver *Deliverer) *WorkItemStageChanged {
	return &WorkItemStageChanged{c: c, deliver: deliver}
}

func (s *WorkItemStageChanged) Handle(ctx context.Context, e trackerContract.WorkItemStageChanged) error {
	data, err := content(map[string]any{
		"actor_id": e.ActorId,
		"stage_id": e.StageId,
	})

	if err != nil {
		return err
	}
	n := domain.Notification{
		NotificationType: e.EventName(),
		Content:          data,
		RefId:            e.WorkItemId,
		RefType:          refWorkItem,
		Created:          s.c.Now(),
	}
	n.AddRecipients(e.ActorId, e.AssigneeIds...)
	return s.deliver.Deliver(ctx, n)
}
