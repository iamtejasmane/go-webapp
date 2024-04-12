package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Server starting on port %s", portNumber)
	_ = http.ListenAndServe(":8080", nil)
}
