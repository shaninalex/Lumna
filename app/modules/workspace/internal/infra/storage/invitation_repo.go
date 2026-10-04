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

func (s *InvitationRepo) Create(ctx context.Context, invitation domain.Invitation) (domain.Invitation, error) {
	invitation.ID = 0
	invitation.CreatedAt = s.clock.Now()
	record := toInvitationRecord(invitation)
	if err := gorm.G[workspaceInvitationRecord](s.db.From(ctx)).Create(ctx, &record); err != nil {
		return domain.Invitation{}, err
	}
	return toInvitationDomain(record), nil
}

func (s *InvitationRepo) GetByHash(ctx context.Context, hash string) (domain.Invitation, error) {
	record, err := gorm.G[workspaceInvitationRecord](s.db.From(ctx)).
		Where("token_hash = ?", hash).
		First(ctx)
	if err != nil {
		return domain.Invitation{}, err
	}
	return toInvitationDomain(record), err
}

func (s *InvitationRepo) ListByWorkspaceId(ctx context.Context, workspaceId int) ([]domain.Invitation, error) {
	records, err := gorm.G[workspaceInvitationRecord](s.db.From(ctx)).
		Where("workspace_id = ?", workspaceId).
		Find(ctx)
	if err != nil {
		return []domain.Invitation{}, err
	}
	invitations := make([]domain.Invitation, len(records))
	for i, inv := range records {
		invitations[i] = toInvitationDomain(inv)
	}
	return invitations, nil
}
