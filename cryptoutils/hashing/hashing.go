package hashing

import (
	"crypto/sha256"
	"encoding/hex"
)

// HashSHA256 takes a string as input and returns its SHA-256 hash in hexadecimal format.
func HashSHA256(data string) string {
	hasher := sha256.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}