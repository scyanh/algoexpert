package main

import (
	"fmt"
	"github.com/scyanh/algoexpert/strings"
)

func main() {
	fmt.Println("hello")
	//str:="1921680"
	//str:="245245245245"
	str:="1"
	//str:="9743"

	f := strings.NewValidIPAddress()
	res := f.ValidIPAddresses(str)
	fmt.Println("res=", res)
}
