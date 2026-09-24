package contract

import "context"

type Pusher interface {
	Push(ctx context.Context, identityIDs []int, v NotificationView)
}
