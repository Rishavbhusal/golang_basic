package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to SR  resturant"
	fmt.Println(welcome)
	fmt.Println("Enter your food number to order")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Your food number is :", input)
	fmt.Printf("Your datastype is : %T ", input)

}
