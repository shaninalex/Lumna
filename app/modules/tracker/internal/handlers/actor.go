package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/core/actor"
)

// actorID returns the identity that owns the request, 0 if there is none.
// Handlers resolve it once and put it into events, so subscribers never
// read the actor from ctx.
func actorID(ctx context.Context) int {
	if a, ok := actor.From(ctx); ok {
		return a.IdentityID
	}
	return 0
}
