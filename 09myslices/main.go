package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("we are at slices ")
	var fruitList = []string{"mango", "banana", "peach", "banana", "apple"}
	fmt.Printf("the type of fruitList is : %T \n", fruitList)
	fmt.Println("the length is", len(fruitList))
	fruitList = append(fruitList, "green apple", "Tomato")
	fmt.Println(fruitList)
	fruitList = append(fruitList[1:4])
	fmt.Println(fruitList)
	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highscore := make([]int, 4)
	highscore[0] = 234
	highscore[1] = 121
	highscore[2] = 2000
	highscore[3] = 2
	fmt.Println(highscore)

	highscore = append(highscore, 4000, 32, 12)
	fmt.Println(highscore)

	sort.Ints(highscore)
	fmt.Println(highscore)
	fmt.Println(sort.IntsAreSorted(highscore))

	fruits := make([]string, 4)
	fruits[0] = "mango"
	fruits[1] = "orange"
	fruits[2] = "Apple"
	fruits[3] = "Banana"
	fmt.Println(fruits)
	fmt.Println("the length is", len(fruits))

	// how to remove a value from slices based on index

	var courses = []string{"python", "javascript", "React", "NodeJS", "Ruby"}
	fmt.Println(courses)
	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
