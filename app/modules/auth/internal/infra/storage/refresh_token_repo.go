package storage

import (
	"context"
	"database/sql"
	"errors"

	"gitlab.com/shaninalex/lumna/app/modules/auth/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/database"
	"gorm.io/gorm"
)

// RefreshTokenRepo implements [domain.RefreshTokenRepo].
type RefreshTokenRepo struct{ db *database.DB }

var _ domain.RefreshTokenRepo = (*RefreshTokenRepo)(nil)

func NewRefreshTokenRepo(db *database.DB) *RefreshTokenRepo {
	return &RefreshTokenRepo{db: db}
}

func (r *RefreshTokenRepo) Save(ctx context.Context, t domain.RefreshToken) (domain.RefreshToken, error) {
	record := refreshTokenRecord{
		ID:         t.ID,
		IdentityID: t.IdentityID,
		Hash:       t.Hash,
		ClientID:   sql.NullString{String: t.ClientID, Valid: t.ClientID != ""},
		Scopes:     sql.NullString{String: t.Scopes, Valid: t.Scopes != ""},
		ExpiresAt:  t.ExpiresAt,
		Revoked:    t.Revoked,
		CreatedAt:  t.CreatedAt,
	}
	if err := r.db.From(ctx).Save(&record).Error; err != nil {
		return domain.RefreshToken{}, err
	}

	return toDomainRefreshToken(record), nil
}

func (r *RefreshTokenRepo) ByHash(ctx context.Context, hash string) (domain.RefreshToken, error) {
	rec, err := gorm.G[refreshTokenRecord](r.db.From(ctx)).Where("hash = ?", hash).First(ctx)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.RefreshToken{}, domain.ErrTokenUnknown
	}
	if err != nil {
		return domain.RefreshToken{}, err
	}
	return toDomainRefreshToken(rec), nil
}
