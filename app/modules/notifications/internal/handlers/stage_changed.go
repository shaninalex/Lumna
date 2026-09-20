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
}

func NewWorkItemStageChanged(repo domain.NotificationRepo, c clock.Clock, eventBus *bus.EventBus) *WorkItemStageChanged {
	return &WorkItemStageChanged{
		c:        c,
		repo:     repo,
		eventBus: eventBus,
	}
}

func (s *WorkItemStageChanged) Handle(ctx context.Context, changed contract.WorkItemStageChanged) error {
	n := &domain.Notification{
		NotificationType: changed.EventName(),
		Content:          "Work item stage changed", // todo: use i18n _(""), use standard message library
		RefId:            &changed.WorkItemId,
		Created:          s.c.Now(),
	}
	if a, ok := actor.From(ctx); ok {
		n.IdentityId = &a.IdentityID
	}

	if err := s.repo.Save(ctx, n); err != nil {
		return err
	}

	notify := notificationsContract.NotificationCreated{Source: changed.EventName()}
	if err := s.eventBus.Publish(ctx, notify); err != nil {
		return err
	}

	return nil
}
