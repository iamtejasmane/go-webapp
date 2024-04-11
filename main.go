package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the main home page function.
func Home(w http.ResponseWriter, r *http.Request) {

}

// About is the about page of the app.
func About(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Server starting on port %s", portNumber)
	_ = http.ListenAndServe(":8080", nil)
}
