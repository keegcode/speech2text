package recognizer

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type FileType string

const (
	Video FileType = "video"
	Audio FileType = "audio"
)

type Media struct {
	FileType FileType
	Size     int64
	Path     string
	Language string
}

func (m *Media) GetAudio() (*Media, error) {
	if m.FileType == Audio {
		return m, nil
	}

	return m.extractAudioFromVideo()
}

func (m *Media) extractAudioFromVideo() (*Media, error) {
	audioPath := fmt.Sprintf("%s.%s", strings.Split(m.Path, ".")[0], ".mp3")
	args := []string{"-i", m.Path, "-q:a", "192", "-map", "a", audioPath}

	err := exec.Command("ffmpeg", args...).Run()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(audioPath)
	if err != nil {
		return nil, err
	}

	fileStat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	m.Path = audioPath
	m.FileType = Audio
	m.Size = fileStat.Size()

	return m, nil
}
