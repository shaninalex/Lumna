package actor

import "context"

type Actor struct {
	IdentityID  int
	WorkspaceID int
	Roles       []string
	IsSystem    bool // CLI, migrations, outbox-relay
}

func System() Actor { return Actor{IsSystem: true} }

type actorKey struct{}

func With(ctx context.Context, a Actor) context.Context {
	return context.WithValue(ctx, actorKey{}, a)
}

func From(ctx context.Context) (Actor, bool) {
	a, ok := ctx.Value(actorKey{}).(Actor)
	return a, ok
}
