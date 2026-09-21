package storage

import (
	"context"
	"database/sql"

	"gitlab.com/shaninalex/lumna/app/core/errs"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/database"
	"gorm.io/gorm"
)

type ScopeRepo struct {
	db *database.DB
}

var _ domain.ScopeRepo = (*ScopeRepo)(nil)

func NewScopeRepo(db *database.DB) *ScopeRepo {
	return &ScopeRepo{
		db: db,
	}
}

func scopeToRecord(scope domain.Scope) scopeRecord {
	return scopeRecord{
		ID:          scope.ID,
		ProjectID:   scope.ProjectID,
		Name:        scope.Name,
		Description: sql.NullString{String: scope.Description, Valid: scope.Description != ""},
		CreatedAt:   scope.CreatedAt,
		UpdatedAt:   sql.NullTime{Time: scope.UpdatedAt, Valid: !scope.UpdatedAt.IsZero()},
	}
}

func (s *ScopeRepo) Create(ctx context.Context, scope domain.Scope) (domain.Scope, error) {
	record := scopeToRecord(scope)
	record.ID = 0
	if err := gorm.G[scopeRecord](s.db.From(ctx)).Create(ctx, &record); err != nil {
		return domain.Scope{}, err
	}
	return scopeRecordToDomain(record), nil
}

func (s *ScopeRepo) Update(ctx context.Context, scope domain.Scope) error {
	if scope.ID == 0 {
		return errs.Validation("scope_id_required", "scope id is required to update a scope")
	}
	record := scopeToRecord(scope)
	res := s.db.From(ctx).Save(&record)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errs.NotFound("scope_not_found", "scope not found")
	}
	return nil
}

func (s *ScopeRepo) ListByProject(ctx context.Context, projectId int) ([]domain.Scope, error) {
	records, err := gorm.G[scopeRecord](s.db.From(ctx)).
		Preload("Stages", nil).
		Where("project_id = ?", projectId).
		Find(ctx)
	if err != nil {
		return nil, err
	}
	scopes := make([]domain.Scope, len(records))
	for i, record := range records {
		scopes[i] = scopeRecordToDomain(record)
	}
	return scopes, nil
}
