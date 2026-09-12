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

func (s IdentityWorkspaceRepo) Save(ctx context.Context, t *domain.IdentityWorkspace) error {
	record := identityWorkspaceRecord{
		IdentityID:  t.IdentityID,
		WorkspaceID: t.WorkspaceID,
		CreatedAt:   t.CreatedAt,
	}
	if err := s.db.From(ctx).Save(&record).Error; err != nil {
		return err
	}
	return nil
}

func (s IdentityWorkspaceRepo) ByIdentityId(ctx context.Context, identityId int) ([]domain.IdentityWorkspace, error) {
	//TODO implement me
	panic("implement me")
}

func (s IdentityWorkspaceRepo) ByWorkspaceId(ctx context.Context, workspaceId int) ([]domain.IdentityWorkspace, error) {
	//TODO implement me
	panic("implement me")
}

func (s IdentityWorkspaceRepo) Delete(ctx context.Context, identityId, workspaceId int) (bool, error) {
	//TODO implement me
	panic("implement me")
}
