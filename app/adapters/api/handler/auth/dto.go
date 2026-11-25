package auth

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type passwordLoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a passwordLoginDTO) Verify() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Email, validation.Required, is.Email),
		validation.Field(&a.Password, validation.Required),
	)
}
