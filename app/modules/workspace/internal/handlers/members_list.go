package handlers

import (
	"context"

	identityContract "gitlab.com/shaninalex/lumna/app/modules/identity/contract"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/contract"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/internal/domain"
)

type MembersList struct {
	identityWorkspaceRepo domain.IdentityWorkspaceRepo
	profiler              identityContract.Profiler
}

func NewMembersList(identityWorkspaceRepo domain.IdentityWorkspaceRepo, profiler identityContract.Profiler) *MembersList {
	return &MembersList{
		identityWorkspaceRepo: identityWorkspaceRepo,
		profiler:              profiler,
	}
}

func (s *MembersList) Handle(ctx context.Context, cmd contract.MembersList) ([]contract.MemberView, error) {
	iws, err := s.identityWorkspaceRepo.ByWorkspaceId(ctx, cmd.WorkspaceId)
	if err != nil {
		return []contract.MemberView{}, err
	}

	ids := make([]int, len(iws))
	for i, iw := range iws {
		ids[i] = iw.IdentityID
	}

	profiles, err := s.profiler.ByIDs(ctx, ids)
	if err != nil {
		return []contract.MemberView{}, err
	}

	members := make([]contract.MemberView, len(iws))
	for i, iw := range iws {
		member := contract.MemberView{
			ID:          iw.IdentityID,
			Role:        iw.Role,
			DateJoined:  iw.CreatedAt,
			WorkspaceId: iw.WorkspaceID,
		}

		for _, profile := range profiles {
			if profile.ID == iw.IdentityID {
				member.Email = profile.Email
				member.FullName = profile.FullName
			}
		}

		members[i] = member
	}

	return members, nil
}
