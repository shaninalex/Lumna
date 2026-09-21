package handlers

import (
	"gitlab.com/shaninalex/lumna/app/modules/identity/contract"
	"gitlab.com/shaninalex/lumna/app/modules/identity/internal/domain"
)

func toProfileView(i domain.Identity) contract.ProfileView {
	return contract.ProfileView{
		ID:       i.ID,
		Email:    i.Email,
		FullName: i.FullName,
		Active:   i.Active,
	}
}
