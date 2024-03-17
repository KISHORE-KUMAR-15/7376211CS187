package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Kishore Kumar k")
	})

	http.HandleFunc("/Dept", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "im from CSE")
	})

	http.HandleFunc("/Rollnum", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "7376211cs187")
	})

	fmt.Printf("Server running (port=8080), route: http://localhost:8080/helloworld\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
