package domain

// Projection - is a View, saved filter setup for work items list.
// Projection helps setup items more flexible than set of scopes and stages.
// Keeping Projection as a separate entity - you can even share filters.
type Projection struct {
	ID int
}
