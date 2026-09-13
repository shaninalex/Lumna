package board

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type boardDTO struct {
	Id        int        `json:"id"`
	Title     string     `json:"title"`
	ProjectId int        `json:"project_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type boardCreateDTO struct {
	Title     string `json:"title"`
	ProjectId int    `json:"project_id"`
}

func (a boardCreateDTO) Validate() error {
	return validation.ValidateStruct(
		&a,
		validation.Field(&a.Title, validation.Required, validation.Length(3, 50)),
		validation.Field(&a.ProjectId, validation.Required),
	)
}
