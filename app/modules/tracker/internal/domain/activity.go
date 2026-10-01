package domain

import "time"

// Activity related to WorkItem, Scope, Stage
type Activity struct {
	ID           int
	IdentityID   int
	EntityID     int
	IdentityType int
	Content      string
	CreatedAt    time.Time
}
