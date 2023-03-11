package sorted

import (
	"fmt"
	"testing"

	"github.com/samber/lo"
)

var benchmarkCases = []struct {
	length int
	max    int
}{
	{length: 30, max: 100},
	{length: 300, max: 1000},
	{length: 3000, max: 10000},
}

var functions = []struct {
	name     string
	function func(self, other []int)
}{
	{name: "Intersect", function: func(self, other []int) {
		Intersect(self, other, Compare[int])
	}},
	{name: "lo.Intersect", function: func(self, other []int) {
		lo.Intersect(self, other)
	}},
	{name: "Difference", function: func(self, other []int) {
		Difference(self, other, Compare[int])
		Difference(other, self, Compare[int])
	}},
	{name: "lo.Difference", function: func(self, other []int) {
		lo.Difference(self, other)
	}},
	{name: "Unique", function: func(self, other []int) {
		Unique(self, Equal[int])
		Unique(other, Equal[int])
	}},
	{name: "lo.Uniq", function: func(self, other []int) {
		lo.Uniq(self)
		lo.Uniq(other)
	}},
}

func BenchmarkAll(b *testing.B) {
	for _, benchmarkCase := range benchmarkCases {
		self := sortedSlice(benchmarkCase.length, benchmarkCase.max, 1)
		other := sortedSlice(benchmarkCase.length, benchmarkCase.max, 2)

		for _, function := range functions {
			b.Run(fmt.Sprintf("%s/%d/%d", function.name, benchmarkCase.length, benchmarkCase.max), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					function.function(self, other)
				}
			})
		}
	}
}
