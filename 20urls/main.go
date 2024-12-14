package main

import (
	"fmt"
	"net/url"
)

const myurl string = "http://localhost:4000/web/contact"

func main() {
	fmt.Println("Welcome to handling urls in golang")
	fmt.Println(myurl)
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	qparams := result.Query()
	fmt.Printf("The type of data type : %T \n", qparams)
	fmt.Println(qparams["contact"])

	partsOfUrl := &url.URL{
		Scheme: "http",
		Host:   "localhost:4000",
		Path:   "/web/contact",
	}
	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
