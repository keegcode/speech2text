package recognizer

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

type OpenAIRecognizer struct {
	apiKey string
}

func (r OpenAIRecognizer) GetToken() string {
	return r.apiKey
}

func (r OpenAIRecognizer) RecognizeTextInAudio(m Media) (string, error) {
	file, err := os.Open(m.Path)
	if err != nil {
		return "", err
	}

	defer file.Close()

	buffer := new(bytes.Buffer)
	writer := multipart.NewWriter(buffer)

	part, err := writer.CreateFormFile("file", m.Path)
	if err != nil {
		return "", err
	}

	io.Copy(part, file)

	writer.WriteField("model", "whisper-1")
	writer.WriteField("response_format", "text")
	writer.WriteField("language", "uk")

	err = writer.Close()
	if err != nil {
		return "", err
	}

	uri := "https://api.openai.com/v1/audio/transcriptions"

	req, err := http.NewRequest("POST", uri, buffer)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", r.GetToken()))

	client := new(http.Client)
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	body := new(bytes.Buffer)

	io.Copy(body, resp.Body)

	if resp.StatusCode != 200 {
		log.Print(fmt.Sprintf("Failed to request the OpenAI API, statusCode: %d, status: %s, body: %s", resp.StatusCode, resp.Status, string(body.Bytes())))
		return "", errors.New("Failed To Recognize Speech!")
	}

	return string(body.Bytes()), nil
}
