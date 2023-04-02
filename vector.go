// Package vector provides a generic wrapper for Go slices and includes some helper methods
// similar to that of a C++ Vector<T>
package vector

import "errors"

// Vector wraps a Go slice of any supported type
type Vector[T any] struct {
	Slice []T
}

// New returns a new Vector with a slice of type T initialized to the provided length and capacity
func New[T any](length, capacity int) Vector[T] {
	return Vector[T]{make([]T, length, capacity)}
}

// Push adds item(s) to the end of the slice, abstracting away Go's interesting
// slice = append(slice, data) pattern
func (v *Vector[T]) Push(items ...T) {
	v.Slice = append(v.Slice, items...)
}

// Pop removes the last item of the slice and returns it
func (v *Vector[T]) Pop() (T, error) {
	length := len(v.Slice)
	if length > 0 {
		popped := v.Slice[length-1]
		v.Slice = v.Slice[:length-1]
		return popped, nil
	}
	// If the slice is empty, return the 0 value of the generic type and the error
	var t T
	return t, errors.New("slice is empty")
}

// Size returns the number of elements in the underlying slice
func (v *Vector[T]) Size() int {
	return len(v.Slice)
}

// Capacity returns the capacity of the underlying slice
func (v *Vector[T]) Capacity() int {
	return cap(v.Slice)
}

// Shrink reduces the capacity to match the current length
func (v *Vector[T]) Shrink() {
	if cap(v.Slice) > len(v.Slice) {
		shrunk := make([]T, len(v.Slice))
		copy(shrunk, v.Slice)
		v.Slice = shrunk
	}
}
