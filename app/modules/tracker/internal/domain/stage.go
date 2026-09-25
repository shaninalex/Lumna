package domain

import "time"

// StageCategory helps classify stages for analytics, filters, or the UI (e.g., to determine when a task is "completed").
type StageCategory string

const (
	StageCategoryBacklog    StageCategory = "backlog"
	StageCategoryUnstarted  StageCategory = "unstarted"
	StageCategoryInProgress StageCategory = "in_progress"
	StageCategoryCompleted  StageCategory = "completed"
	StageCategoryCanceled   StageCategory = "canceled"
)

// Stage represents a specific step or phase in the work lifecycle within a Scope.
type Stage struct {
	ID          int
	ScopeID     int
	Name        string
	Description string
	Category    StageCategory

	// Position - within the workflow. Using float64 solves the issue of mass updates.
	Position float64

	// Work In Progress limit
	WIPLimit int

	WorkItems []WorkItem

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *Stage) UpdatePosition(position float64, timestamp time.Time) {
	s.Position = position
	s.UpdatedAt = timestamp
}
