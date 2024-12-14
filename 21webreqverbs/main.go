package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to webverb video")
	// PerformGetReqeust()
	PerformPostRequest()
	// PerformPostFormRequest()
}

// func PerformGetReqeust() {
// 	myurl := "http://localhost:8000/get"
// 	response, err := http.Get(myurl)
// 	NulErrorHandling(err)

// 	defer response.Body.Close()

// 	fmt.Println("Status code", response.StatusCode)
// 	fmt.Println("content length", response.ContentLength)

// 	var responseString strings.Builder
// 	content, _ := io.ReadAll(response.Body)
// 	byteCount, _ := responseString.Write(content)
// 	fmt.Println("ByteCount is :", byteCount)
// 	fmt.Println(responseString.String())

// 	// content, _ := io.ReadAll(response.Body)
// 	// fmt.Println(string(content))
// }

func PerformPostRequest() {
	myurl := "http://localhost:8000/post"
	// fake json payload
	requestBody := strings.NewReader(`
	{
	"Course" : "Lets go with golang",
	"price" : 0,
	"platform":"golang.com"
	}
	`)
	response, err := http.Post(myurl, "application/json", requestBody)
	NulErrorHandling(err)
	defer response.Body.Close()
	content, err := io.ReadAll(response.Body)
	NulErrorHandling(err)
	fmt.Println("the content is ", string(content))

}

// func PerformPostFormRequest() {
// 	myurl := "http://localhost:8000/postform"
// 	data := url.Values{}
// 	data.Add("firstname", "Rishav")
// 	data.Add("Sirname", "Bhusal")
// 	data.Add("Email", "rishav@gmail.com")
// 	response, err := http.PostForm(myurl, data)
// 	NulErrorHandling(err)
// 	defer response.Body.Close()
// 	content, err := io.ReadAll(response.Body)
// 	NulErrorHandling(err)
// 	fmt.Println(string(content))

// }
func NulErrorHandling(err error) {
	if err != nil {
		panic(err)
	}
}

