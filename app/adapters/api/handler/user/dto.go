package user

import (
	"net/mail"
)

type profileDTO struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Active   bool   `json:"active"`
}

type changePasswordDTO struct {
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

func (s changePasswordDTO) Verify() error {
	if s.Password != s.PasswordConfirm {
		return ErrPasswordsDoesNotMatch
	}

	return nil
}

type accountChangeDTO struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

func (s accountChangeDTO) Verify() error {
	if _, err := mail.ParseAddress(s.Email); err != nil {
		return err
	}

	return nil
}
