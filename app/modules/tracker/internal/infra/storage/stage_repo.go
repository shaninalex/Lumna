package storage

import (
	"context"

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
		Description: &stage.Description,
		Category:    string(stage.Category),
		Position:    stage.Position,
		WipLimit:    stage.WIPLimit,
		CreatedAt:   stage.CreatedAt,
		UpdatedAt:   &stage.UpdatedAt,
	}
	if err := s.db.From(ctx).Save(&record).Error; err != nil {
		return err
	}
	stage.ID = record.ID
	return nil
}

func (s *StageRepo) GetById(ctx context.Context, stageId int) (*domain.Stage, error) {
	record, err := gorm.G[stageRecord](s.db.From(ctx)).Where("id = ?", stageId).First(ctx)
	if err != nil {
		return nil, err
	}
	return &domain.Stage{
		ID:          record.ID,
		ScopeID:     record.ScopeID,
		Name:        record.Name,
		Description: *record.Description,
		Category:    domain.StageCategory(record.Category),
		Position:    record.Position,
		WIPLimit:    record.WipLimit,
		CreatedAt:   record.CreatedAt,
		UpdatedAt:   *record.UpdatedAt,
	}, nil
}

func (s *StageRepo) Get(ctx context.Context, scopeId int) ([]domain.Stage, error) {
	records, err := gorm.G[stageRecord](s.db.From(ctx)).Where("scope_id = ?", scopeId).Find(ctx)
	if err != nil {
		return nil, err
	}
	var stages []domain.Stage
	for _, record := range records {
		stages = append(stages, domain.Stage{
			ID:          record.ID,
			ScopeID:     record.ScopeID,
			Name:        record.Name,
			Description: *record.Description,
			Category:    domain.StageCategory(record.Category),
			Position:    record.Position,
			WIPLimit:    record.WipLimit,
			CreatedAt:   record.CreatedAt,
			UpdatedAt:   *record.UpdatedAt,
		})
	}
	return stages, nil
}
