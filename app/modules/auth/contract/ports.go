package contract

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/core/actor"
	"gitlab.com/shaninalex/lumna/app/core/bus"
)

type IdentityAuthenticator interface {

	// VerifyPassword authenticate user by email/password
	VerifyPassword(ctx context.Context, email string, password bus.Secret) (int, error)
}

type IdentityProvisioner interface {

	// EnsureIdentityByEmail - Get or create identity. Need for OAuth flow.
	EnsureIdentityByEmail(ctx context.Context, email, fullName string) (int, error)
}

// Verifier turns an access token into the actor it stands for.
type Verifier interface {
	Verify(ctx context.Context, accessToken string) (actor.Actor, error)
}
