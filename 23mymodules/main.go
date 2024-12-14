package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Here is mymodule topic")
	greeting()
	r := mux.NewRouter()
	r.HandleFunc("/", showServer).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeting() {
	fmt.Println("This is greeting from golang")
}

func showServer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the moodule topic of golang </h1>"))
}
