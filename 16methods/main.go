package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Methods in golang")

	rishav := User{"Rishav", "rishav@gmail.com", true, 12}
	fmt.Println(rishav)
	rishav.GetStatus()
	rishav.NewEmail()
}

func (u User) GetStatus() {
	fmt.Println("IS user active :", u.Status)
}
func (u User) NewEmail() {
	u.Email = "test@gmail.com"
	fmt.Println("The new email of this user is : ", u.Email)

}
