package storage

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/workspace/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/database"
	"gorm.io/gorm"
)

type IdentityWorkspaceRepo struct {
	db *database.DB
}

var _ domain.IdentityWorkspaceRepo = (*IdentityWorkspaceRepo)(nil)

func NewIdentityWorkspaceRepo(db *database.DB) *IdentityWorkspaceRepo {
	return &IdentityWorkspaceRepo{db: db}
}

func (s *IdentityWorkspaceRepo) Save(ctx context.Context, t domain.IdentityWorkspace) (domain.IdentityWorkspace, error) {
	record := identityWorkspaceRecord{
		IdentityID:  t.IdentityID,
		WorkspaceID: t.WorkspaceID,
		CreatedAt:   t.CreatedAt,
		Role:        t.Role,
	}
	if err := s.db.From(ctx).Save(&record).Error; err != nil {
		return domain.IdentityWorkspace{}, err
	}
	return domain.IdentityWorkspace{
		IdentityID:  record.IdentityID,
		WorkspaceID: record.WorkspaceID,
		CreatedAt:   record.CreatedAt,
		Role:        t.Role,
	}, nil
}

func (s *IdentityWorkspaceRepo) ByWorkspaceId(ctx context.Context, workspaceId int) ([]domain.IdentityWorkspace, error) {
	records, err := gorm.G[identityWorkspaceRecord](s.db.From(ctx)).Where("workspace_id = ?", workspaceId).Find(ctx)
	if err != nil {
		return []domain.IdentityWorkspace{}, err
	}
	identities := make([]domain.IdentityWorkspace, len(records))
	for i, record := range records {
		identities[i] = domain.IdentityWorkspace{
			IdentityID:  record.IdentityID,
			WorkspaceID: record.WorkspaceID,
			CreatedAt:   record.CreatedAt,
			Role:        record.Role,
		}
	}
	return identities, nil
}
