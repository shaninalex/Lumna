package domain

import "time"

type Sprint struct {
	ID          int
	ProjectID   int
	Name        string
	Description string
	StartDate   time.Time
	EndDate     time.Time
	IsActive    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
