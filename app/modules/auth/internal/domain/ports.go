package domain

import (
	"context"
	"time"
)

type RefreshTokenRepo interface {
	Save(ctx context.Context, t RefreshToken) (RefreshToken, error)
	ByHash(ctx context.Context, hash string) (RefreshToken, error)
	Expired(ctx context.Context, now time.Time) ([]RefreshToken, error)
	Remove(ctx context.Context, id int) error
	RemoveBatch(ctx context.Context, ids []int) error
	ListIdentityID(ctx context.Context, identityID int) ([]RefreshToken, error)
	ByID(ctx context.Context, id int) (RefreshToken, error)
}

type TokenCodec[P any] interface {
	Issue(p P, now time.Time) (token string, ttl time.Duration, err error)
	Parse(token string) (P, error)
}

type RefreshTokens interface {
	Issue() (plain, hash string, ttl time.Duration, err error)
	Hash(plain string) string
}
