package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/identity/contract"
	"gitlab.com/shaninalex/lumna/app/modules/identity/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

type Register struct {
	identities domain.IdentityRepo
	creds      domain.CredentialRepo
	hasher     domain.Hasher
	mailer     contract.Mailer
	clock      clock.Clock
}

func NewRegister(
	i domain.IdentityRepo,
	c domain.CredentialRepo,
	h domain.Hasher,
	m contract.Mailer,
	clk clock.Clock,
) *Register {
	return &Register{i, c, h, m, clk}
}

// Handle - executes ExecRegister command
func (u *Register) Handle(ctx context.Context, cmd contract.Register) (contract.ProfileView, error) {
	var zero contract.ProfileView

	if existing, err := u.identities.ByEmail(ctx, cmd.Email); err == nil && existing.ID != 0 {
		return zero, domain.ErrEmailTaken
	}

	identity, err := u.identities.Save(ctx, domain.NewIdentity(cmd.Email, cmd.FullName, u.clock.Now()))
	if err != nil {
		return zero, err
	}

	hash, err := u.hasher.Hash(cmd.Password.Reveal())
	if err != nil {
		return zero, err
	}
	_, err = u.creds.Save(ctx, domain.NewPasswordCredential(identity.ID, hash, identity.Email))
	if err != nil {
		return zero, err
	}

	//if err := outbox.Record(ctx, contract.Registered{
	//	IdentityID: ident.ID, Email: ident.Email, At: u.clock.Now(),
	//}); err != nil {
	//	return zero, err
	//}

	return toProfileView(identity), nil
}
