package email

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/mail"
	"time"

	"github.com/pkg/errors"
	"gitlab.com/shaninalex/lumna/app/platform/database"
	"gorm.io/gorm"
)

type Repository interface {
	PendingEmails(ctx context.Context) ([]Entry, error)
	Create(ctx context.Context, entry Entry) error
	Update(ctx context.Context, entry Entry) error
}

type Writer interface {
	Send(ctx context.Context, e string) error
}

func NewRepository(db *database.DB) Repository {
	return &emailRepository{
		db: db,
	}
}

type recordEmail struct {
	ID        int `gorm:"primaryKey;autoIncrement"`
	Subject   string
	Content   string
	Type      string
	Status    string `gorm:"default:pending"`
	Attempts  sql.NullInt16
	Receivers string
	Message   sql.NullString
	SendAt    sql.NullTime `gorm:"column:send_at;default:null"`
	CreatedAt time.Time    `gorm:"autoCreateTime"`
}

func (recordEmail) TableName() string { return "email_queue" }

func toDomainEntry(r recordEmail) Entry {
	e := Entry{
		ID:        r.ID,
		Subject:   r.Subject,
		Content:   r.Content,
		Type:      r.Type,
		Status:    r.Status,
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

func toRecordEntry(d Entry) recordEmail {
	r := recordEmail{
		ID:        d.ID,
		Subject:   d.Subject,
		Content:   d.Content,
		Type:      d.Type,
		Status:    d.Status,
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

type emailRepository struct {
	db *database.DB
}

var _ Repository = (*emailRepository)(nil)

func (s *emailRepository) PendingEmails(ctx context.Context) ([]Entry, error) {
	records, err := gorm.G[recordEmail](s.db.From(ctx)).Where("status = ?", "pending").Find(ctx)
	if err != nil {
		return nil, err
	}
	entries := make([]Entry, len(records))
	for i, e := range records {
		entries[i] = toDomainEntry(e)
	}
	return entries, nil
}

func (s *emailRepository) Create(ctx context.Context, entry Entry) error {
	record := toRecordEntry(entry)
	record.ID = 0
	record.CreatedAt = time.Now()
	return gorm.G[recordEmail](s.db.From(ctx)).Create(ctx, &record)
}

func (s *emailRepository) Update(ctx context.Context, entry Entry) error {
	record := toRecordEntry(entry)
	rows, err := gorm.G[recordEmail](s.db.From(ctx)).Updates(ctx, record)
	if rows <= 0 {
		return EntryWasNotFoundError
	}
	if err != nil {
		return errors.Wrap(EntryWasNotUpdatedError, err.Error())
	}
	return nil
}
