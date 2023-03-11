package sorted

import "fmt"

func ExampleIntersect() {
	result := Intersect([]int{1, 2, 4, 5}, []int{2, 3, 5}, Compare[int])
	fmt.Println(result)
	// Output: [2 5]
}

func ExampleDifference() {
	result := Difference([]int{1, 2, 4, 5}, []int{2, 3, 5}, Compare[int])
	fmt.Println(result)
	// Output: [1 4]
}

func ExampleUnique() {
	result := Unique([]int{1, 2, 2, 2, 4, 6, 6}, Equal[int])
	fmt.Println(result)
	// Output: [1 2 4 6]
}
