package main

import (
	"fmt"
	"github.com/scyanh/algoexpert/recursion"
)

func main() {
	fmt.Println("hello")
	arr:=[]int{1, 2, 3}

	res := recursion.NewPowerset().Powerset(arr)
	fmt.Println("res=", res)
}
