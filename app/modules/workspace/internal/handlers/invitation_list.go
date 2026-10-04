package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/workspace/contract"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/internal/domain"
)

type InvitationList struct {
	repo domain.InvitationRepo
}

func NewInvitationList(repo domain.InvitationRepo) *InvitationList {
	return &InvitationList{repo: repo}
}

func (s *InvitationList) Handle(ctx context.Context, query contract.InvitationList) ([]contract.InvitationView, error) {
	invitationsDomains, err := s.repo.ListByWorkspaceId(ctx, query.WorkspaceId)
	if err != nil {
		return []contract.InvitationView{}, err
	}
	invitations := make([]contract.InvitationView, len(invitationsDomains))
	for i, invitation := range invitationsDomains {
		invitations[i] = toInvitationView(invitation)
	}
	return invitations, nil
}
