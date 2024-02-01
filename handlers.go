package main

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
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

	filePath, err := writeFile(handler.Filename, file)
	if err != nil {
		log.Print(err)
		http.Error(w, "Failed to upload the file!", http.StatusBadRequest)
		return
	}

	language := r.URL.Query().Get("lang")
	if language == "" {
		language = "en"
	}

	recognizer := LocalRecognizer{}
	media := Media{fileType: FileType(fileType), size: handler.Size, path: filePath, recognizer: recognizer, language: language}

	text, err := media.GenerateText()
	if err != nil {
		http.Error(w, "Failed to upload the file!", http.StatusBadRequest)
		return
	}

	w.Write([]byte(text))
}

func Home(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func writeFile(filename string, file multipart.File) (string, error) {
	path, err := getPath(filename)
	if err != nil {
		return "", err
	}

	dst, err := os.Create(path)
	if err != nil {
		return "", err
	}

	defer dst.Close()
	if _, err := io.Copy(dst, file); err != nil {
		return "", err
	}

	return path, nil
}

func getPath(filename string) (string, error) {
	id, err := uuid.NewV7()
	if err != nil {
		log.Print(err)
		return "", err
	}

	return fmt.Sprintf("/tmp/%s%s", id, filepath.Ext(filename)), nil
}
