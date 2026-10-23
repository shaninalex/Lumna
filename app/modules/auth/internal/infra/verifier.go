package infra

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/auth/contract"
	"gitlab.com/shaninalex/lumna/app/modules/auth/internal/domain"
)

type Verifier[P, T any] struct {
	codec domain.TokenCodec[P]
	view  func(P) T
}

var _ contract.Verifier[int] = (*Verifier[struct{}, int])(nil)

func NewVerifier[P, T any](c domain.TokenCodec[P], view func(P) T) *Verifier[P, T] {
	return &Verifier[P, T]{
		codec: c,
		view:  view,
	}
}

func (v *Verifier[P, T]) Verify(_ context.Context, token string) (T, error) {
	var zero T

	if token == "" {
		return zero, domain.ErrBadAccessToken
	}

	p, err := v.codec.Parse(token)
	if err != nil {
		return zero, domain.ErrBadAccessToken
	}

	return v.view(p), nil
}

func Identity[T any](p T) T { return p }
