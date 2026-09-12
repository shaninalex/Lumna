package domain

import "time"

type Workspace struct {
	ID         int
	Title      string
	Active     bool
	OwnerEmail string
	CreatedAt  time.Time
	UpdatedAt  *time.Time
}

func NewWorkspace(title, email string, active bool, now time.Time) *Workspace {
	return &Workspace{
		Title:      title,
		OwnerEmail: email,
		Active:     active,
		CreatedAt:  now,
	}
}
