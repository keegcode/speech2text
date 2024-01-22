package main

type FileType string

const (
	Video FileType = "video"
	Audio FileType = "audio"
)

type Media struct {
	fileType FileType
	size     int64
}

func (m Media) generateText() string {
	return "Sample text"
}
