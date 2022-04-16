package main

import (
	"fmt"
	"github.com/scyanh/algoexpert/strings"
)

func main() {
	fmt.Println("hello")
	str:="AAAAAAAAAAAAABBCCCCDD"
	//input := []int{7, 10, 12, 7, 9, 14}

	f := strings.NewRunLengthEncoding()
	res := f.RunLengthEncoding(str)
	fmt.Println("res=", res)
}
