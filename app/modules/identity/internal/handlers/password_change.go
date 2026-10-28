package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/core/actor"
	"gitlab.com/shaninalex/lumna/app/modules/identity/contract"
	"gitlab.com/shaninalex/lumna/app/modules/identity/internal/domain"
)

type PasswordChange struct {
	identityRepo domain.IdentityRepo
	credsRepo    domain.CredentialRepo
	hasher       domain.Hasher
}

func NewPasswordChange(
	identityRepo domain.IdentityRepo,
	credsRepo domain.CredentialRepo,
	hasher domain.Hasher,
) *PasswordChange {
	return &PasswordChange{identityRepo, credsRepo, hasher}
}

func (s *PasswordChange) Handle(ctx context.Context, cmd contract.PasswordChange) (bool, error) {
	identity, err := s.identityRepo.ByID(ctx, cmd.IdentityId)
	if err != nil {
		return false, err
	}

	// change password should be ONLY via authorized action
	a, ok := actor.From(ctx)
	if !ok || a.IdentityID == 0 || a.IdentityID != identity.ID {
		return false, domain.ErrBadCredentials
	}

	hash, err := s.hasher.Hash(cmd.Password.Reveal())
	if err != nil {
		return false, err
	}

	credential, err := s.credsRepo.ByIdentityAndProvider(ctx, cmd.IdentityId, "email")
	if err != nil {
		return false, err
	}

	credential.PasswordHash = hash
	if err = s.credsRepo.Update(ctx, credential); err != nil {
		return false, err
	}

	return true, nil
}
