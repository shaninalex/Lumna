package contract

import "context"

type Pusher interface {
	Push(ctx context.Context, identityIDs []int, v NotificationView)
}

type EmailDispatcher interface {
	// Run sends pending emails until ctx is cancelled. Blocks.
	Run(ctx context.Context)
}
