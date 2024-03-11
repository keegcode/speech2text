package recognizer

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
  "path/filepath"
  "runtime"
)

type LocalRecognizer struct {
}

func (r LocalRecognizer) GetToken() string {
	return ""
}

func (r LocalRecognizer) RecognizeTextInAudio(m *Media) (string, error) {
  cmd, err := filepath.Abs("./whisper.cpp/main")
  if err != nil {
    return "", err
  }

  model, err := filepath.Abs("./whisper.cpp/models/ggml-large-v3.bin")
  if err != nil {
    return "", err
  }

	args := []string{
    m.Path, 
    "-m", model, 
    "-t", fmt.Sprintf("%d", runtime.NumCPU()), 
    "-osrt", 
    "-l", m.Language, 
  }

	err = exec.Command(cmd, args...).Run()
	if err != nil {
		return "", err
	}

	path := fmt.Sprintf("%s.%s", strings.Split(m.Path, ".")[0], "wav.srt")

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
