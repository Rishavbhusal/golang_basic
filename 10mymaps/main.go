package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in golang")
	language := make(map[int]string)
	language[100] = "javascrips"
	language[1] = "Python"
	language[2] = "Ruby"
	language[3] = "React"
	fmt.Println(language)
	fmt.Println("js here:", language[100])

	delete(language, 2)
	fmt.Println("List of language", language)

	// loops

	for key, value := range language {
		fmt.Printf("For key %v amd for value %v \n", key, value)
	}
	for _, value := range language {
		fmt.Printf("For key value amd for value %v \n", value)
	}

}
