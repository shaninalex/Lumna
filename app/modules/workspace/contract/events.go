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

type InvitationAccepted struct {
	InvitationId  int
	WorkspaceId   int
	RecipientsIds []int
}

func (e InvitationAccepted) EventName() string {
	return "invitation/accepted"
}
