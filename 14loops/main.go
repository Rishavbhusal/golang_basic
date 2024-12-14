package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops in golang")
	days := []string{"sunday", "Monday", "Tuesday", "wednessday", "Thursdayy", "Friday"}
	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	for i := range days {
		fmt.Println(days[i])
	}

	// for each loops in golang

	for index, value := range days {
		fmt.Printf("The index is %v and value is %v \n", index, value)
	}

	// goto loop

	rougueValue := 1

	for rougueValue <= 20 {

		if rougueValue == 2 {
			goto rishav
		}
		if rougueValue == 10 {

			fmt.Println("Break")
			rougueValue++
			continue

		}

		fmt.Println("Value is ", rougueValue)
		rougueValue++
	}

rishav:
	fmt.Println("Here we used goto")
}
