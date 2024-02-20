package main

import (
	"errors"
	"log"
	"net/http"
)

const maxFileSizeMB = 25 << 20

var allowedContentTypes = map[string]bool{
	"audio/mpeg": true,
	"audio/mp4":  true,
	"video/mpeg": true,
	"video/mp4":  true,
	"audio/wav":  true,
	"audio/webm": true,
	"video/webm": true,
}

type MultipartData struct {
	Path        string
	ContentType string
	Size        int64
}

func ProcessMultipart(w http.ResponseWriter, r *http.Request) (*MultipartData, error) {
	r.ParseMultipartForm(maxFileSizeMB)

	file, handler, err := r.FormFile("media")
	if err != nil {
		log.Print(err)
		return nil, errors.New("failed to upload the file")
	}

	defer file.Close()

	log.Printf("Uploaded File: %+v\n", handler.Filename)
	log.Printf("File Size: %+v\n", handler.Size)
	log.Printf("MIME Header: %+v\n", handler.Header)

	contentType := handler.Header.Get("Content-Type")

	if !allowedContentTypes[contentType] {
		return nil, errors.New("failed to upload the file")
	}

	filePath, err := writeFile(handler.Filename, file)
	if err != nil {
		log.Print(err)
		return nil, errors.New("failed to upload the file")
	}

	return &MultipartData{Path: filePath, ContentType: contentType, Size: handler.Size}, nil
}
