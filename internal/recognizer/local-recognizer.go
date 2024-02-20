package recognizer

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type LocalRecognizer struct {
}

func (r LocalRecognizer) GetToken() string {
	return ""
}

func (r LocalRecognizer) RecognizeTextInAudio(m *Media) (string, error) {
	args := []string{m.Path, "--model", "small", "--output_format", "srt", "--language", m.Language, "--task", "transcribe", "-o", "/tmp/"}
	err := exec.Command("whisper", args...).Run()
	if err != nil {
		return "", err
	}

	path := fmt.Sprintf("%s.%s", strings.Split(m.Path, ".")[0], "srt")

	file, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	err = os.Remove(path)
	if err != nil {
		return "", err
	}

	return string(file), nil
}
