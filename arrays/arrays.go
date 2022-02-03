package arrays

import (
	"fmt"
	"sort"
)

type sortedSquaredArrayService struct {
}

func NewSortedSquaredArrayService() sortedSquaredArrayService {
	return sortedSquaredArrayService{}
}

// Sorted squared array
// Write a function that takes in a non-empty array of integers that are sorted
// in ascending order and returns a new array of the same length with the squares
// of the original integers also sorted in ascending order.
// O(nlogn) time | O(n) space - where n is the length of the input array
func (s sortedSquaredArrayService) SortedSquaredArray(array []int) []int {
	fmt.Println("SortedSquaredArray")

	sortedSquares := make([]int, len(array))

	for idx, val := range array {
		sortedSquares[idx] = val * val
	}
	sort.Ints(sortedSquares)

	return sortedSquares
}

// O(n) time | O(n) space - where n is the length of the input array
func (s sortedSquaredArrayService) SortedSquaredArrayV2(array []int) []int {
	fmt.Println("SortedSquaredArrayV2")

	sortedSquares := make([]int, len(array))

	smallerValueIdx := 0
	largerValueIdx := len(array) - 1

	for idx := len(array) - 1; idx >= 0; idx-- {
		smallerValue := abs(array[smallerValueIdx])
		largerValue := abs(array[largerValueIdx])
		if smallerValue >= largerValue {
			sortedSquares[idx] = smallerValue * smallerValue
			smallerValueIdx++
		} else {
			sortedSquares[idx] = largerValue * largerValue
			largerValueIdx--
		}
	}

	return sortedSquares
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
