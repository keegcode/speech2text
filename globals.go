package main

var ALLOWED_CONTENT_TYPES = map[string]bool{
	"audio/mpeg": true,
	"video/mp4":  true,
	"video/mpeg": true,
	"audio/wav":  true,
	"audio/webm": true,
	"video/webm": true,
}

const MAX_FILE_SIZE_MB = 25 << 20
