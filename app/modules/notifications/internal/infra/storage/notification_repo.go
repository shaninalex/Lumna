package storage

import (
	"context"
	"database/sql"

	"gitlab.com/shaninalex/lumna/app/modules/notifications/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/database"
)

// NotificationRepo implements [domain.NotificationRepo].
type NotificationRepo struct {
	db *database.DB
}

func NewNotificationRepo(db *database.DB) *NotificationRepo { return &NotificationRepo{db: db} }

var _ domain.NotificationRepo = (*NotificationRepo)(nil)

func (r *NotificationRepo) Save(ctx context.Context, n domain.Notification) (domain.Notification, error) {
	record := notificationRecord{
		IdentityID: sql.NullInt64{Int64: int64(n.IdentityId), Valid: n.IdentityId != 0},
		Type:       n.NotificationType,
		Content:    n.Content,
		RefID:      sql.NullInt64{Int64: int64(n.RefId), Valid: n.RefId != 0},
		Priority:   sql.NullString{String: n.Priority, Valid: n.Priority != ""},
		CreatedAt:  n.Created,
	}
	if err := r.db.From(ctx).Save(&record).Error; err != nil {
		return domain.Notification{}, err
	}
	return toNotificationDomain(record), nil
}

func toNotificationDomain(n notificationRecord) domain.Notification {
	dn := domain.Notification{
		ID:               n.ID,
		IdentityId:       int(n.IdentityID.Int64),
		NotificationType: n.Type,
		Content:          n.Content,
		RefId:            int(n.RefID.Int64),
		Priority:         n.Priority.String,
		Created:          n.CreatedAt,
	}
	if n.ReadAt.Valid {
		dn.ReadAt = n.ReadAt.Time
	}
	return dn
}
