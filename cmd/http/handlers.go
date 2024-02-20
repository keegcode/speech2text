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

func (app *application) UploadFile(w http.ResponseWriter, r *http.Request) {
	file, err := ProcessMultipart(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	path := file.Path
	contentType := file.ContentType
	size := file.Size

	fileType := strings.Split(contentType, "/")[0]
	language := r.URL.Query().Get("lang")

	media := recognizer.Media{FileType: recognizer.FileType(fileType), Size: size, Path: path, Language: language}

	audio, err := media.GetAudio()
	if err != nil {
		http.Error(w, "Failed to extract the audio!", http.StatusBadRequest)
		return
	}

	text, err := app.recognizer.RecognizeTextInAudio(audio)
	if err != nil {
		http.Error(w, "Failed to upload the file!", http.StatusBadRequest)
		return
	}

	err = os.Remove(audio.Path)
	if err != nil {
		http.Error(w, "Failed to remove the file!", http.StatusBadRequest)
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
