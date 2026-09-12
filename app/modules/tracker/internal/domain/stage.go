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
// Examples: "Referencing", "Planning", "In Review", "Deployed".
type Stage struct {
	ID          int
	ScopeID     int
	Name        string
	Description string
	Category    StageCategory

	// Position defines the stage's order in the workflow (e.g., 1 -> Research, 2 -> Progress, 3 -> Done).
	// float64 is used to facilitate reordering (Fractional Indexing / Lexorank).
	Position float64

	// Work In Progress limit
	WIPLimit *int

	CreatedAt time.Time
	UpdatedAt time.Time
}
