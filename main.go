package main

import (
	"fmt"
	"github.com/scyanh/algoexpert/famousalgorithms"
)

func main() {
	fmt.Println("hello")
	input := []int{3, 5, -9, 1, 3, -2, 3, 4, 7, 2, -9, 6, 3, 1, -5, 4}
	//input := []int{7, 10, 12, 7, 9, 14}

	f := famousalgorithms.NewKadanesAlgorith()
	res := f.KadanesAlgorithm(input)
	fmt.Println("res=", res)
}
