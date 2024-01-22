package main

import "errors"

type FileType string

const (
	Video FileType = "video"
	Audio FileType = "audio"
)

type Media struct {
	fileType FileType
	size     int64
	path     string
}

func (m Media) GenerateText() (string, error) {
	switch m.fileType {
	case Audio:
		return m.processAudio()
	case Video:
		return m.processVideo()
	default:
		return "", errors.New("Invalid File Type")
	}
}

func (m Media) processVideo() (string, error) {
	media, err := m.extractAudioFromVideo()
	if err != nil {
		return "", err
	}

	text, err := media.processAudio()
	if err != nil {
		return "", err
	}

	return text, nil
}

func (m Media) extractAudioFromVideo() (Media, error) {
	return m, nil
}

func (m Media) processAudio() (string, error) {
	return "Sample text", nil
}
