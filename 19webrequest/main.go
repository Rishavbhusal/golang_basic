package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "http://localhost:4000/web/home"

func main() {
	fmt.Println("Web request")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The response data type is %T \n", response)

	defer response.Body.Close()
	dataByte, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(dataByte)
	fmt.Println(content)

}
