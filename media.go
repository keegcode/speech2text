package main

import (
	"errors"
)

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
	audio, err := m.extractAudioFromVideo()
	if err != nil {
		return "", err
	}

	text, err := audio.processAudio()
	if err != nil {
		return "", err
	}

	return text, nil
}

func (m Media) extractAudioFromVideo() (Media, error) {
	return m, nil
}

func (m Media) processAudio() (string, error) {
	return m.recognizer.RecognizeTextInAudio(m)
}
