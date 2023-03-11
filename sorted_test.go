package sorted

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	name  string
	self  []int
	other []int
}{
	{
		name:  "Empty",
		self:  []int{},
		other: []int{},
	},
	{
		name:  "SamenessOne",
		self:  []int{1},
		other: []int{1},
	},
	{
		name:  "DifferenceOne",
		self:  []int{1},
		other: []int{2},
	},
	{
		name:  "SamenessTwo",
		self:  []int{1, 2},
		other: []int{1, 2},
	},
	{
		name:  "DifferenceTwo",
		self:  []int{1, 3},
		other: []int{2, 4},
	},
	{
		name:  "Long1",
		self:  sortedSlice(10, 20, 0),
		other: sortedSlice(10, 20, 1),
	},
	{
		name:  "Long2",
		self:  sortedSlice(10, 20, 3),
		other: sortedSlice(10, 20, 4),
	},
	{
		name:  "Long3",
		self:  sortedSlice(100, 200, 3),
		other: sortedSlice(100, 200, 4),
	},
}

func sortedSlice(length, max int, seed int64) []int {
	random := rand.New(rand.NewSource(seed))
	slice := lo.Times(length, func(_ int) int { return random.Intn(max) })
	sort.Ints(slice)
	return slice
}

func TestIntersect(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			expected := lo.Intersect(testCase.other, testCase.self)
			actual := Intersect(testCase.self, testCase.other, Compare[int])

			assert.Equal(t, expected, actual, "self: %v, other: %v", testCase.self, testCase.other)
		})
	}
}

func TestDifference(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			expectedSelf, expectedOther := lo.Difference(testCase.self, testCase.other)
			actualSelf := Difference(testCase.self, testCase.other, Compare[int])
			actualOther := Difference(testCase.other, testCase.self, Compare[int])

			assert.Equal(t, expectedSelf, actualSelf, "self: %v, other: %v", testCase.self, testCase.other)
			assert.Equal(t, expectedOther, actualOther, "self: %v, other: %v", testCase.self, testCase.other)
		})
	}
}

func TestUnique(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			expectedSelf := lo.Uniq(testCase.self)
			expectedOther := lo.Uniq(testCase.other)
			actualSelf := Unique(testCase.self, Equal[int])
			actualOther := Unique(testCase.other, Equal[int])

			assert.Equal(t, expectedSelf, actualSelf, "self: %v", testCase.self)
			assert.Equal(t, expectedOther, actualOther, "self: %v", testCase.other)
		})
	}
}
