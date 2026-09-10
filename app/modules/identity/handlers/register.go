package handlers

import (
	"context"
	"time"

	"gitlab.com/shaninalex/lumna/app/modules/identity/contract"
	"gitlab.com/shaninalex/lumna/app/modules/identity/internal/domain"
)

type Register struct {
	identities domain.IdentityRepo
	creds      domain.CredentialRepo
	hasher     domain.Hasher
	mailer     contract.Mailer
}

func NewRegister(
	i domain.IdentityRepo,
	c domain.CredentialRepo,
	h domain.Hasher,
	m contract.Mailer,
) *Register {
	return &Register{i, c, h, m}
}

// Handle — implements bus.RegisterCommand.
func (u *Register) Handle(ctx context.Context, cmd contract.Register) (contract.ProfileView, error) {
	var zero contract.ProfileView

	if existing, err := u.identities.ByEmail(ctx, cmd.Email); err == nil && existing != nil {
		return zero, domain.ErrEmailTaken
	}

	ident, err := domain.NewIdentity(cmd.Email, cmd.FullName, time.Now())
	if err != nil {
		return zero, err
	}

	// Транзакція вже відкрита middleware — окремий tx-код тут не потрібен.
	if err := u.identities.Save(ctx, ident); err != nil {
		return zero, err
	}

	hash, err := u.hasher.Hash(cmd.Password.Reveal())
	if err != nil {
		return zero, err
	}
	cred := domain.NewPasswordCredential(ident.ID, hash, ident.Email)
	if err := u.creds.Save(ctx, cred); err != nil {
		return zero, err
	}

	//if err := outbox.Record(ctx, contract.Registered{
	//	IdentityID: ident.ID, Email: ident.Email, At: u.clock.Now(),
	//}); err != nil {
	//	return zero, err
	//}

	return contract.ProfileView{
		ID:       ident.ID,
		Email:    ident.Email,
		FullName: ident.FullName,
		Active:   ident.Active,
	}, nil
}
