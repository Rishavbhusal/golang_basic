package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to json ")
	// EncodedJson()
	DecodeJson()
}

func EncodedJson() {
	digitalCourses := []course{
		{"ReactJS Bootcamp", 299, "LEarning.in", "abc212", []string{"web-dev", "js"}},
		{"Angular ", 254, "Angu.in", "bcd21", []string{"Full-stack", "js"}},
		{"Mern", 323, "Mern.in", "kdfgd", nil}}

	finalJson, err := json.MarshalIndent(digitalCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", finalJson)
}

func DecodeJson() {
	jsonDataWeb := []byte(`
	       {
                "coursename": "Angular ",
                "Price": 254,
                "website": "Angu.in",
                "tags": [
                        "Full-stack",
                        "js"
                ]
        }
		`)
	var digitalCourse course
	checkValid := json.Valid(jsonDataWeb)
	if checkValid {
		fmt.Println("Json valid")
		json.Unmarshal(jsonDataWeb, &digitalCourse)
		fmt.Printf("%#v \n", digitalCourse)

	} else {
		fmt.Println("Json wqas invalid")
	}

	// some cases where you want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataWeb, &myOnlineData)
	fmt.Printf("%#v \n", myOnlineData)

	for key, value := range myOnlineData {
		fmt.Printf("KEy is %v and value is %v and type is : %T \n", key, value, value)
	}

}
