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

func Send(
	ctx context.Context,
	cfg config.EmailerConfig,
	from, subject, html, txt string,
	receivers []mail.Address,
) error {
	m := elib.NewEmail()
	m.From = from
	m.Subject = subject
	m.To = emailsToSlice(receivers)
	m.HTML = []byte(html)
	m.Text = []byte(txt)

	smtpHost := cfg.Host
	smtpPort := cfg.Port
	smtpUser := cfg.User
	smtpPass := cfg.Password

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
