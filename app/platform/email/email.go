package email

type Email struct {
	EmailType string
	Subject   string
	Content   string
}

func NewEmail(emailType string, subject string, content string) *Email {
	return &Email{
		EmailType: emailType,
		Subject:   subject,
		Content:   content,
	}
}

func InviteEmail(subject string, content string) *Email {
	return NewEmail("invite", subject, content)
}
