package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This need to go in file - gofiles.txt"
	file, err := os.Create("./mygofile.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length is :", length)
	defer file.Close()
	reaFile("./mygofile.txt")
}

func reaFile(filename string) {
	databyte, err := os.ReadFile(filename)
	checkNilError(err)
	fmt.Println("Text data inside the file \n", string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}

}
