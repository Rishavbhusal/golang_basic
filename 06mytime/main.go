package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study in golang")
	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	createDate := time.Date(2004, time.December, 06, 02, 04, 06, 01, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 Monday"))
}
