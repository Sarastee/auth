package common

// Pointer ...
func Pointer[T any](element T) *T {
	return &element
}
