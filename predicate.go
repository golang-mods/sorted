package sorted

import "golang.org/x/exp/constraints"

// Compare two values return -1 for less, +1 for grater, 0 for same.
func Compare[T constraints.Ordered](self T, other T) int {
	if self < other {
		return -1
	} else if self > other {
		return 1
	}
	return 0
}

// Compare two values returns true if they are same.
func Equal[T comparable](self, other T) bool {
	return self == other
}
