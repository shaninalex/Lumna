package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/core/actor"
	"gitlab.com/shaninalex/lumna/app/modules/auth/contract"
	"gitlab.com/shaninalex/lumna/app/modules/auth/internal/domain"
)

type RefreshTokenDelete struct {
	repo domain.RefreshTokenRepo
}

func NewRefreshTokenDelete(repo domain.RefreshTokenRepo) *RefreshTokenDelete {
	return &RefreshTokenDelete{repo}
}

func (s *RefreshTokenDelete) Handle(ctx context.Context, cmd contract.RefreshTokenDelete) (bool, error) {
	token, err := s.repo.ByID(ctx, cmd.TokenID)
	if err != nil {
		return false, err
	}

	a, ok := actor.From(ctx)
	if !ok && a.IdentityID != token.IdentityID {
		return false, domain.ErrTokenNotBelongs
	}

	if err = s.repo.Remove(ctx, token.ID); err != nil {
		return false, err
	}

	return true, nil
}
