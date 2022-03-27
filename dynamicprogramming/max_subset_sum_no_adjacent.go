package dynamicprogramming

import (
	"fmt"
)

type maxSubsetSumNoAdjacent struct {
}

func NewMaxSubsetSumNoAdjacentService() maxSubsetSumNoAdjacent {
	return maxSubsetSumNoAdjacent{}
}

// MaxSubsetSumNoAdjacent
// Write a function that takes in an array of positive integers and returns the
// maximum sum of non-adjacent elements in the array.
//
// If the input array is empty, the function should return 0
//
// Sample input [75, 105, 120, 75, 90, 135]
// Sample output 330 // 75 + 120 + 135
func (s maxSubsetSumNoAdjacent) MaxSubsetSumNoAdjacent(array []int) int {
	fmt.Println("SortedSquaredArray")

	if len(array) == 0 {
		return 0
	} else if len(array) == 1 {
		return array[0]
	}

	maxSums := make([]int, len(array))
	maxSums[0], maxSums[1] = array[0], s.max(array[0], array[1])
	for i := 2; i < len(array); i++ {
		maxSums[i] = s.max(maxSums[i-1], maxSums[i-2]+array[i])
	}
	return maxSums[len(array)-1]
}

func (s maxSubsetSumNoAdjacent) max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
