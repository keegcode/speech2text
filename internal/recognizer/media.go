package recognizer

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

func (m Media) extractAudioFromVideo() (Media, error) {
	return m, nil
}
