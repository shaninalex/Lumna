package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/core/actor"
	"gitlab.com/shaninalex/lumna/app/modules/identity/contract"
	"gitlab.com/shaninalex/lumna/app/modules/identity/internal/domain"
)

type AccountDelete struct {
	repo domain.IdentityRepo
}

func NewAccountDelete(repo domain.IdentityRepo) *AccountDelete {
	return &AccountDelete{repo: repo}
}

func (s *AccountDelete) Handle(ctx context.Context, cmd contract.AccountDelete) (bool, error) {
	a, ok := actor.From(ctx)
	if !ok && a.IdentityID != cmd.IdentityID {
		return false, domain.ErrBadCredentials
	}
	if err := s.repo.Delete(ctx, cmd.IdentityID); err != nil {
		return false, err
	}
	return true, nil
}
