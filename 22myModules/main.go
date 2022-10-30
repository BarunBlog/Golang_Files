package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in golang")
	greeter()

	// url routes
	route := mux.NewRouter()
	route.HandleFunc("/", serveHome).Methods("GET")

	//http.ListenAndServe(":4000", route)

	log.Fatal(http.ListenAndServe(":4000", route)) // if anything failed then it will logs the error
}

func greeter() {
	fmt.Println("Hey there mod users")
}

// writer and reader. reader is reading the value whoever sending the request
// writer is writing the request
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to golang series on youtube </h1>"))
}
