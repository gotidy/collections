package slice

import (
	"fmt"
)

// Index returns the index of the first instance of v in s, or -1 if v is not present in s.
func Index[T comparable](s []T, v T) int {
	for i, item := range s {
		if v == item {
			return i
		}
	}

	return -1
}

// Insert the value at the specified position of the slice.
func Insert[T any](s []T, pos int, v T) []T {
	if pos > len(s) || pos < 0 {
		panic(fmt.Sprintf("pos %d is out of range 0..%d", pos, len(s)))
	}
	len := len(s)
	if cap(s) > len {
		s = s[:len+1]
		if pos < len {
			copy(s[pos+1:], s[pos:len])
		}
		s[pos] = v
		return s
	}

	result := make([]T, len+1)
	if pos > 0 {
		copy(result[:pos], s[:pos])
	}
	if pos < len {
		copy(result[pos+1:], s[pos:])
	}
	result[pos] = v
	return result
}

// Reverse items of the slice.
func Reverse[T any](s []T) {
	for i, j := 0, len(s)-1; i < len(s)/2; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
}

// Map turns a []T1 to a []T2 using a mapping function.
// This function has two type parameters, T1 and T2.
// This works with slices of any type.
func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	result := make([]T2, len(s))
	for i, v := range s {
		result[i] = f(v)
	}

	return result
}

// Filter filters values from a slice using a filter function.
// It returns a new slice with only the elements of s
// for which f returned true.
func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}

	return r
}

// Reduce reduces a []T to a single value using a reduction function.
func Reduce[T, R any](s []T, initializer R, f func(R, T) R) R {
	r := initializer
	for _, v := range s {
		r = f(r, v)
	}

	return r
}
