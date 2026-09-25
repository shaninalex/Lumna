package contract

import "time"

type ScopeView struct {
	Id          int
	Name        string
	Description string
	ProjectId   int
	StageCount  int
	IssueCount  int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type StageView struct {
	Id          int
	ScopeId     int
	Name        string
	Description string
	Category    string
	Position    float64
	WIPLimit    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type StageDeleteView struct {
	StageId      int
	DeletedTasks []int
}

type WorkItemView struct {
	Id          int
	Title       string
	Description string
	ProjectId   int
	StageId     int
	ScopeId     int
	Rank        float64
	DueTo       time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Assignees   []int
}
