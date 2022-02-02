package arrays

import (
	"fmt"
	"sort"
)

type arrayObj struct {
}

func NewArrayObj() arrayObj {
	return arrayObj{}
}

// Sorted squared array
// Write a function that takes in a non-empty array of integers that are sorted
// in ascending order and returns a new array of the same length with the squares
// of the original integers also sorted in ascending order.
func (ao arrayObj) SortedSquaredArray(array []int) []int {
	fmt.Println("SortedSquaredArray")

	sortedSquares := make([]int, len(array))

	for idx, val := range array {
		sortedSquares[idx] = val * val
	}
	sort.Ints(sortedSquares)

	return sortedSquares
}
