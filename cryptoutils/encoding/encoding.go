package encoding

import (
	"encoding/base64"
)

// Base64Encode takes a byte slice as input and returns its Base64 encoded string.
func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Base64Decode takes a Base64 encoded string as input and returns the decoded byte slice and an error if any.
func Base64Decode(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}
