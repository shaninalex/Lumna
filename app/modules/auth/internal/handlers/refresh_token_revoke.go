package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/core/actor"
	"gitlab.com/shaninalex/lumna/app/modules/auth/contract"
	"gitlab.com/shaninalex/lumna/app/modules/auth/internal/domain"
)

type RefreshTokenRevoke struct {
	repo domain.RefreshTokenRepo
}

func NewRefreshTokenRevoke(repo domain.RefreshTokenRepo) *RefreshTokenRevoke {
	return &RefreshTokenRevoke{repo}
}

func (s *RefreshTokenRevoke) Handle(ctx context.Context, cmd contract.RefreshTokenRevoke) (contract.RefreshTokenView, error) {
	token, err := s.repo.ByID(ctx, cmd.TokenID)
	if err != nil {
		return contract.RefreshTokenView{}, err
	}

	a, ok := actor.From(ctx)
	if !ok && a.IdentityID != token.IdentityID {
		return contract.RefreshTokenView{}, domain.ErrTokenNotBelongs
	}

	token.Revoke()
	// NOTE: we have to add revoked_at

	if token, err = s.repo.Save(ctx, token); err != nil {
		return contract.RefreshTokenView{}, err
	}

	return contract.RefreshTokenView{
		ID:         token.ID,
		IdentityID: token.IdentityID,
		ClientID:   token.ClientID,
		Scopes:     token.Scopes,
		ExpiresAt:  token.ExpiresAt,
		Revoked:    token.Revoked,
		CreatedAt:  token.CreatedAt,
	}, nil
}
