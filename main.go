package main

import (
	"fmt"
	"github.com/scyanh/algoexpert/recursion"
)

func main() {
	fmt.Println("hello")
	arr:=[]int{1, 2, 3}

	res := recursion.NewPermutations().GetPermutations(arr)
	fmt.Println("res=", res)
}
