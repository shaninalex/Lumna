package domain

import "time"

// Scope is a working context or space, groups together related stages and tasks.
type Scope struct {
	ID          int
	ProjectID   int
	Name        string
	Description string

	Stages []Stage

	CreatedAt time.Time
	UpdatedAt time.Time
}
