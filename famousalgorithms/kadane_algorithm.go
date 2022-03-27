package famousalgorithms

type kadanesAlgorithm struct {
}

func NewKadanesAlgorith() kadanesAlgorithm {
	return kadanesAlgorithm{}
}

// KadanesAlgorithm
// Write a function that takes in a non-empty array of integers and returns the
// maximum sum that can be obtained by summing up all of the integers in a
// non-empty subarray of the input array. A subarray must only contain adjacent
// numbers (numbers next to each other in the input array).
//
// Sample input [3, 5, -9, 1, 3, -2, 3, 4, 7, 2, -9, 6, 3, 1, -5, 4]
// Sample output 19 // [1, 3, -2, 3, 4, 7, 2, -9, 6, 3, 1]
//
func (k kadanesAlgorithm) KadanesAlgorithm(array []int) int {
	maxNumber := array[0]
	maxToHere := array[0]

	for i := 1; i < len(array); i++ {
		maxToHere = k.max(array[i], maxToHere + array[i])
		maxNumber = k.max(maxToHere, maxNumber)
	}

	return maxNumber
}

func (kadanesAlgorithm) max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
