package handlers

import (
	"context"
	"fmt"

	"gitlab.com/shaninalex/lumna/app/core/actor"
	"gitlab.com/shaninalex/lumna/app/modules/notifications/internal/domain"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

type WorkItemStageChanged struct {
	c    clock.Clock
	repo domain.NotificationRepo
}

func NewWorkItemStageChanged(repo domain.NotificationRepo, c clock.Clock) *WorkItemStageChanged {
	return &WorkItemStageChanged{
		c:    c,
		repo: repo,
	}
}

func (s *WorkItemStageChanged) Handle(ctx context.Context, changed contract.WorkItemStageChanged) error {
	fmt.Println("work item stage changed", changed.WorkItemId, changed.StageId, changed.EventName())
	n := &domain.Notification{
		NotificationType: changed.EventName(),
		Content:          "Work item stage changed",
		RefId:            &changed.WorkItemId,
		Created:          s.c.Now(),
	}
	if a, ok := actor.From(ctx); ok {
		n.IdentityId = &a.IdentityID
	}

	// todo: call bus.Publish if needed
	return s.repo.Save(ctx, n)
}
