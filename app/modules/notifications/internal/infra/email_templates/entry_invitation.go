package email_templates

import (
	"bytes"
	"context"
	_ "embed"
	"html/template"

	"github.com/vanng822/go-premailer/premailer"
)

//go:embed entry_invitation.txt
var entryInvitationTxt string

type EntryInvitation struct {
	invitationUrl string
	txtTemplate   string
}

func (s *EntryInvitation) GenerateTemplates(ctx context.Context) ([]byte, []byte) {
	return s.generateHTML(ctx), s.generateTXT()
}

func NewEntryInvitation(invitationUrl string) *EntryInvitation {
	return &EntryInvitation{
		invitationUrl: invitationUrl,
		txtTemplate:   entryInvitationTxt,
	}
}

func (s *EntryInvitation) generateHTML(ctx context.Context) []byte {
	var buffer bytes.Buffer
	if err := s.template().Render(ctx, &buffer); err != nil {
		panic(err)
	}

	pre, err := premailer.NewPremailerFromBytes(buffer.Bytes(), premailer.NewOptions())
	if err != nil {
		panic(err)
	}

	html, err := pre.Transform()
	if err != nil {
		panic(err)
	}

	return []byte(html)
}

func (s *EntryInvitation) generateTXT() []byte {
	var buffer bytes.Buffer
	t := template.Must(template.New("invitation_email").Parse(s.txtTemplate))
	err := t.Execute(&buffer, map[string]any{"invitationLink": s.invitationUrl})
	if err != nil {
		panic(err)
	}
	return buffer.Bytes()
}
