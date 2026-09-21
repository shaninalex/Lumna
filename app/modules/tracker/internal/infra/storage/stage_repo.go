package storage

import (
	"context"
	"database/sql"
	"errors"

	"gitlab.com/shaninalex/lumna/app/core/errs"
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

func stageToRecord(stage domain.Stage) stageRecord {
	return stageRecord{
		ID:          stage.ID,
		ScopeID:     stage.ScopeID,
		Name:        stage.Name,
		Description: sql.NullString{String: stage.Description, Valid: stage.Description != ""},
		Category:    string(stage.Category),
		Position:    stage.Position,
		WipLimit:    sql.NullInt32{Int32: int32(stage.WIPLimit), Valid: stage.WIPLimit != 0},
		CreatedAt:   stage.CreatedAt,
		UpdatedAt:   sql.NullTime{Time: stage.UpdatedAt, Valid: !stage.UpdatedAt.IsZero()},
	}
}

func (s *StageRepo) Create(ctx context.Context, stage domain.Stage) (domain.Stage, error) {
	record := stageToRecord(stage)
	record.ID = 0
	if err := gorm.G[stageRecord](s.db.From(ctx)).Create(ctx, &record); err != nil {
		return domain.Stage{}, err
	}
	return stageRecordToDomain(record), nil
}

func (s *StageRepo) Update(ctx context.Context, stage domain.Stage) error {
	if stage.ID == 0 {
		return errs.Validation("stage_id_required", "stage id is required to update a stage")
	}
	record := stageToRecord(stage)
	res := s.db.From(ctx).Save(&record)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errs.NotFound("stage_not_found", "stage not found")
	}
	return nil
}

func (s *StageRepo) Get(ctx context.Context, stageId int) (domain.Stage, error) {
	record, err := gorm.G[stageRecord](s.db.From(ctx)).
		Preload("WorkItems", nil).
		Where("id = ?", stageId).
		First(ctx)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.Stage{}, errs.NotFound("stage_not_found", "stage not found")
	}
	if err != nil {
		return domain.Stage{}, err
	}
	return stageRecordToDomain(record), nil
}

func (s *StageRepo) ListByScope(ctx context.Context, scopeId int) ([]domain.Stage, error) {
	records, err := gorm.G[stageRecord](s.db.From(ctx)).
		Preload("WorkItems", nil).
		Where("scope_id = ?", scopeId).
		Find(ctx)
	if err != nil {
		return nil, err
	}
	stages := make([]domain.Stage, len(records))
	for i, record := range records {
		stages[i] = stageRecordToDomain(record)
	}
	return stages, nil
}

func (s *StageRepo) Delete(ctx context.Context, stageId int) error {
	r, err := gorm.G[stageRecord](s.db.From(ctx)).Where("id = ?", stageId).Delete(ctx)
	if err != nil {
		return err
	}
	if r == 0 {
		return errs.NotFound("stage_not_found", "stage not found")
	}
	return nil
}
