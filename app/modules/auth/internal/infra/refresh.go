package infra

import (
	"context"
	"fmt"
	"time"

	"gitlab.com/shaninalex/lumna/app/core/secure"
	"gitlab.com/shaninalex/lumna/app/modules/auth/contract"
	"gitlab.com/shaninalex/lumna/app/modules/auth/internal/domain"
)

type RefreshTokens struct {
	ttl    time.Duration
	repo   domain.RefreshTokenRepo
	access domain.TokenCodec[domain.AccessClaims]
}

var _ domain.RefreshTokens = (*RefreshTokens)(nil)

func NewRefreshTokens(
	ttl time.Duration,
	repo domain.RefreshTokenRepo,
	access domain.TokenCodec[domain.AccessClaims],
) *RefreshTokens {
	return &RefreshTokens{
		ttl:    ttl,
		repo:   repo,
		access: access,
	}
}

func (r *RefreshTokens) Issue() (string, string, time.Duration, error) {
	plain, err := secure.NewOpaque()
	if err != nil {
		return "", "", 0, fmt.Errorf("auth: generate refresh token: %w", err)
	}
	return plain, r.Hash(plain), r.ttl, nil
}

func (r *RefreshTokens) Hash(plain string) string {
	return secure.HashOpaque(plain)
}

func (r *RefreshTokens) Refresh(ctx context.Context, plain string, now time.Time) (contract.SessionView, error) {
	var zero contract.SessionView

	if plain == "" {
		return zero, domain.ErrEmptyToken
	}

	stored, err := r.repo.ByHash(ctx, r.Hash(plain))
	if err != nil {
		return zero, err
	}

	if err := stored.Usable(now); err != nil {
		return zero, err
	}

	access, accessTTL, err := r.access.Issue(domain.AccessClaims{IdentityID: stored.IdentityID}, now)
	if err != nil {
		return zero, err
	}

	return contract.SessionView{
		IdentityID:           stored.IdentityID,
		AccessToken:          access,
		AccessTokenDuration:  accessTTL,
		RefreshToken:         plain,
		RefreshTokenDuration: stored.ExpiresAt.Sub(now),
	}, nil
}
