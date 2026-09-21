package storage

import (
	"context"
	"database/sql"
	"errors"

	"gitlab.com/shaninalex/lumna/app/core/errs"
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

func workItemToRecord(wi domain.WorkItem) workItemRecord {
	return workItemRecord{
		ID:          wi.ID,
		ProjectID:   wi.ProjectID,
		Type:        string(wi.Type),
		ParentID:    sql.NullInt64{Int64: int64(wi.ParentID), Valid: wi.ParentID != 0},
		Title:       wi.Title,
		Description: sql.NullString{String: wi.Description, Valid: wi.Description != ""},
		ScopeID:     sql.NullInt64{Int64: int64(wi.ScopeID), Valid: wi.ScopeID != 0},
		StageID:     sql.NullInt64{Int64: int64(wi.StageID), Valid: wi.StageID != 0},
		Rank:        wi.Rank,
		CreatedAt:   wi.CreatedAt,
		UpdatedAt:   sql.NullTime{Time: wi.UpdatedAt, Valid: !wi.UpdatedAt.IsZero()},
		DueTo:       sql.NullTime{Time: wi.DueTo, Valid: !wi.DueTo.IsZero()},
	}
}

func (s *WorkingItemRepo) Create(ctx context.Context, wi domain.WorkItem) (domain.WorkItem, error) {
	record := workItemToRecord(wi)
	record.ID = 0
	record.CreatedAt = s.clock.Now()
	if err := gorm.G[workItemRecord](s.db.From(ctx)).Create(ctx, &record); err != nil {
		return domain.WorkItem{}, err
	}
	created := workItemToDomain(record)
	created.AssigneeIDs = wi.AssigneeIDs
	return created, nil
}

func (s *WorkingItemRepo) Update(ctx context.Context, wi domain.WorkItem) error {
	if wi.ID == 0 {
		return errs.Validation("work_item_id_required", "work item id is required to update a work item")
	}
	record := workItemToRecord(wi)
	res := s.db.From(ctx).Save(&record)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errs.NotFound("work_item_not_found", "work item not found")
	}
	return nil
}

func (s *WorkingItemRepo) Get(ctx context.Context, itemId int) (domain.WorkItem, error) {
	record, err := gorm.G[workItemRecord](s.db.From(ctx)).
		Preload("Assignees", nil).
		Where("id = ?", itemId).
		First(ctx)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.WorkItem{}, errs.NotFound("work_item_not_found", "work item not found")
	}
	if err != nil {
		return domain.WorkItem{}, err
	}
	return workItemToDomain(record), nil
}

func (s *WorkingItemRepo) ListByScope(ctx context.Context, scopeId int) ([]domain.WorkItem, error) {
	records, err := gorm.G[workItemRecord](s.db.From(ctx)).
		Preload("Assignees", nil).
		Where("scope_id = ?", scopeId).
		Find(ctx)
	if err != nil {
		return nil, err
	}
	return workItemsToDomain(records), nil
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
		Where("work_item_id = ? and identity_id = ?", itemId, identity).
		Delete(ctx)
	if err != nil {
		return err
	}
	if r == 0 {
		return errs.NotFound("work_item_assignment_not_found", "work item assignment not found")
	}
	return nil
}

func (s *WorkingItemRepo) Delete(ctx context.Context, itemId int) error {
	r, err := gorm.G[workItemRecord](s.db.From(ctx)).
		Where("id = ?", itemId).
		Delete(ctx)
	if err != nil {
		return err
	}
	if r == 0 {
		return errs.NotFound("work_item_not_found", "work item not found")
	}
	return nil
}

func (s *WorkingItemRepo) BatchDelete(ctx context.Context, itemIds []int) error {
	if len(itemIds) == 0 {
		return nil
	}
	_, err := gorm.G[workItemRecord](s.db.From(ctx)).
		Where("id IN ?", itemIds).
		Delete(ctx)
	return err
}
