package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Simple Hello API")
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World")
	})

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}