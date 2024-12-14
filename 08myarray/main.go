package main

import "fmt"

func main() {
	fmt.Println("Lets learn about array in golang")
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[2] = "PEach"
	fmt.Println("The fruits are : ", fruitList)
	fmt.Println("The fruits are : ", len(fruitList))

	var vegetable = [5]string{"Potato", "Tomato", "Beans", "Mushroom"}

	fmt.Println("The vegitasbles are: ", vegetable)
	fmt.Println("The len of vegy is: ", len(vegetable))
}
