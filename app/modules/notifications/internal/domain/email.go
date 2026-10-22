package domain

import (
	"net/mail"
	"time"

	"gitlab.com/shaninalex/lumna/app/core/errs"
)

type EntryType string

var (
	InvitationEntryType EntryType = "invitation"
)

type EmailStatus string

var (
	EmailStatusPending EmailStatus = "pending"
	EmailStatusSending EmailStatus = "sending"
	EmailStatusSent    EmailStatus = "sent"
	EmailStatusError   EmailStatus = "error"
)

type Email struct {
	ID        int
	Subject   string
	Html      string
	Text      string
	Type      EntryType
	Status    EmailStatus
	Attempts  int16
	Receivers []mail.Address
	Message   string
	SendAt    time.Time
	CreatedAt time.Time
}

func (s *Email) Retry(msg string) {
	s.Message = msg
	s.Attempts += 1
}

func (s *Email) Failed(msg string) {
	s.Message = msg
	s.Status = "failed"
}

func (s *Email) Sent(t time.Time) {
	s.SendAt = t
	s.Status = "sent"
	s.Message = ""
}

var (
	EntryWasNotFoundError      = errs.NotFound("EML00", "email entry was not found")
	EntryWasNotUpdatedError    = errs.Platform("EML01", "email entry was not updated")
	ProcessorUnableToTickError = errs.Platform("EML02", "email processor unable to process tick entries batch")
	EntryTypeUndefinedError    = errs.Platform("EML04", "email type is undefined")
)
