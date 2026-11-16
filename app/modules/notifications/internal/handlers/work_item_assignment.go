package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/notifications/internal/domain"
	trackerContract "gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

const refWorkItem = "work_item"

// WorkItemAssigned notifies the assignee, unless they assigned themselves.
type WorkItemAssigned struct {
	c       clock.Clock
	deliver *Deliverer
}

func NewWorkItemAssigned(c clock.Clock, deliver *Deliverer) *WorkItemAssigned {
	return &WorkItemAssigned{c: c, deliver: deliver}
}

func (s *WorkItemAssigned) Handle(ctx context.Context, e trackerContract.WorkItemAssigned) error {
	return deliverAssignment(ctx, s.deliver, s.c, e.EventName(), e.WorkItemId, e.AssigneeId, e.ActorId)
}

// WorkItemUnassigned notifies the removed assignee, unless they removed
// themselves.
type WorkItemUnassigned struct {
	c       clock.Clock
	deliver *Deliverer
}

func NewWorkItemUnassigned(c clock.Clock, deliver *Deliverer) *WorkItemUnassigned {
	return &WorkItemUnassigned{c: c, deliver: deliver}
}

func (s *WorkItemUnassigned) Handle(ctx context.Context, e trackerContract.WorkItemUnassigned) error {
	return deliverAssignment(ctx, s.deliver, s.c, e.EventName(), e.WorkItemId, e.AssigneeId, e.ActorId)
}

func deliverAssignment(ctx context.Context, d *Deliverer, c clock.Clock, typ string, workItemID, assigneeID, actorID int) error {
	data, err := content(map[string]any{"actor_id": actorID})
	if err != nil {
		return err
	}
	n := domain.Notification{
		NotificationType: typ,
		Content:          data,
		RefId:            workItemID,
		RefType:          refWorkItem,
		Created:          c.Now(),
	}
	n.AddRecipients(actorID, assigneeID)
	return d.Deliver(ctx, n)
}
