package infra

import (
	"fmt"
	"time"

	"gitlab.com/shaninalex/lumna/app/core/secure"
	"gitlab.com/shaninalex/lumna/app/modules/auth/internal/domain"
)

type RefreshTokens struct {
	ttl time.Duration
}

var _ domain.RefreshTokens = (*RefreshTokens)(nil)

func NewRefreshTokens(ttl time.Duration) *RefreshTokens {
	return &RefreshTokens{ttl: ttl}
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
