package recognizer

type Recognizer interface {
	GetToken() string
	RecognizeTextInAudio(m Media) (string, error)
}
