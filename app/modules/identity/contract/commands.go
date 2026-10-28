package contract

import (
	"context"

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

func ExecRegister(ctx context.Context, a *core.App, c Register) (ProfileView, error) {
	return bus.Execute[Register, ProfileView](ctx, a.Commands, c)
}

type PasswordChange struct {
	IdentityId int
	Password   secure.Secret
}

func (PasswordChange) Permission() (action string, scope int) { return "", 0 }

func ExecPasswordChange(ctx context.Context, a *core.App, c PasswordChange) (bool, error) {
	return bus.Execute[PasswordChange, bool](ctx, a.Commands, c)
}
