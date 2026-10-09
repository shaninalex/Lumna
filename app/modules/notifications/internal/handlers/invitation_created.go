package handlers

import (
	"context"

	workspaceContract "gitlab.com/shaninalex/lumna/app/modules/workspace/contract"
)

type InvitationCreated struct{}

func NewInvitationCreated() *InvitationCreated {
	return &InvitationCreated{}
}

func (s *InvitationCreated) Handle(ctx context.Context, event workspaceContract.InvitationCreated) error {
	// build and save email in db
	return nil
}
