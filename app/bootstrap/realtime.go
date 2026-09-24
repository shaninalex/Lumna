package bootstrap

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/notifications/contract"
	"gitlab.com/shaninalex/lumna/app/platform/database"
	"gitlab.com/shaninalex/lumna/app/platform/realtime"
)

type notificationPusher struct{ hub *realtime.Hub }

func (p notificationPusher) Push(ctx context.Context, ids []int, v contract.NotificationView) {
	database.AfterCommit(ctx, func() {
		p.hub.Send(ids, realtime.Envelope{Name: "notification.created", Payload: v})
	})
}
