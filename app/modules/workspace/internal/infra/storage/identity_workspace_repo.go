package storage

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/workspace/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/database"
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
