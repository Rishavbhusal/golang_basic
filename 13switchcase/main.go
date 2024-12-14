package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Switch case in golang")
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of diceNumber is:", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You can move to spot 1")
		// fallthrough
	case 2:
		fmt.Println("You can move to spot 2")
	case 3:
		fmt.Println("You can move to spot 3")
	case 4:
		fmt.Println("You can move to spot 4")
	case 5:
		fmt.Println("You can move to spot 5")
	case 7:
		fmt.Println("You can movev to spot 6 and you can roll dice again")
	default:
		fmt.Println("What was it?")
	}

}
