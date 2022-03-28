package main

import (
	"fmt"
	"github.com/scyanh/algoexpert/famousalgorithms"
)

func main() {
	fmt.Println("hello")
	input := []int{75, 105, 120, 75, 90, 135}
	//input := []int{7, 10, 12, 7, 9, 14}

	f := famousalgorithms.NewKadanesAlgorith()
	res := f.KadanesAlgorithm(input)
	fmt.Println("res=", res)
}
