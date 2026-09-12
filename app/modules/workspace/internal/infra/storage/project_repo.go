package storage

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/workspace/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/database"
	"gorm.io/gorm"
)

type ProjectRepo struct {
	db *database.DB
}

var _ domain.ProjectRepo = (*ProjectRepo)(nil)

func NewProjectRepo(db *database.DB) *ProjectRepo {
	return &ProjectRepo{db: db}
}

func (s ProjectRepo) Save(ctx context.Context, t *domain.Project) error {
	record := projectRecord{
		Title:       t.Title,
		WorkspaceId: t.WorkspaceId,
		OwnerId:     t.OwnerId,
		Meta:        t.Meta,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
	if err := s.db.From(ctx).Save(&record).Error; err != nil {
		return err
	}
	t.ID = record.ID
	return nil
}

func (s ProjectRepo) ByWorkspaceId(ctx context.Context, workspaceId int) ([]domain.Project, error) {
	result, err := gorm.G[projectRecord](s.db.From(ctx)).Where("workspace_id = ?", workspaceId).Find(ctx)
	if err != nil {
		return nil, err
	}

	projects := make([]domain.Project, len(result))
	for i := range result {
		projects[i] = toDomainProject(result[i])
	}
	return projects, nil
}
