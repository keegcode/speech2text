package main

type FileType string

const (
	Video FileType = "video"
	Audio FileType = "audio"
)

type Media struct {
	fileType   FileType
	size       int64
	path       string
	language   string
	recognizer Recognizer
}

type Recognizer interface {
	GetToken() string
	RecognizeTextInAudio(m Media) (string, error)
}
