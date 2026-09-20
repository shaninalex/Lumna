package storage

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/notifications/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/database"
)

// NotificationRepo implements [domain.NotificationRepo].
type NotificationRepo struct {
	db *database.DB
}

func NewNotificationRepo(db *database.DB) *NotificationRepo { return &NotificationRepo{db: db} }

var _ domain.NotificationRepo = (*NotificationRepo)(nil)

func (r *NotificationRepo) Save(ctx context.Context, n *domain.Notification) error {
	record := notificationRecord{
		IdentityID: n.IdentityId,
		Type:       n.NotificationType,
		Content:    n.Content,
		RefID:      n.RefId,
		Priority:   n.Priority,
		CreatedAt:  n.Created,
	}
	if err := r.db.From(ctx).Save(&record).Error; err != nil {
		return err
	}

	n.ID = record.ID
	return nil
}
