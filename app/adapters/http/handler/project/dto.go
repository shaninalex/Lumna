package project

import "time"

type projectDTO struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	WorkspaceId int        `json:"workspace_id"`
	OwnerId     int        `json:"owner_id"`
	Meta        *string    `json:"meta,omitempty"`
	CreatedAt   time.Time  `json:"created_At"`
	UpdatedAt   *time.Time `json:"updated_At,omitempty"`
}
