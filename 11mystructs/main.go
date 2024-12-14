package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("This is structs in golang")

	// there is no concept of inheritance, clid , parent etc

	rishav := User{"Rishav", "Rishav@go.dev", true, 12}

	// v is used to fetch value and +v is used to obained value by key too
	fmt.Printf("The details of rishav is : %+v \n", rishav)
	fmt.Printf("The name is %v and the email of the user is %v \n ", rishav.Name, rishav.Email)
}
