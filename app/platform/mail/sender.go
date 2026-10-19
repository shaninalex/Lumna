package mail

import (
	"context"
	"fmt"
	"net/mail"
	"net/smtp"

	elib "github.com/jordan-wright/email"
	"gitlab.com/shaninalex/lumna/app/core/errs"
	"gitlab.com/shaninalex/lumna/app/platform/config"
)

type Sender interface {
	Send(ctx context.Context, subject, html, txt string, receivers []mail.Address) error
}

type sender struct {
	cfg config.EmailerConfig
}

var _ Sender = (*sender)(nil)

func NewSender(cfg config.EmailerConfig) Sender {
	return &sender{cfg: cfg}
}

func (s *sender) Send(ctx context.Context, subject, html, txt string, receivers []mail.Address) error {
	m := elib.NewEmail()
	m.From = s.cfg.From
	m.Subject = subject
	m.To = emailsToSlice(receivers)
	m.HTML = []byte(html)
	m.Text = []byte(txt)

	smtpHost := s.cfg.Host
	smtpPort := s.cfg.Port
	smtpUser := s.cfg.User
	smtpPass := s.cfg.Password

	emailServerAddr := fmt.Sprintf("%s:%v", smtpHost, smtpPort)
	if err := m.Send(emailServerAddr, smtp.PlainAuth("", smtpUser, smtpPass.Reveal(), smtpHost)); err != nil {
		return errs.Platform(
			"EML03",
			fmt.Sprintf(
				"failed to send email: subject=%q, to=%v: %s",
				subject, receivers, err.Error(),
			),
		)
	}
	return nil
}

func emailsToSlice(emails []mail.Address) []string {
	var toS []string
	for _, value := range emails {
		if value.Address == "" {
			continue
		}
		toS = append(toS, value.String())
	}
	return toS
}
