package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/upload", uploadFile)

	log.Print("Starting server on port 8000")
	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}
