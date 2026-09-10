package ptr

// Pointer return pointer of a provided value
func P[T any](v T) *T {
	return &v
}
