package storage

import (
	"context"
	"database/sql"
	"errors"
	"time"

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

func (r *RefreshTokenRepo) Expired(ctx context.Context, now time.Time) ([]domain.RefreshToken, error) {
	records, err := gorm.G[refreshTokenRecord](r.db.From(ctx)).Where("expires_at <= ?", now).Find(ctx)
	if err != nil {
		return nil, err
	}

	tokens := make([]domain.RefreshToken, 0, len(records))
	for _, rec := range records {
		tokens = append(tokens, toDomainRefreshToken(rec))
	}

	return tokens, nil
}

func (r *RefreshTokenRepo) Remove(ctx context.Context, id int) error {
	rows, err := gorm.G[refreshTokenRecord](r.db.From(ctx)).Where("id = ?", id).Delete(ctx)
	if rows == 0 {
		return domain.ErrTokenUnknown
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.ErrNotFound
	}
	if err != nil {
		return err
	}
	return nil
}

func (r *RefreshTokenRepo) RemoveBatch(ctx context.Context, ids []int) error {
	if len(ids) == 0 {
		return nil
	}
	rows, err := gorm.G[refreshTokenRecord](r.db.From(ctx)).Where("id IN ?", ids).Delete(ctx)
	if err != nil {
		return err
	}
	if rows == 0 {
		return domain.ErrTokenUnknown
	}

	return nil
}

func (r *RefreshTokenRepo) ListIdentityID(ctx context.Context, id int) ([]domain.RefreshToken, error) {
	records, err := gorm.G[refreshTokenRecord](r.db.From(ctx)).Where("identity_id = ?", id).Find(ctx)
	if err != nil {
		return nil, err
	}

	tokens := make([]domain.RefreshToken, 0, len(records))
	for _, rec := range records {
		tokens = append(tokens, toDomainRefreshToken(rec))
	}

	return tokens, nil
}

func (r *RefreshTokenRepo) ByID(ctx context.Context, id int) (domain.RefreshToken, error) {
	record, err := gorm.G[refreshTokenRecord](r.db.From(ctx)).Where("id = ?", id).First(ctx)
	if err != nil {
		return domain.RefreshToken{}, err
	}

	return toDomainRefreshToken(record), nil
}
