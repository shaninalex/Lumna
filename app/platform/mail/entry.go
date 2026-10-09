package mail

import (
	"context"
	"net/mail"

	"gitlab.com/shaninalex/lumna/app/platform/config"
)

type Message struct {
	To          []mail.Address
	Subject     string
	HTML, Text  []byte
	Attachments []Attachment
}

type Attachment struct {
	Id int
}

type Transport interface {
	Send(ctx context.Context, m Message) error // SMTP
}

// Queue - email_queue + worker/retries
type Queue interface {
	Enqueue(ctx context.Context, m Message) error
}

type EntryContentBuilder interface {
	GenerateTemplates(ctx context.Context) ([]byte, []byte)
}

func MakeEntry(e *Entry, cfg config.EmailerConfig) (EntryContentBuilder, error) {
	switch e.Type {
	case InvitationEntryType:
		return NewEntryInvitation(e, cfg.ServerUrl), nil
	default:
		return nil, EntryTypeUndefinedError
	}
}
