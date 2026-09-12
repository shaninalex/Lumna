package domain

import "time"

type Project struct {
	ID          int
	Title       string
	WorkspaceId int
	OwnerId     int
	Meta        *string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}

func NewProject(title string, workspaceId, ownerId int) *Project {
	return &Project{
		Title:       title,
		WorkspaceId: workspaceId,
		OwnerId:     ownerId,
		CreatedAt:   time.Now(),
	}
}
