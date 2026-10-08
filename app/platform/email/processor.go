package email

import (
	"context"
	"fmt"
	"log/slog"
	"runtime/debug"
	"time"

	"github.com/pkg/errors"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
	"gitlab.com/shaninalex/lumna/app/platform/config"
)

var (
	ProcessLoopDuration   = 15 * time.Second
	MaxSendAttemptsAmount = int16(3)
)

type Processor struct {
	cfg   *config.Config
	log   *slog.Logger
	clock clock.Clock
	repo  Repository
}

func NewSender(cfg *config.Config, log *slog.Logger, clock clock.Clock, repo Repository) *Processor {
	s := &Processor{cfg: cfg, log: log, clock: clock, repo: repo}
	return s
}

func (s *Processor) ScheduleEmail(ctx context.Context, entry Entry) {
	if err := s.repo.Create(ctx, entry); err != nil {
		s.log.Error(err.Error())
	}
}

func (s *Processor) Process(ctx context.Context) {
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

func (s *Processor) tick(ctx context.Context) {
	defer func() {
		if r := recover(); r != nil {
			s.log.Error(fmt.Sprintf("[Email Processor]: tick panic: %v", r), "stack", string(debug.Stack()))
		}
	}()

	entries, err := s.repo.PendingEmails(ctx)
	if err != nil {
		s.log.Error(errors.Wrap(ProcessorUnableToTickError, err.Error()).Error())
		return
	}

	s.log.Info(fmt.Sprintf("[Email Processor]: %d emails are going to send", len(entries)))
	for _, entry := range entries {
		s.processEntry(ctx, entry)
	}
}

func (s *Processor) processEntry(ctx context.Context, entry Entry) {
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

	if err := s.send(ctx, entry); err != nil {
		msg := err.Error()
		s.log.Error(msg)

		if entry.Attempts == MaxSendAttemptsAmount {
			entry.Retry(msg)
		} else {
			entry.Failed(msg)
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
}

func (s *Processor) send(ctx context.Context, entry Entry) error {
	fmt.Printf("Sending email: %s\n", entry.Type)
	return nil
}
