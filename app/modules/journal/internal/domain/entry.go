package domain

import "time"

// Entry is a single recorded fact about an entity.
type Entry struct {
	ID         int
	ActorID    int
	EntityType string
	EntityID   int
	EventType  string
	Data       string
	CreatedAt  time.Time
}
