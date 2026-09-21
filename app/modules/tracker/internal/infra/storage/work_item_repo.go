package storage

import (
	"context"
	"database/sql"
	"errors"

	"gitlab.com/shaninalex/lumna/app/modules/tracker/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
	"gitlab.com/shaninalex/lumna/app/platform/database"
	"gorm.io/gorm"
)

type WorkingItemRepo struct {
	db    *database.DB
	clock clock.Clock
}

var _ domain.WorkingItemRepo = (*WorkingItemRepo)(nil)

func NewWorkingItemRepo(db *database.DB, clock clock.Clock) *WorkingItemRepo {
	return &WorkingItemRepo{
		db:    db,
		clock: clock,
	}
}

func (s *WorkingItemRepo) Save(ctx context.Context, wi *domain.WorkItem) error {
	record := workItemRecord{
		ID:          wi.ID,
		ProjectID:   wi.ProjectID,
		Type:        string(wi.Type),
		ParentID:    sql.NullInt64{Int64: int64(wi.ParentID), Valid: wi.ParentID != 0},
		Title:       wi.Title,
		Description: sql.NullString{String: wi.Description, Valid: wi.Description != ""},
		ScopeID:     sql.NullInt64{Int64: int64(wi.ScopeID), Valid: wi.ScopeID != 0},
		StageID:     sql.NullInt64{Int64: int64(wi.StageID), Valid: wi.StageID != 0},
		Rank:        wi.Rank,
		CreatedAt:   s.clock.Now(),
		DueTo:       sql.NullTime{Time: wi.DueTo, Valid: wi.DueTo.IsZero()},
	}
	if err := s.db.From(ctx).Save(&record).Error; err != nil {
		return err
	}
	wi.ID = record.ID
	wi.CreatedAt = record.CreatedAt
	return nil
}

func (s *WorkingItemRepo) List(ctx context.Context, scopeId int) ([]domain.WorkItem, error) {
	records, err := gorm.G[workItemRecord](s.db.From(ctx)).
		Preload("Assignees", nil).
		Where("scope_id = ?", scopeId).
		Find(ctx)
	if err != nil {
		return nil, err
	}

	workItems := make([]domain.WorkItem, len(records))
	for i, record := range records {
		workItems[i] = workItemToDomain(record)
	}

	return workItems, nil
}

func (s *WorkingItemRepo) Get(ctx context.Context, itemId int) (*domain.WorkItem, error) {
	record, err := gorm.G[workItemRecord](s.db.From(ctx)).
		Preload("Assignees", nil).
		Where("id = ?", itemId).
		First(ctx)
	if err != nil {
		return nil, err
	}

	d := workItemToDomain(record)
	return &d, nil
}

func (s *WorkingItemRepo) Assignment(ctx context.Context, identity, itemId int) error {
	_, err := gorm.G[workItemAssignRecord](s.db.From(ctx)).
		Where("work_item_id = ? and identity_id = ?", itemId, identity).
		First(ctx)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		record := workItemAssignRecord{WorkItemID: itemId, IdentityID: identity}
		return gorm.G[workItemAssignRecord](s.db.From(ctx)).Create(ctx, &record)
	}
	if err != nil {
		return err
	}

	r, err := gorm.G[workItemAssignRecord](s.db.From(ctx)).
		Where("work_item_id = ? and identity_id", itemId, identity).
		Delete(ctx)
	if err != nil {
		return err
	}
	if r <= 0 {
		return errors.New("unable to delete: work_item_assignment not found in database")
	}
	return nil
}

func (s *WorkingItemRepo) Delete(ctx context.Context, itemId int) (bool, error) {
	r, err := gorm.G[workItemRecord](s.db.From(ctx)).
		Where("id = ?", itemId).
		Delete(ctx)
	if err != nil {
		return false, err
	}
	if r <= 0 {
		return false, errors.New("unable to delete work_item, something went wrong")
	}
	return true, nil
}

func (s *WorkingItemRepo) BatchDelete(ctx context.Context, itemIds []int) (bool, error) {
	_, err := gorm.G[workItemRecord](s.db.From(ctx)).
		Where("id IN ?", itemIds).
		Delete(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}
