package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/notifications/internal/domain"
	trackerContract "gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

type WorkItemMoved struct {
	clk     clock.Clock
	deliver *Deliverer
}

func NewWorkItemMoved(clk clock.Clock, deliver *Deliverer) *WorkItemMoved {
	return &WorkItemMoved{clk: clk, deliver: deliver}
}

func (s *WorkItemMoved) Handle(ctx context.Context, e trackerContract.WorkItemMoved) error {
	data, _ := content(map[string]any{"actor_id": e.ActorId})
	n := domain.Notification{
		NotificationType: e.EventName(),
		Content:          data,
		RefId:            e.WorkItemId,
		RefType:          refWorkItem,
		Created:          s.clk.Now(),
	}
	n.AddRecipients(e.ActorId, e.Assignees...)
	return s.deliver.Deliver(ctx, n)
}
