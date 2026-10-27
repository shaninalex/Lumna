package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/auth/contract"
	"gitlab.com/shaninalex/lumna/app/modules/auth/internal/domain"
)

type RefreshTokenList struct {
	repo domain.RefreshTokenRepo
}

func NewRefreshTokenList(repo domain.RefreshTokenRepo) *RefreshTokenList {
	return &RefreshTokenList{repo}
}

func (s *RefreshTokenList) Handle(ctx context.Context, query contract.RefreshTokensList) ([]contract.RefreshTokenView, error) {
	tokens, err := s.repo.ListIdentityID(ctx, query.IdentityID)
	if err != nil {
		return []contract.RefreshTokenView{}, nil
	}

	views := make([]contract.RefreshTokenView, 0, len(tokens))
	for _, t := range tokens {
		views = append(views, contract.RefreshTokenView{
			ID:         t.ID,
			IdentityID: t.IdentityID,
			ClientID:   t.ClientID,
			Scopes:     t.Scopes,
			ExpiresAt:  t.ExpiresAt,
			Revoked:    t.Revoked,
			CreatedAt:  t.CreatedAt,
		})
	}

	return views, nil
}
