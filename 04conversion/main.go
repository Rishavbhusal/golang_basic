package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Welcome to our pizza store"
	fmt.Println(welcome)
	fmt.Println("Enter your rating from 1 to 5")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Your rating to our pizza store is:", input)
	numInput, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("your rating with one added is :", numInput+1)
	}
}
