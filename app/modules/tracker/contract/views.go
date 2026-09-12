package contract

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

// WorkItem - is the core entity. Main idea is instead of defining every entity type as a separate table or domain type,
// define single polymorphic type for all of them. This is how Linear, Gitlab Work Items and others do.
type WorkItem struct {
	ID        int
	ProjectID int
	Type      WorkItemType

	// ParentID allows hierarchy structure of different entities:
	// Epic -> Story/Task, Task -> Subtask
	ParentID    *int
	Title       string
	Description string

	// StatusID - reference on Workflow state (Todo, In Progress, Done)
	StatusID    int
	Priority    Priority
	AssigneeIDs []int
	SprintID    *int
	Estimate    *int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
