package main

import (
	"fmt"
)

func main() {
	fmt.Println("IFelse ingolang")
	var loginCount = 10
	var result string
	if loginCount < 10 {
		result = "regular USer"
	} else if loginCount == 10 {
		result = "exactly 10 user"
	} else {
		result = "Bulk users"
	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("It is even")
	} else {
		fmt.Println("ODdd")
	}
	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is greeater than 10")
	}
}
