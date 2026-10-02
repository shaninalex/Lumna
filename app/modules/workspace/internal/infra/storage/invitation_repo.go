package storage

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/workspace/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
	"gitlab.com/shaninalex/lumna/app/platform/database"
	"gorm.io/gorm"
)

type InvitationRepo struct {
	db    *database.DB
	clock clock.Clock
}

func NewInvitationRepo(db *database.DB, clock clock.Clock) *InvitationRepo {
	return &InvitationRepo{db: db, clock: clock}
}

func (s *InvitationRepo) Save(ctx context.Context, invitation domain.Invitation) (domain.Invitation, error) {
	invitation.ID = 0
	invitation.CreatedAt = s.clock.Now()
	if err := gorm.G[domain.Invitation](s.db.From(ctx)).Create(ctx, &invitation); err != nil {
		return domain.Invitation{}, err
	}
	return domain.Invitation{}, nil
}

func (s *InvitationRepo) GetByHash(ctx context.Context, hash string) (domain.Invitation, error) {
	return gorm.G[domain.Invitation](s.db.From(ctx)).
		Where("token_hash = ?", hash).
		First(ctx)
}

func (s *InvitationRepo) ListByWorkspaceId(ctx context.Context, workspaceId int) ([]domain.Invitation, error) {
	return gorm.G[domain.Invitation](s.db.From(ctx)).Where("workspace_id = ?", workspaceId).Find(ctx)
}
