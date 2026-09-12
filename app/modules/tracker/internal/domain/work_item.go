package domain

import "time"

type WorkItemType string

var (
	WorkItemTypeEpic    WorkItemType = "epic"
	WorkItemTypeStory   WorkItemType = "story"
	WorkItemTypeTask    WorkItemType = "task"
	WorkItemTypeBug     WorkItemType = "bug"
	WorkItemTypeSubtask WorkItemType = "subtask"
)

type Priority string

var (
	PriorityHigh    Priority = "high"
	PriorityWarning Priority = "warning"
	PriorityNormal  Priority = "normal"
)

type EstimateType string

const (
	EstimateTypeStoryPoints EstimateType = "story_points" // Scrum points (1, 2, 3, 5, 8...)
	EstimateTypeMinutes     EstimateType = "minutes"      // time in minutes
	EstimateTypeTShirt      EstimateType = "t_shirt"      // XS, S, M, L, XL
)

type Estimate struct {
	Value int          `json:"value"`
	Type  EstimateType `json:"type"`
}

// WorkItem - is the core entity. Main idea is instead of defining every entity type as a separate table or domain type,
// define single polymorphic type for all of them. This is how Linear, Gitlab Work Items and others do.
type WorkItem struct {
	ID        int
	ProjectID int
	Type      WorkItemType

	ParentID    *int
	Title       string
	Description string

	// Work context and execution stage:
	ScopeID int
	StageID int

	// Rank - Position within the stage. Using float64 solves the issue of mass updates.
	Rank float64

	Priority    Priority
	AssigneeIDs []int
	SprintID    *int
	Estimate    *Estimate

	CreatedAt time.Time
	UpdatedAt time.Time
}
