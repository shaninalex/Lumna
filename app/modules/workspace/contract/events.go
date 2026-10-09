package contract

import "net/mail"

type InvitationCreated struct {
	Receiver           mail.Address
	Role               string
	InvitationLinkHash string
	InitiatorId        int
}

func (e InvitationCreated) EventName() string {
	return "invitation/created"
}
