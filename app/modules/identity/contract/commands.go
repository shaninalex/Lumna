package contract

import (
	"context"
	"net/mail"

	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/bus"
	"gitlab.com/shaninalex/lumna/app/core/secure"
)

type Register struct {
	Email    string
	Password secure.Secret
	FullName string
}

func (Register) Permission() (action string, scope int) { return "", 0 }

func ExecRegister(ctx context.Context, a *core.App, cmd Register) (ProfileView, error) {
	return bus.Execute[Register, ProfileView](ctx, a.Commands, cmd)
}

type PasswordChange struct {
	IdentityId int
	Password   secure.Secret
}

func (PasswordChange) Permission() (action string, scope int) { return "", 0 }

func ExecPasswordChange(ctx context.Context, a *core.App, cmd PasswordChange) (bool, error) {
	return bus.Execute[PasswordChange, bool](ctx, a.Commands, cmd)
}

type AccountChange struct {
	IdentityID int
	Email      *mail.Address
	FullName   string
}

func (AccountChange) Permission() (action string, scope int) { return "", 0 }

func ExecAccountChange(ctx context.Context, a *core.App, cmd AccountChange) (ProfileView, error) {
	return bus.Execute[AccountChange, ProfileView](ctx, a.Commands, cmd)
}
