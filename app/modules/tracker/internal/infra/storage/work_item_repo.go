package storage

import (
	"context"
	"errors"

	"gitlab.com/shaninalex/lumna/app/modules/tracker/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/database"
	"gorm.io/gorm"
)

type WorkingItemRepo struct {
	db *database.DB
}

var _ domain.WorkingItemRepo = (*WorkingItemRepo)(nil)

func NewWorkingItemRepo(db *database.DB) *WorkingItemRepo {
	return &WorkingItemRepo{
		db: db,
	}
}

func (s *WorkingItemRepo) Save(ctx context.Context, wi *domain.WorkItem) error {
	record := workItemRecord{
		ID:          wi.ID,
		ProjectID:   wi.ProjectID,
		Type:        string(wi.Type),
		ParentID:    wi.ParentID,
		Title:       wi.Title,
		Description: wi.Description,
		ScopeID:     wi.ScopeID,
		StageID:     wi.StageID,
		Rank:        wi.Rank,
		CreatedAt:   wi.CreatedAt,
		UpdatedAt:   wi.UpdatedAt,
		DueTo:       wi.DueTo,
	}
	if err := s.db.From(ctx).Save(&record).Error; err != nil {
		return err
	}

	wi.ID = record.ID
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
		assigneeIDs := make([]int, len(record.Assignees))
		for j, assignee := range record.Assignees {
			assigneeIDs[j] = assignee.IdentityID
		}
		workItems[i] = domain.WorkItem{
			ID:          record.ID,
			ProjectID:   record.ProjectID,
			Type:        domain.WorkItemType(record.Type),
			ParentID:    record.ParentID,
			Title:       record.Title,
			Description: record.Description,
			ScopeID:     record.ScopeID,
			StageID:     record.StageID,
			Rank:        record.Rank,
			DueTo:       record.DueTo,
			CreatedAt:   record.CreatedAt,
			UpdatedAt:   record.UpdatedAt,
			AssigneeIDs: assigneeIDs,
		}
	}

	return workItems, nil
}

func (s *WorkingItemRepo) Get(ctx context.Context, itemId int) (*domain.WorkItem, error) {
	record, err := gorm.G[workItemRecord](s.db.From(ctx)).
		Where("id = ?", itemId).
		First(ctx)
	if err != nil {
		return nil, err
	}
	return &domain.WorkItem{
		ID:          record.ID,
		ProjectID:   record.ProjectID,
		Type:        domain.WorkItemType(record.Type),
		ParentID:    record.ParentID,
		Title:       record.Title,
		Description: record.Description,
		ScopeID:     record.ScopeID,
		StageID:     record.StageID,
		Rank:        record.Rank,
		DueTo:       record.DueTo,
		CreatedAt:   record.CreatedAt,
		UpdatedAt:   record.UpdatedAt,
	}, nil
}

func (s *WorkingItemRepo) Assignment(ctx context.Context, identity, itemId int) error {
	_, err := gorm.G[workItemAssignRecord](s.db.From(ctx)).
		Where("work_item_id = ? and identity_id", itemId, identity).
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
