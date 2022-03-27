package dynamicprogramming

import (
	"fmt"
)

type numberOfWaysToMakeChange struct {
}

func NewNumberOfWaysToMakeChangeService() numberOfWaysToMakeChange {
	return numberOfWaysToMakeChange{}
}

// NumberOfWaysToMakeChange
// Given an array of distinct positive integers representing coin denominations and a
//  single non-negative integer "n" representing a target amount of
//  money, write a function that returns the number of ways to make change for
//  that target amount using the given coin denominations.
//
func (s numberOfWaysToMakeChange) NumberOfWaysToMakeChange(array []int) int {
	fmt.Println("NumberOfWaysToMakeChange")

	return 1
}

func (s numberOfWaysToMakeChange) max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
