package infra

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/identity/contract"
	"gitlab.com/shaninalex/lumna/app/modules/identity/internal/domain"
)

type Profiler struct {
	repo domain.IdentityRepo
}

func NewProfiler(repo domain.IdentityRepo) *Profiler {
	return &Profiler{repo: repo}
}

func (s *Profiler) ByID(ctx context.Context, id int) (contract.ProfileView, error) {
	identity, err := s.repo.ByID(ctx, id)
	if err != nil {
		return contract.ProfileView{}, err
	}
	return contract.ProfileView{
		ID:       identity.ID,
		Email:    identity.Email,
		FullName: identity.FullName,
		Active:   identity.Active,
	}, nil
}

func (s *Profiler) ByIDs(ctx context.Context, ids []int) ([]contract.ProfileView, error) {
	identity, err := s.repo.ListByIDs(ctx, ids)
	if err != nil {
		return []contract.ProfileView{}, err
	}

	result := make([]contract.ProfileView, 0, len(identity))
	for _, idn := range identity {
		result = append(result, contract.ProfileView{
			ID:       idn.ID,
			Email:    idn.Email,
			FullName: idn.FullName,
			Active:   idn.Active,
		})
	}
	return result, nil
}
