package api

import (
	"fmt"
	"net/http"
)

// HandleRoot handles the root endpoint
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received request...")
	fmt.Fprintf(w, "Hello, you've reached the Global Warfront server!")
}

// Hello handles the hello endpoint
func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}
