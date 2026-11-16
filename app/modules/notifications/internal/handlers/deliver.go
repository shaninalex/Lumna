package handlers

import (
	"context"
	"encoding/json"

	"gitlab.com/shaninalex/lumna/app/core/bus"
	"gitlab.com/shaninalex/lumna/app/modules/notifications/contract"
	"gitlab.com/shaninalex/lumna/app/modules/notifications/internal/domain"
)

// Deliverer stores a notification with its recipients, announces it on the
// bus and pushes it to the recipients after the transaction commits.
// Event handlers only decide what the notification is and who gets it.
type Deliverer struct {
	repo     domain.NotificationRepo
	eventBus *bus.EventBus
	pusher   contract.Pusher
}

func NewDeliverer(repo domain.NotificationRepo, eventBus *bus.EventBus, pusher contract.Pusher) *Deliverer {
	return &Deliverer{repo: repo, eventBus: eventBus, pusher: pusher}
}

func (d *Deliverer) Deliver(ctx context.Context, n domain.Notification) error {
	if len(n.Recipients) == 0 {
		return nil
	}

	n, err := d.repo.Save(ctx, n)
	if err != nil {
		return err
	}

	recipients := n.RecipientIDs()
	if err := d.eventBus.Publish(ctx, contract.NotificationCreated{
		NotificationId: n.ID,
		Type:           n.NotificationType,
		RecipientIds:   recipients,
	}); err != nil {
		return err
	}

	d.pusher.Push(ctx, recipients, contract.NotificationView{
		ID:        n.ID,
		Type:      n.NotificationType,
		Content:   n.Content,
		RefID:     n.RefId,
		RefType:   n.RefType,
		CreatedAt: n.Created,
	})
	return nil
}

// content encodes the data the client renders the notification text from.
func content(v any) (string, error) {
	b, err := json.Marshal(v)
	return string(b), err
}
