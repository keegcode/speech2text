package main

import (
	"log"
	"net/http"
	"strings"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed!", http.StatusMethodNotAllowed)
		return
	}

	r.ParseMultipartForm(MAX_FILE_SIZE_MB)

	file, handler, err := r.FormFile("media")
	if err != nil {
		log.Print(err)
		http.Error(w, "Failed to upload the file!", http.StatusInternalServerError)
		return
	}

	defer file.Close()

	log.Printf("Uploaded File: %+v\n", handler.Filename)
	log.Printf("File Size: %+v\n", handler.Size)
	log.Printf("MIME Header: %+v\n", handler.Header)

	contentType := handler.Header.Get("Content-Type")

	if !ALLOWED_CONTENT_TYPES[contentType] {
		http.Error(w, "Failed to upload the file!", http.StatusBadRequest)
		return
	}

	fileType := strings.Split(contentType, "/")[0]

	media := Media{fileType: FileType(fileType), size: handler.Size}
	text := media.generateText()

	w.Write([]byte(text))
}

func home(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
