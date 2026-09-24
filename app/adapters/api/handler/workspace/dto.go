package workspace

import (
	"time"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type workspaceDTO struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Active     bool      `json:"active"`
	OwnerEmail string    `json:"owner_email"`
	CreatedAt  time.Time `json:"created_at"`
}

type workspaceCreateDTO struct {
	Title string `json:"title" binding:"required"`
	Email string `json:"email" binding:"required"`
}

func (a workspaceCreateDTO) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Title, validation.Required, validation.Length(3, 50)),
		validation.Field(&a.Email, validation.Required, is.Email),
	)
}
