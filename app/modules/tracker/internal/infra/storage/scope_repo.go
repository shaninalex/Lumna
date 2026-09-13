package storage

import (
	"context"

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

func (s ScopeRepo) Save(ctx context.Context, scope *domain.Scope) error {
	record := scopeRecord{
		ProjectID:   scope.ProjectID,
		Name:        scope.Name,
		Description: &scope.Description,
		CreatedAt:   scope.CreatedAt,
		UpdatedAt:   scope.UpdatedAt,
	}
	if err := s.db.From(ctx).Save(&record).Error; err != nil {
		return err
	}

	scope.ID = record.ID
	return nil
}

func (s ScopeRepo) Get(ctx context.Context, projectId int) ([]domain.Scope, error) {
	records, err := gorm.G[scopeRecord](s.db.From(ctx)).
		Preload("Stages", nil).
		Where("project_id = ?", projectId).
		Find(ctx)
	if err != nil {
		return nil, err
	}
	var result []domain.Scope
	for _, record := range records {
		var stages []domain.Stage
		for _, stage := range record.Stages {
			st := domain.Stage{
				ID:          stage.ID,
				ScopeID:     stage.ScopeID,
				Name:        stage.Name,
				Description: *stage.Description,
				Position:    stage.Position,
				WIPLimit:    stage.WIPLimit,
				CreatedAt:   stage.CreatedAt,
				UpdatedAt:   *stage.UpdatedAt,
			}
			if stage.Category != nil {
				st.Category = domain.StageCategory(*stage.Category)
			}
			stages = append(stages, st)
		}
		result = append(result, domain.Scope{
			ID:          record.ID,
			ProjectID:   record.ProjectID,
			Name:        record.Name,
			Stages:      stages,
			Description: *record.Description,
			CreatedAt:   record.CreatedAt,
			UpdatedAt:   record.UpdatedAt,
		})
	}

	return result, nil
}
