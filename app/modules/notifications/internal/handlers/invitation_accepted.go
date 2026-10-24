package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/notifications/contract"
	"gitlab.com/shaninalex/lumna/app/modules/notifications/internal/domain"
	workspaceContract "gitlab.com/shaninalex/lumna/app/modules/workspace/contract"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

type InvitationAccepted struct {
	pusher contract.Pusher
	clk    clock.Clock
	repo   domain.NotificationRepo
}

func NewInvitationAccepted(
	pusher contract.Pusher,
	clk clock.Clock,
	repo domain.NotificationRepo,
) *InvitationAccepted {
	return &InvitationAccepted{pusher, clk, repo}
}

func (s *InvitationAccepted) Handle(ctx context.Context, event workspaceContract.InvitationAccepted) error {
	n := domain.Notification{
		NotificationType: event.EventName(),
		Content:          "Invitation accepted",
		Created:          s.clk.Now(),
	}

	n, err := s.repo.Save(ctx, n)
	if err != nil {
		return err
	}
	s.pusher.Push(ctx, event.RecipientsIds, contract.NotificationView{
		ID:        n.ID,
		Type:      event.EventName(),
		Content:   n.Content,
		CreatedAt: n.Created,
	})

	return nil
}
