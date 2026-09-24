package auth

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type passwordLoginDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (a passwordLoginDTO) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Email, validation.Required, is.Email),
		validation.Field(&a.Password, validation.Required),
	)
}
