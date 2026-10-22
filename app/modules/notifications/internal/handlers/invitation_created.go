package handlers

import (
	"context"
	"fmt"
	"net/mail"

	"gitlab.com/shaninalex/lumna/app/modules/notifications/internal/domain"
	"gitlab.com/shaninalex/lumna/app/modules/notifications/internal/infra/email_templates"
	workspaceContract "gitlab.com/shaninalex/lumna/app/modules/workspace/contract"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

type InvitationCreated struct {
	repo        domain.EmailRepo
	clock       clock.Clock
	redirectUrl string
}

func NewInvitationCreated(
	repo domain.EmailRepo,
	clock clock.Clock,
	url string,
) *InvitationCreated {
	return &InvitationCreated{
		repo:        repo,
		clock:       clock,
		redirectUrl: url,
	}
}

func (s *InvitationCreated) Handle(ctx context.Context, event workspaceContract.InvitationCreated) error {
	url := fmt.Sprintf("%s%s%s", s.redirectUrl, "/confirm-invitation?key=", event.InvitationLinkHash)
	emailBody := email_templates.NewEntryInvitation(url)
	html, text := emailBody.GenerateTemplates(ctx)
	email := domain.Email{
		Subject:   "Invitation in Lumna workspace!",
		Html:      string(html),
		Text:      string(text),
		Type:      "invitation",
		Status:    domain.EmailStatusPending,
		Receivers: []mail.Address{event.Receiver},
		CreatedAt: s.clock.Now(),
	}
	return s.repo.Create(ctx, email)
}
