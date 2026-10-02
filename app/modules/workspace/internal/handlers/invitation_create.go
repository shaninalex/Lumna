package handlers

import (
	"context"
	"fmt"

	"gitlab.com/shaninalex/lumna/app/modules/workspace/contract"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

type CreateInvitation struct {
	repo   domain.InvitationRepo
	hasher domain.TokenHasher
	clock  clock.Clock
}

func NewCreateInvitation(repo domain.InvitationRepo, hasher domain.TokenHasher, clk clock.Clock) *CreateInvitation {
	return &CreateInvitation{
		repo:   repo,
		hasher: hasher,
		clock:  clk,
	}
}

func (s *CreateInvitation) Handle(ctx context.Context, cmd contract.CreateInvitation) (contract.InvitationView, error) {
	token, hash, err := s.hasher.CreateToken()
	if err != nil {
		return contract.InvitationView{}, err
	}
	cmd.Invitation.TokenHash = hash
	inv, err := s.repo.Save(ctx, toInvitationDomain(cmd.Invitation))
	if err != nil {
		return contract.InvitationView{}, err
	}

	fmt.Printf("token will be send in email: %s", token)
	return toInvitationView(inv), nil
}
