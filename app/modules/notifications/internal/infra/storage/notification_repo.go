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

// Save stores the notification together with its recipients. It runs in the
// caller's transaction when there is one.
func (r *NotificationRepo) Save(ctx context.Context, n domain.Notification) (domain.Notification, error) {
	record := notificationRecord{
		Type:      n.NotificationType,
		Content:   n.Content,
		RefID:     sql.NullInt64{Int64: int64(n.RefId), Valid: n.RefId != 0},
		RefType:   sql.NullString{String: n.RefType, Valid: n.RefType != ""},
		Priority:  sql.NullString{String: n.Priority, Valid: n.Priority != ""},
		CreatedAt: n.Created,
	}
	tx := r.db.From(ctx)
	if err := tx.Create(&record).Error; err != nil {
		return domain.Notification{}, err
	}

	recipients := make([]notificationRecipientRecord, len(n.Recipients))
	for i, rc := range n.Recipients {
		recipients[i] = notificationRecipientRecord{IdentityID: rc.IdentityID, NotificationID: record.ID}
	}
	if len(recipients) > 0 {
		if err := tx.Create(&recipients).Error; err != nil {
			return domain.Notification{}, err
		}
	}

	return toNotificationDomain(record, recipients), nil
}

func toNotificationDomain(n notificationRecord, recipients []notificationRecipientRecord) domain.Notification {
	dn := domain.Notification{
		ID:               n.ID,
		NotificationType: n.Type,
		Content:          n.Content,
		RefId:            int(n.RefID.Int64),
		RefType:          n.RefType.String,
		Priority:         n.Priority.String,
		Created:          n.CreatedAt,
		Recipients:       make([]domain.Recipient, len(recipients)),
	}
	for i, r := range recipients {
		dn.Recipients[i] = domain.Recipient{
			ID:             r.ID,
			NotificationID: r.NotificationID,
			IdentityID:     r.IdentityID,
		}
		if r.ReadAt.Valid {
			dn.Recipients[i].ReadAt = r.ReadAt.Time
		}
	}
	return dn
}
