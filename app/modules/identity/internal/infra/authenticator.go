package infra

import (
	"context"
	"errors"

	"gitlab.com/shaninalex/lumna/app/core/bus"
	"gitlab.com/shaninalex/lumna/app/modules/identity/internal/domain"
)

// Authenticator implements contract.Authenticator
type Authenticator struct {
	identities  domain.IdentityRepo
	credentials domain.CredentialRepo
	hasher      domain.Hasher
}

func NewAuthenticator(
	i domain.IdentityRepo,
	c domain.CredentialRepo,
	h domain.Hasher,
) *Authenticator {
	return &Authenticator{identities: i, credentials: c, hasher: h}
}

// VerifyPassword returns the identity id on success.
func (a *Authenticator) VerifyPassword(ctx context.Context, email string, password bus.Secret) (int, error) {
	ident, err := a.identities.ByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return 0, domain.ErrBadCredentials
		}
		return 0, err
	}

	if !ident.Active {
		return 0, domain.ErrNotActive
	}

	cred, err := a.credentials.ByIdentityAndProvider(ctx, ident.ID, domain.EmailCredentialProvider)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return 0, domain.ErrBadCredentials
		}
		return 0, err
	}

	if cred.PasswordHash == "" {
		return 0, domain.ErrBadCredentials
	}

	if err := a.hasher.Verify(cred.PasswordHash, password.Reveal()); err != nil {
		return 0, domain.ErrBadCredentials
	}

	return ident.ID, nil
}
