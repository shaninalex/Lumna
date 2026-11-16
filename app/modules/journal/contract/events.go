package contract

import "gitlab.com/shaninalex/lumna/app/core/bus"

// Recordable is an event the journal stores. Producer modules implement it on
// their own events; the journal subscribes to the interface, so it never
// imports producers' contracts. The whole event is stored as JSON data.
type Recordable interface {
	bus.Event
	// JournalSubject is the entity the event happened to, e.g. ("work_item", 42).
	JournalSubject() (entityType string, entityID int)
	// JournalActor is the identity that caused the event, 0 for system.
	JournalActor() int
}
