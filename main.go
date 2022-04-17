package main

import (
	"fmt"
	"github.com/scyanh/algoexpert/strings"
)

func main() {
	fmt.Println("hello")
	str:="Bste!hetsi ogEAxpelrt x "
	str2:="AlgoExpert is the Best!"
	//input := []int{7, 10, 12, 7, 9, 14}

	f := strings.NewGenerateDocument()
	res := f.GenerateDocument(str, str2)
	fmt.Println("res=", res)
}
