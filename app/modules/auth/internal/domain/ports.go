package domain

import (
	"context"
	"time"
)

type RefreshTokenRepo interface {
	Save(ctx context.Context, t *RefreshToken) error
	ByHash(ctx context.Context, hash string) (*RefreshToken, error)
}

type Tokens interface {
	IssueAccess(identityID int, now time.Time) (token string, ttl time.Duration, err error)
	IssueRefresh() (plain string, hash string, ttl time.Duration, err error)
	HashRefresh(plain string) string
}
