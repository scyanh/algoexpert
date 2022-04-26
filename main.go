package main

import (
	"fmt"
	"github.com/scyanh/algoexpert/strings"
)

func main() {
	fmt.Println("hello")
	str:=[]string{"this", "that", "did", "deed", "them!", "a"}


	res := strings.NewMinimumCharactersForWords().MinimumCharactersForWords(str)
	fmt.Println("res=", res)
}
