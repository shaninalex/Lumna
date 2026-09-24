package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/core/actor"
	"gitlab.com/shaninalex/lumna/app/core/bus"
	notificationsContract "gitlab.com/shaninalex/lumna/app/modules/notifications/contract"
	"gitlab.com/shaninalex/lumna/app/modules/notifications/internal/domain"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

type WorkItemStageChanged struct {
	c        clock.Clock
	repo     domain.NotificationRepo
	eventBus *bus.EventBus
	pusher   notificationsContract.Pusher
}

func NewWorkItemStageChanged(
	repo domain.NotificationRepo,
	c clock.Clock,
	eventBus *bus.EventBus,
	pusher notificationsContract.Pusher,
) *WorkItemStageChanged {
	return &WorkItemStageChanged{
		c:        c,
		repo:     repo,
		eventBus: eventBus,
		pusher:   pusher,
	}
}

func (s *WorkItemStageChanged) Handle(ctx context.Context, changed contract.WorkItemStageChanged) error {
	n := domain.Notification{
		NotificationType: changed.EventName(),
		Content:          "Work item stage changed", // todo: use i18n _(""), use standard message library
		RefId:            changed.WorkItemId,
		Created:          s.c.Now(),
	}
	if a, ok := actor.From(ctx); ok {
		n.IdentityId = a.IdentityID
	}

	n, err := s.repo.Save(ctx, n)
	if err != nil {
		return err
	}

	// to internal subscribers
	notify := notificationsContract.NotificationCreated{Source: changed.EventName()}
	if err := s.eventBus.Publish(ctx, notify); err != nil {
		return err
	}

	// to adapters
	// NOTE: assignees id's and watchers
	s.pusher.Push(ctx, []int{n.IdentityId}, notificationsContract.NotificationView{
		ID:        n.ID,
		Type:      changed.EventName(),
		Content:   n.Content,
		RefID:     n.RefId,
		CreatedAt: n.Created,
	})

	return nil
}
