package storage

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/mail"
	"time"

	"gitlab.com/shaninalex/lumna/app/modules/notifications/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/database"
	"gorm.io/gorm"
)

type EmailRepository struct {
	db *database.DB
}

func NewEmailRepo(db *database.DB) domain.EmailRepo {
	return &EmailRepository{
		db: db,
	}
}

var _ domain.EmailRepo = (*EmailRepository)(nil)

func (s *EmailRepository) PendingEmails(ctx context.Context) ([]domain.Email, error) {
	records, err := gorm.G[recordEmail](s.db.From(ctx)).Where("status = ?", "pending").Find(ctx)
	if err != nil {
		return nil, err
	}
	entries := make([]domain.Email, len(records))
	for i, e := range records {
		entries[i] = toDomainEntry(e)
	}
	return entries, nil
}

func (s *EmailRepository) Create(ctx context.Context, entry domain.Email) error {
	record := toRecordEntry(entry)
	record.ID = 0
	record.CreatedAt = time.Now()
	return gorm.G[recordEmail](s.db.From(ctx)).Create(ctx, &record)
}

func (s *EmailRepository) Update(ctx context.Context, entry domain.Email) error {
	record := toRecordEntry(entry)
	rows, err := gorm.G[recordEmail](s.db.From(ctx)).Updates(ctx, record)
	if rows <= 0 {
		return domain.EntryWasNotFoundError
	}
	if err != nil {
		return errors.Join(domain.EntryWasNotUpdatedError, err)
	}
	return nil
}

func toDomainEntry(r recordEmail) domain.Email {
	e := domain.Email{
		ID:        r.ID,
		Subject:   r.Subject,
		Content:   r.Content,
		Type:      domain.EntryType(r.Type),
		Status:    domain.EmailStatus(r.Status),
		Attempts:  r.Attempts.Int16,
		Message:   r.Message.String,
		SendAt:    r.SendAt.Time,
		CreatedAt: r.CreatedAt,
	}
	var receivers []mail.Address
	if err := json.Unmarshal([]byte(r.Receivers), &receivers); err == nil {
		e.Receivers = receivers
	}
	return e
}

func toRecordEntry(d domain.Email) recordEmail {
	r := recordEmail{
		ID:        d.ID,
		Subject:   d.Subject,
		Content:   d.Content,
		Type:      string(d.Type),
		Status:    string(d.Status),
		Attempts:  sql.NullInt16{Int16: d.Attempts, Valid: d.Attempts > 0},
		Message:   sql.NullString{String: d.Message, Valid: d.Message != ""},
		SendAt:    sql.NullTime{Time: d.SendAt, Valid: !d.SendAt.IsZero()},
		CreatedAt: d.CreatedAt,
	}

	if b, err := json.Marshal(d.Receivers); err == nil {
		r.Receivers = string(b)
	}

	return r
}
