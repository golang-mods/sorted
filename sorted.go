package sorted

// Intersect returns the intersection between two collections.
// The collection of arguments must be sorted.
func Intersect[T any, U any](self []T, other []U, compare func(T, U) int) []T {
	result := []T{}

	for i, j := 0, 0; i < len(self) && j < len(other); {
		c := compare(self[i], other[j])
		switch {
		case c < 0:
			i++
		case c > 0:
			j++
		default:
			result = append(result, self[i])
			i++
		}
	}

	return result
}

// Difference returns the difference between two collections.
// The collection of arguments must be sorted.
func Difference[T any, U any](self []T, other []U, compare func(T, U) int) []T {
	result := []T{}

	i := 0
	for j := 0; i < len(self) && j < len(other); {
		c := compare(self[i], other[j])
		switch {
		case c < 0:
			result = append(result, self[i])
			i++
		case c > 0:
			j++
		default:
			i++
		}
	}

	return append(result, self[i:]...)
}

// Unique returns the removes duplicates a collection.
// The collection of arguments must be sorted.
func Unique[T any](self []T, equal func(T, T) bool) []T {
	result := []T{}

	if len(self) == 0 {
		return result
	}

	i := 0
	for j := 1; i < len(self) && j < len(self); j++ {
		if !equal(self[i], self[j]) {
			result = append(result, self[i])
			i = j
		}
	}

	return append(result, self[i])
}
