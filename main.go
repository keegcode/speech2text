package main

import (
	"log"
	"net/http"

	"github.com/gofor-little/env"
)

func main() {
	env.Load(".env")

	mux := http.NewServeMux()

	mux.HandleFunc("/", Home)
	mux.HandleFunc("/upload", UploadFile)

	log.Print("Starting server on port 8000")
	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}
