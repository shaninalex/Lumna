package auth

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"gitlab.com/shaninalex/lumna/app/core/errs"
)

type passwordLoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a passwordLoginDTO) Verify() error {
	err := validation.ValidateStruct(&a,
		validation.Field(&a.Email, validation.Required, is.Email),
		validation.Field(&a.Password, validation.Required),
	)
	if err != nil {
		return errs.Unauthenticated("AUTH001", "invalid credentials", err)
	}

	return nil
}
