package main

import "fmt"

// This is an public variable cz the first character is written in capital
const LoginToken string = "dgdnggdfgdfkg"

func main() {

	// datatypes

	var username string = "Rishav"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var Numbers int = 213232
	fmt.Println(Numbers)
	fmt.Printf("Variable is of type: %T \n", Numbers)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.2434534343232
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)
	var allFloat float64 = 255.2434534343232
	fmt.Println(allFloat)
	fmt.Printf("Variable is of type: %T \n", allFloat)

	// default values and some aliases

	var variable int
	fmt.Println(variable)
	fmt.Printf("The data is of type : %T \n", variable)

	var anotherVariable string
	fmt.Println(anotherVariable)
	fmt.Printf("The data is of type : %T \n", anotherVariable)

	// Implicit type
	// where datatype is added by lexer
	var container = 2344342
	fmt.Println(container)

	// no var style
	// this style is only allowed inside of the function

	numberofUser := 3000
	fmt.Println(numberofUser)
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type : %T \n", LoginToken)

}
