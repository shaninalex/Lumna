package project

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type projectDTO struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	WorkspaceId int        `json:"workspace_id"`
	OwnerId     int        `json:"owner_id"`
	Meta        *string    `json:"meta,omitempty"`
	CreatedAt   time.Time  `json:"created_At"`
	UpdatedAt   *time.Time `json:"updated_At,omitempty"`
}

type projectCreateDTO struct {
	Title       string `json:"title"`
	WorkspaceId int    `json:"workspace_id"`
}

func (a projectCreateDTO) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Title, validation.Required, validation.Length(3, 50)),
		validation.Field(&a.WorkspaceId, validation.Required),
	)
}
