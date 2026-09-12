package domain

import "time"

// Scope is a working context or space (e.g., "Engineers Scope", "Marketing Scope", "Q3 Release Scope").
// It groups together related stages and tasks.
type Scope struct {
	ID          int
	ProjectID   int
	Name        string
	Description string

	// Stages contains the sequence of stages within this context
	Stages []Stage

	CreatedAt time.Time
	UpdatedAt time.Time
}
