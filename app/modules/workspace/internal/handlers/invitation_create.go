package handlers

import (
	"context"
	"net/mail"
	"time"

	"gitlab.com/shaninalex/lumna/app/core/errs"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/contract"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
	"gitlab.com/shaninalex/lumna/app/platform/email"
)

type CreateInvitation struct {
	repo      domain.InvitationRepo
	hasher    domain.TokenHasher
	clock     clock.Clock
	emlSender email.Sender
}

func NewCreateInvitation(repo domain.InvitationRepo, hasher domain.TokenHasher, clk clock.Clock, emlSender email.Sender) *CreateInvitation {
	return &CreateInvitation{
		repo:      repo,
		hasher:    hasher,
		clock:     clk,
		emlSender: emlSender,
	}
}

var CreateInvitationAlreadyExistsError = errs.Conflict("WSP01", "user already invited in this workspace")
var CreateInvitationAlreadyAcceptedError = errs.Conflict("WSP02", "user already accepted in this workspace")
var CreateInvitationAlreadyRevokedError = errs.Conflict("WSP03", "invitation was revoked for this user")

func (s *CreateInvitation) Handle(ctx context.Context, cmd contract.CreateInvitation) (contract.InvitationView, error) {
	_, hash, err := s.hasher.CreateToken()
	if err != nil {
		return contract.InvitationView{}, err
	}
	invitations, err := s.repo.ListByWorkspaceId(ctx, cmd.Invitation.WorkspaceID)
	if err != nil {
		return contract.InvitationView{}, err
	}

	for _, inv := range invitations {
		if inv.Email == cmd.Invitation.Email {
			if !inv.AcceptedAt.IsZero() {
				return contract.InvitationView{}, CreateInvitationAlreadyAcceptedError
			} else if !inv.RevokedAt.IsZero() {
				return contract.InvitationView{}, CreateInvitationAlreadyRevokedError
			}

			return contract.InvitationView{}, CreateInvitationAlreadyExistsError
		}
	}

	domainInvitation := toInvitationDomain(cmd.Invitation)
	domainInvitation.TokenHash = hash
	domainInvitation.ExpiresAt = s.clock.Now().Add(24 * time.Hour)

	inv, err := s.repo.Create(ctx, domainInvitation)
	if err != nil {
		return contract.InvitationView{}, err
	}

	// TODO: use constructor
	eml := email.Entry{
		Subject:   "invitation in workspace",
		Content:   "Email content with invitation link",
		Type:      "invitation",
		Receivers: []mail.Address{{Address: domainInvitation.Email}},
	}
	s.emlSender.ScheduleEmail(ctx, eml)

	return toInvitationView(inv), nil
}
