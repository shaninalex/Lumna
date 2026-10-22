package infra

import (
	"context"
	"fmt"
	"log/slog"
	"runtime/debug"
	"time"

	"gitlab.com/shaninalex/lumna/app/modules/notifications/contract"
	"gitlab.com/shaninalex/lumna/app/modules/notifications/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
	"gitlab.com/shaninalex/lumna/app/platform/mail"
)

var (
	ProcessLoopDuration   = 15 * time.Second
	MaxSendAttemptsAmount = int16(3)
)

type EmailDispatcher struct {
	repo   domain.EmailRepo
	log    *slog.Logger
	sender mail.Sender
	clock  clock.Clock
}

func NewEmailDispatcher(
	repo domain.EmailRepo,
	log *slog.Logger,
	sender mail.Sender,
	clock clock.Clock,
) contract.EmailDispatcher {
	return &EmailDispatcher{
		repo:   repo,
		log:    log,
		sender: sender,
		clock:  clock,
	}
}

var _ contract.EmailDispatcher = (*EmailDispatcher)(nil)

func (s *EmailDispatcher) Run(ctx context.Context) {
	ticker := time.NewTicker(ProcessLoopDuration)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			s.tick(ctx)
		case <-ctx.Done():
			return
		}
	}
}

func (s *EmailDispatcher) tick(ctx context.Context) {
	defer func() {
		if r := recover(); r != nil {
			s.log.Error(fmt.Sprintf("[EmailDispatcher]: tick panic: %v", r), "stack", string(debug.Stack()))
		}
	}()
	entries, err := s.repo.PendingEmails(ctx)
	if err != nil {
		s.log.Error(fmt.Sprintf("[EmailDispatcher]: unable to load pending emails: %s", err.Error()))
		return
	}
	s.log.Info(fmt.Sprintf("[EmailDispatcher]: %d emails are going to send", len(entries)))

	// An entry that started sending must also get its status saved, so it is
	// processed with a context that shutdown does not cancel. Shutdown only
	// stops us from picking up the next entry.
	work := context.WithoutCancel(ctx)
	for _, entry := range entries {
		if ctx.Err() != nil {
			return
		}
		s.processEntry(work, entry)
	}
}

func (s *EmailDispatcher) processEntry(ctx context.Context, entry domain.Email) {
	defer func() {
		if r := recover(); r != nil {
			msg := fmt.Sprintf("panic while sending email: %v", r)
			s.log.Error(msg)

			entry.Failed(msg)
			if err := s.repo.Update(ctx, entry); err != nil {
				s.log.Error(err.Error())
			}
		}
	}()

	if err := s.sender.Send(ctx, entry.Subject, entry.Html, entry.Text, entry.Receivers); err != nil {
		msg := err.Error()
		s.log.Error(msg)
		if entry.Attempts == MaxSendAttemptsAmount {
			entry.Failed(msg)
		} else {
			entry.Retry(msg)
		}

		if err = s.repo.Update(ctx, entry); err != nil {
			s.log.Error(err.Error())
		}
		return
	}

	entry.Sent(s.clock.Now())
	if err := s.repo.Update(ctx, entry); err != nil {
		s.log.Error(err.Error())
	}
	s.log.Info(fmt.Sprintf("[EmailDispatcher]: sent. Type=%s receivers=%v", entry.Type, entry.Receivers))
}
