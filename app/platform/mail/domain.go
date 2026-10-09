package mail

import (
	"context"
	"net/mail"
	"time"
)

type EntryType string

var (
	InvitationEntryType EntryType = "invitation"
)

type Entry struct {
	ID        int
	Subject   string
	Content   string
	Type      EntryType
	Status    string
	Attempts  int16
	Receivers []mail.Address
	Message   string
	SendAt    time.Time
	CreatedAt time.Time
}

func (s *Entry) Retry(msg string) {
	s.Message = msg
	s.Attempts += 1
}

func (s *Entry) Failed(msg string) {
	s.Message = msg
	s.Status = "failed"
}

func (s *Entry) Sent(t time.Time) {
	s.SendAt = t
	s.Status = "sent"
	s.Message = ""
}

type Sender interface {
	ScheduleEmail(ctx context.Context, entry Entry)
}
