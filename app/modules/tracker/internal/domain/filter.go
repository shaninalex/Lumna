package domain

// Filter - is a View, saved filter setup for work items list.
// I can't call it "view" because of architecture guidelines.
// Filter helps setup items more flexible than set of scopes and stages.
// Keeping Filter as a separate entity - you can even share filters.
type Filter struct {
	// TODO: entity is not implemented and not documented
}
