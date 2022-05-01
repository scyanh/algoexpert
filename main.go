package main

import (
	"fmt"
	"github.com/scyanh/algoexpert/recursion"
)

func main() {
	fmt.Println("hello")
	//arr:=[]int{1, 2, 3}
	phoneNumber:="123"

	res := recursion.NewPhoneNumberMnemonics().PhoneNumberMnemonics(phoneNumber)
	fmt.Println("res=", res)
}
