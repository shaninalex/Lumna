package storage

import (
	"context"
	"errors"

	"gitlab.com/shaninalex/lumna/app/modules/workspace/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/database"
	"gorm.io/gorm"
)

type WorkspaceRepo struct {
	db *database.DB
}

var _ domain.WorkspaceRepo = (*WorkspaceRepo)(nil)

func NewWorkspaceRepo(db *database.DB) *WorkspaceRepo { return &WorkspaceRepo{db: db} }

func (s WorkspaceRepo) Save(ctx context.Context, t *domain.Workspace) error {
	record := workspaceRecord{
		ID:         t.ID,
		Title:      t.Title,
		OwnerEmail: t.OwnerEmail,
		Active:     t.Active,
		CreatedAt:  t.CreatedAt,
	}
	if err := s.db.From(ctx).Save(&record).Error; err != nil {
		return err
	}

	t.ID = record.ID
	return nil
}

func (s WorkspaceRepo) ById(ctx context.Context, id int) (*domain.Workspace, error) {
	rec, err := gorm.G[workspaceRecord](s.db.From(ctx)).Where("id = ?", id).First(ctx)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrWorkspaceNotFound
	}
	if err != nil {
		return nil, err
	}
	w := toDomainWorkspace(rec)
	return &w, nil
}

func (s WorkspaceRepo) List(ctx context.Context) ([]domain.Workspace, error) {
	rec, err := gorm.G[workspaceRecord](s.db.From(ctx)).Find(ctx)
	if err != nil {
		return nil, err
	}

	workspaces := make([]domain.Workspace, len(rec))
	for i := range rec {
		workspaces[i] = toDomainWorkspace(rec[i])
	}
	return workspaces, nil
}
