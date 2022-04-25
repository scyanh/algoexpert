package main

import (
	"fmt"
	"github.com/scyanh/algoexpert/strings"
)

func main() {
	fmt.Println("hello")
	str:="a b c d e"
	//str:="AlgoExpert is the best!"

	f := strings.NewReverseWordsInString()
	res := f.ReverseWordsInString(str)
	fmt.Println("res=", res)
}
