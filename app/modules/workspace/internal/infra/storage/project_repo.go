package storage

import (
	"context"
	"errors"

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

func (s *ProjectRepo) Save(ctx context.Context, project domain.Project) (domain.Project, error) {
	record := toRecordProject(project)
	if err := s.db.From(ctx).Save(&record).Error; err != nil {
		return domain.Project{}, err
	}
	return toDomainProject(record), nil
}

func (s *ProjectRepo) ByWorkspaceId(ctx context.Context, workspaceId int) ([]domain.Project, error) {
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

func (s *ProjectRepo) Update(ctx context.Context, project domain.Project) (domain.Project, error) {
	record := toRecordProject(project)
	if err := s.db.From(ctx).Where("id = ?", project.ID).Updates(&record).Error; err != nil {
		return domain.Project{}, err
	}
	return toDomainProject(record), nil
}

func (s *ProjectRepo) GetById(ctx context.Context, projectID int) (domain.Project, error) {
	result, err := gorm.G[projectRecord](s.db.From(ctx)).Where("id = ?", projectID).First(ctx)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.Project{}, domain.ErrProjectNotFound
	}
	if err != nil {
		return domain.Project{}, err
	}
	return toDomainProject(result), nil
}

func (s *ProjectRepo) Remove(ctx context.Context, projectID int) error {
	record, err := gorm.G[projectRecord](s.db.From(ctx)).Where("id = ?", projectID).Delete(ctx)
	if errors.Is(err, gorm.ErrRecordNotFound) || record == 0 {
		return domain.ErrProjectNotFound
	}
	if err != nil {
		return err
	}
	return nil
}
