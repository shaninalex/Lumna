package handlers

import (
	"context"
	"log/slog"

	"gitlab.com/shaninalex/lumna/app/core/bus"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/contract"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

type AcceptInvitation struct {
	repoInvitation        domain.InvitationRepo
	repoIdentityWorkspace domain.IdentityWorkspaceRepo
	clk                   clock.Clock
	bus                   *bus.EventBus
	logger                *slog.Logger
}

func NewAcceptInvitation(
	repoInvitation domain.InvitationRepo,
	repoIdentityWorkspace domain.IdentityWorkspaceRepo,
	clk clock.Clock,
	bus *bus.EventBus,
	logger *slog.Logger,
) *AcceptInvitation {
	return &AcceptInvitation{repoInvitation, repoIdentityWorkspace, clk, bus, logger}
}

func (s *AcceptInvitation) Handle(ctx context.Context, cmd contract.AcceptInvitation) (contract.InvitationView, error) {
	invitation, err := s.repoInvitation.GetById(ctx, cmd.InvitationId)
	if err != nil {
		return contract.InvitationView{}, domain.InvitationNotFoundError
	}

	if !invitation.AcceptedAt.IsZero() {
		return contract.InvitationView{}, domain.CreateInvitationAlreadyAcceptedError
	}

	_, err = s.repoIdentityWorkspace.Save(ctx, domain.IdentityWorkspace{
		IdentityID:  cmd.IdentityId,
		WorkspaceID: invitation.WorkspaceID,
		CreatedAt:   s.clk.Now(),
		Role:        invitation.Role,
	})
	if err != nil {
		return contract.InvitationView{}, domain.InvitationUnableToSaveError
	}

	invitation.Accept(s.clk.Now())

	err = s.repoInvitation.Update(ctx, invitation)
	if err != nil {
		return contract.InvitationView{}, domain.InvitationUnableToSaveError
	}

	ids, err := s.getWorkspaceAdministrators(ctx, invitation.WorkspaceID)
	if err != nil {
		return contract.InvitationView{}, err
	}

	if err = s.bus.Publish(ctx, contract.InvitationAccepted{
		InvitationId:  invitation.ID,
		WorkspaceId:   invitation.WorkspaceID,
		RecipientsIds: ids,
	}); err != nil {
		s.logger.Error(err.Error())
	}

	return toInvitationView(invitation), nil
}

func (s *AcceptInvitation) getWorkspaceAdministrators(ctx context.Context, wpId int) (ids []int, err error) {
	members, err := s.repoIdentityWorkspace.ByWorkspaceId(ctx, wpId)
	if err != nil {
		return nil, err
	}

	var admins []domain.IdentityWorkspace
	for _, m := range members {
		if m.Role == "admin" {
			admins = append(admins, m)
		}
	}

	for _, a := range admins {
		ids = append(ids, a.IdentityID)
	}

	return ids, nil
}
