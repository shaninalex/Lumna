package storage

import (
	"context"
	"database/sql"
	"errors"

	"gitlab.com/shaninalex/lumna/app/modules/tracker/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/database"
	"gorm.io/gorm"
)

type StageRepo struct {
	db *database.DB
}

var _ domain.StageRepo = (*StageRepo)(nil)

func NewStageRepo(db *database.DB) *StageRepo {
	return &StageRepo{
		db: db,
	}
}

func (s *StageRepo) Save(ctx context.Context, stage *domain.Stage) error {
	record := stageRecord{
		ID:          stage.ID,
		ScopeID:     stage.ScopeID,
		Name:        stage.Name,
		Description: sql.NullString{String: stage.Description, Valid: stage.Description != ""},
		Category:    string(stage.Category),
		Position:    stage.Position,
		WipLimit:    sql.NullInt32{Int32: int32(stage.WIPLimit), Valid: stage.WIPLimit != 0},
		CreatedAt:   stage.CreatedAt,
		UpdatedAt:   sql.NullTime{Time: stage.UpdatedAt, Valid: stage.UpdatedAt.IsZero()},
	}
	if err := s.db.From(ctx).Save(&record).Error; err != nil {
		return err
	}
	stage.ID = record.ID
	return nil
}

func (s *StageRepo) GetById(ctx context.Context, stageId int) (*domain.Stage, error) {
	record, err := gorm.G[stageRecord](s.db.From(ctx)).
		Preload("WorkItems", nil).
		Where("id = ?", stageId).
		First(ctx)
	if err != nil {
		return nil, err
	}
	w := stageRecordToDomain(record)
	return &w, nil
}

func (s *StageRepo) Get(ctx context.Context, scopeId int) ([]domain.Stage, error) {
	records, err := gorm.G[stageRecord](s.db.From(ctx)).
		Preload("WorkItems", nil).
		Where("scope_id = ?", scopeId).
		Find(ctx)
	if err != nil {
		return nil, err
	}
	var stages []domain.Stage
	for _, record := range records {
		stages = append(stages, stageRecordToDomain(record))
	}
	return stages, nil
}

func (s *StageRepo) Delete(ctx context.Context, stageId int) (bool, error) {
	r, err := gorm.G[stageRecord](s.db.From(ctx)).Where("id = ?", stageId).Delete(ctx)
	if err != nil {
		return false, err
	}
	if r <= 0 {
		return false, errors.New("unable to delete stage, something went wrong")
	}
	return true, nil
}
