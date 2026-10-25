package main

import (
	"fmt"
	groupie "goupie/Handler"
	"net/http"
)

func main() {

	http.HandleFunc("GET /{$}", groupie.GroupieHandler)
	http.HandleFunc("GET /artist/{id}", groupie.ArtistHandler)

	fmt.Println("http://localhost:8080/")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("500 Internal Server Error")
		return
	}
}
