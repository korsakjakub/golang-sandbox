package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server on port 3000")
	http.HandleFunc("/user", makeHTTPHandler(handleGetUserById))
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
