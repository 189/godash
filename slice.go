package godash

import "errors"

// Iterate over each item in the slice sequentially, passing the result of the previous iteration to the next iteration's function execution
func Reduce[T any, U any](items []T, initial U, reducer func(U, T) U) U {
	accumulated := initial
	for _, item := range items {
		accumulated = reducer(accumulated, item)
	}
	return accumulated
}

// Iterate over each item in a slice sequentially
func Map[T any, U any](items []T, mapper func(T) U) []U {
	result := make([]U, len(items))
	for i, item := range items {
		result[i] = mapper(item)
	}
	return result
}

type Predicate[T comparable] func(T) bool

// Filter the slice object
func Filter[T ~[]E, E comparable](items T, fn Predicate[E]) T {
	var result T
	for _, item := range items {
		if fn(item) {
			result = append(result, item)
		}
	}
	return result
}

// Return true if any Predicate(item) == true
func Some[T ~[]E, E comparable](items T, fn Predicate[E]) bool {
	for _, item := range items {
		if fn(item) {
			return true
		}
	}
	return false
}

// Return true only if all Predicate(item) == true
func Every[T ~[]E, E comparable](items T, fn Predicate[E]) bool {
	for _, item := range items {
		if !fn(item) {
			return false
		}
	}
	return true
}

// Another forms for for loop instead of for .. range, loop will stop when callback with False result
func ForEach[T ~[]E, E any](items T, callback func(idx int, item E) bool) {
	for idx, m := range items {
		if !callback(idx, m) {
			break
		}
	}
}

// Flat returns an array with a single level deep.
func Flat[T ~[]E, E any](items []T) []E {
	totalLen := 0
	for i := range items {
		totalLen += len(items[i])
	}

	result := make(T, 0, totalLen)
	for m := range items {
		result = append(result, items[m]...)
	}

	return result
}

// Append insert new item to the end of the slice
func Push[T ~[]E, E any](s *T, items ...E) {
	*s = append(*s, items...)
}

// Pop removes and returns the last element from a slice
func Pop[T ~[]E, E any](s *T) (E, error) {
	total := len(*s)
	if total == 0 {
		zero := GetZeroValue[E]()
		return zero, errors.New("slice is empty")
	}

	lastItem := (*s)[total-1]
	*s = (*s)[:total-1]

	return lastItem, nil
}

// Shift remove and returns the first element from a slice
func Shift[T ~[]E, E any](s *T) (E, error) {
	if len(*s) == 0 {
		zero := GetZeroValue[E]()
		return zero, nil
	}

	firstItem := (*s)[0]
	*s = (*s)[1:]

	return firstItem, nil
}

// Splice modifies the slice by removing and/or adding elements
func Splice[T ~[]E, E any](s *T, start int, removes int, items ...E) error {
	totalLen := len(*s)

	if start > totalLen || start < 0 {
		return errors.New("start index out of bounds")
	}

	if removes > totalLen-start {
		removes = totalLen - start
	}

	*s = append((*s)[:start], (*s)[start+removes:]...)
	*s = append((*s)[:start], append(items, (*s)[start:]...)...)

	return nil
}

// Splice add the slice to the front of slice
func Unshift[T ~[]E, E any](s *T, items ...E) T {
	*s = append(items, (*s)...)
	return *s
}
