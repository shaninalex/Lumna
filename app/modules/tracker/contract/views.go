package contract

import "time"

type ScopeView struct {
	Id          int
	Name        string
	Description string
	ProjectId   int
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}

type StageView struct {
	Id          int
	ScopeId     int
	Name        string
	Description string
	Category    string
	Position    float64
	WIPLimit    *int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
