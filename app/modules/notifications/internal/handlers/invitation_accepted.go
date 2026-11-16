package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/notifications/internal/domain"
	workspaceContract "gitlab.com/shaninalex/lumna/app/modules/workspace/contract"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

type InvitationAccepted struct {
	clk     clock.Clock
	deliver *Deliverer
}

func NewInvitationAccepted(clk clock.Clock, deliver *Deliverer) *InvitationAccepted {
	return &InvitationAccepted{clk: clk, deliver: deliver}
}

func (s *InvitationAccepted) Handle(ctx context.Context, e workspaceContract.InvitationAccepted) error {
	data, err := content(map[string]any{"invitation_id": e.InvitationId})
	if err != nil {
		return err
	}
	n := domain.Notification{
		NotificationType: e.EventName(),
		Content:          data,
		RefId:            e.WorkspaceId,
		RefType:          "workspace",
		Created:          s.clk.Now(),
	}
	n.AddRecipients(0, e.RecipientsIds...)
	return s.deliver.Deliver(ctx, n)
}
