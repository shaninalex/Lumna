package actor

import "context"

type Actor struct {
	IdentityID  int
	WorkspaceID int
	Roles       []string
	IsSystem    bool // CLI, migrations, outbox-relay
}

func With(ctx context.Context, a Actor) context.Context
func From(ctx context.Context) (Actor, bool)
