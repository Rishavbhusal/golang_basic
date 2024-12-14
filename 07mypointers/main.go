package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hurray! we are in pointer of golang")
	var ptr *int
	fmt.Println("Value of Pointer is :", ptr)

	number := 23
	var pointer = &number
	fmt.Println(" Actual memory address pinted by pointer is :", pointer)
	fmt.Println("Value of actual memory address pinted by pointer is :", *pointer)

	*pointer += 22
	fmt.Println(*pointer)
}
