package contract

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/bus"
)

type Register struct {
	Email    string
	Password bus.Secret
	FullName string
}

func (Register) Permission() (action string, scope int) { return "", 0 }

func ExecRegister(ctx context.Context, a *core.App, c Register) (ProfileView, error) {
	return bus.Execute[Register, ProfileView](ctx, a.Commands, c)
}

type Authenticate struct {
	Email    string
	Password bus.Secret
}

func (Authenticate) Permission() (string, int) { return "", 0 }

func ExecAuthenticate(ctx context.Context, a *core.App, c Authenticate) (SessionView, error) {
	return bus.Execute[Authenticate, SessionView](ctx, a.Commands, c)
}
