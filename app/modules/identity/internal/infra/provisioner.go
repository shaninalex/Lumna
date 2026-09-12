package infra

import (
	"context"
	"errors"

	"gitlab.com/shaninalex/lumna/app/modules/identity/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

// Provisioner implements contract.Provisioner: the port another module uses
// when it has an email from an external provider and needs an identity for
// it, existing or new.
//
// It writes through the repo, which takes the transaction from ctx
// (database.DB.From). So when auth calls this from inside its own command,
// the new identity and auth's own rows commit or roll back together — without
// either module knowing about the other's transaction.
type Provisioner struct {
	identities domain.IdentityRepo
	clock      clock.Clock
}

func NewProvisioner(i domain.IdentityRepo, clk clock.Clock) *Provisioner {
	return &Provisioner{identities: i, clock: clk}
}

// EnsureIdentityByEmail returns the id of an existing identity, or creates one.
func (p *Provisioner) EnsureIdentityByEmail(ctx context.Context, email, fullName string) (int, error) {
	ident, err := p.identities.ByEmail(ctx, email)
	if err == nil {
		return ident.ID, nil
	}
	if !errors.Is(err, domain.ErrNotFound) {
		return 0, err
	}

	ident, err = domain.NewIdentity(email, fullName, p.clock.Now())
	if err != nil {
		return 0, err
	}
	if err := p.identities.Save(ctx, ident); err != nil {
		return 0, err
	}

	return ident.ID, nil
}
