package random

import (
	"crypto/rand"
	"encoding/binary"
)

// GenerateSecureRandomNumber generates a cryptographically secure random uint64 number and returns it along with an error if any.
func GenerateSecureRandomNumber() (uint64, error) {
	buf := make([]byte, 8)
	_, err := rand.Read(buf)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint64(buf), nil
}