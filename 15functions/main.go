package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()
	result := adder(2, 3)
	fmt.Println("The result is", result)
}

// output:datatype
func adder(val1 int, val2 int) int {
	return val1 + val2

}

func greeter() {
	fmt.Println("Namaste from golang")
}
