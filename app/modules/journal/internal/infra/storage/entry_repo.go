package storage

import (
	"context"
	"database/sql"

	"gitlab.com/shaninalex/lumna/app/modules/journal/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/database"
)

// EntryRepo implements [domain.EntryRepo].
type EntryRepo struct {
	db *database.DB
}

func NewEntryRepo(db *database.DB) *EntryRepo { return &EntryRepo{db: db} }

var _ domain.EntryRepo = (*EntryRepo)(nil)

func (r *EntryRepo) Save(ctx context.Context, e domain.Entry) (domain.Entry, error) {
	record := entryRecord{
		ActorID:    sql.NullInt64{Int64: int64(e.ActorID), Valid: e.ActorID != 0},
		EntityType: e.EntityType,
		EntityID:   e.EntityID,
		EventType:  e.EventType,
		Data:       sql.NullString{String: e.Data, Valid: e.Data != ""},
		CreatedAt:  e.CreatedAt,
	}
	if err := r.db.From(ctx).Create(&record).Error; err != nil {
		return domain.Entry{}, err
	}
	e.ID = record.ID
	return e, nil
}
