package handlers

import (
	"context"
	"net/mail"

	"gitlab.com/shaninalex/lumna/app/modules/notifications/internal/domain"
	"gitlab.com/shaninalex/lumna/app/modules/notifications/internal/infra/email_templates"
	workspaceContract "gitlab.com/shaninalex/lumna/app/modules/workspace/contract"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

type InvitationCreated struct {
	repo  domain.EmailRepo
	clock clock.Clock
}

func NewInvitationCreated(
	repo domain.EmailRepo,
	clock clock.Clock,
) *InvitationCreated {
	return &InvitationCreated{
		repo:  repo,
		clock: clock,
	}
}

func (s *InvitationCreated) Handle(ctx context.Context, event workspaceContract.InvitationCreated) error {
	email_body := email_templates.NewEntryInvitation(event.InvitationLinkHash)
	html, _ := email_body.GenerateTemplates(ctx)
	email := domain.Email{
		Subject:   "Invitation in Lumna workspace!",
		Html:      string(html),
		Type:      "invitation",
		Status:    domain.EmailStatusPending,
		Receivers: []mail.Address{event.Receiver},
		CreatedAt: s.clock.Now(),
	}
	return s.repo.Create(ctx, email)
}
