package main

import (
	"fmt"
	"github.com/scyanh/algoexpert/arrays"
)

func main() {
	fmt.Println("hello")

	input := []int{1, 2, 3, 5, 6, 8, 9}
	arrays := arrays.NewArrayObj()

	sortedSquared := arrays.SortedSquaredArray(input)
	fmt.Println("sortedSquared=", sortedSquared)
}
