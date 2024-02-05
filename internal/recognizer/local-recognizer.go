package recognizer

import (
	"os/exec"
)

type LocalRecognizer struct {
}

func (r LocalRecognizer) GetToken() string {
	return ""
}

func (r LocalRecognizer) RecognizeTextInAudio(m *Media) (string, error) {
	args := []string{m.Path, "--model", "small", "--output_format", "txt", "--language", m.Language, "--task", "transcribe", "-o", "./output/"}
	out, err := exec.Command("whisper", args...).Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}
