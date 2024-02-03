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

	"github.com/keegcode/speech2text/internal/recognizer"
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

func (app *application) UploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed!", http.StatusMethodNotAllowed)
		return
	}

	r.ParseMultipartForm(maxFileSizeMB)

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

	if !allowedContentTypes[contentType] {
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

	media := recognizer.Media{FileType: recognizer.FileType(fileType), Size: handler.Size, Path: filePath, Language: language}

	text, err := app.recognizer.RecognizeTextInAudio(media)
	if err != nil {
		http.Error(w, "Failed to upload the file!", http.StatusBadRequest)
		return
	}

	w.Write([]byte(text))
}

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
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
