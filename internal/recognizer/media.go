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
	return m.convertAudioToWAV()
}

func (m *Media) convertAudioToWAV() (*Media, error) {
	audioPath := fmt.Sprintf("%s.%s", strings.Split(m.Path, ".")[0], "wav")
	args := []string{
    "-i", m.Path, 
    "-c:a", "pcm_s16le", 
    "-map", "a", 
    "-ar", "16000", 
    "-ac", "1", 
    audioPath,
  }
	err := exec.Command("ffmpeg", args...).Run()
	if err != nil {
		return nil, err
	}

	err = os.Remove(m.Path)
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
