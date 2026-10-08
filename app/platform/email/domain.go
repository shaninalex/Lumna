package email

import (
	"context"
	"time"
)

type Entry struct {
	ID        int
	Subject   string
	Content   string
	Type      string
	Status    string
	Attempts  int16
	Receivers []string // TODO: make proper validation/append/remove entry methods, or do emailString type properly
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
