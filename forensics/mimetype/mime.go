package mimetype

import (
	"bytes"
	"net/http"
)

func GuessMIMEType(data []byte) (string, error) {
	// Define a map of magic bytes and their corresponding MIME types
	magicBytes := map[string]string{
		"\x89PNG":              "image/png",
		"\xFF\xD8":             "image/jpeg",
		"BM":                   "image/bmp",
		"%PDF-":                "application/pdf",
		"%!PS-":                "application/postscript",
		"\x1A\x45\xDF\xA3":     "video/webm",
		"\x00\x00\x00\x0C":     "audio/x-aiff",
		"\x1F\x8B\x08":         "application/gzip",
		"\x42\x5A\x68":         "application/x-bzip2",
		"\x52\x61\x72\x21\x1A": "application/x-rar-compressed",
		"\x50\x4B\x03\x04":     "application/zip",
		// Add more magic bytes and MIME types as needed
	}

	// Iterate through the map and compare magic bytes
	for magic, mimeType := range magicBytes {
		if bytes.HasPrefix(data, []byte(magic)) {
			return mimeType, nil
		}
	}

	return "", http.ErrNotMultipart
}
