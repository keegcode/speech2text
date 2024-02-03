package main

import (
	"log"
	"net/http"

	"github.com/gofor-little/env"
	"github.com/keegcode/speech2text/internal/recognizer"
)

type application struct {
	recognizer recognizer.Recognizer
}

func main() {
	env.Load(".env")

	app := application{recognizer: recognizer.LocalRecognizer{}}

	log.Print("Starting server on port 8000")
	err := http.ListenAndServe(":8000", app.routes())
	log.Fatal(err)
}
