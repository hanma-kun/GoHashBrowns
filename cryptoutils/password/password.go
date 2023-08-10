package password

import (
	"strings"
	"sync"
)

// The character set used to generate all possible password combinations
const charset = "abcdefghijklmnopqrstuvwxyz"

// BruteForce cracks a password using brute force technique.
func BruteForce(password string, maxLength int) string {
	var wg sync.WaitGroup
	// Channel to send matching password guesses
	found := make(chan string)

	// Create goroutines to generate password combinations of different lengths
	for length := 1; length <= maxLength; length++ {
		wg.Add(1)
		go func(length int) {
			defer wg.Done()
			// Generate combinations of the given length and send to the channel
			generateCombinations(charset, length, "", found)
		}(length)
	}

	// Read from the channel until a matching password guess is found
	for passwordGuess := range found {
		if strings.Compare(password, passwordGuess) == 0 {
			close(found)
			return passwordGuess
		}
	}

	return ""
}

// Recursive function to generate all combinations of characters of a given length
func generateCombinations(charset string, length int, prefix string, found chan<- string) {
	// If the desired length is 0, send the generated combination to the channel
	if length == 0 {
		found <- prefix
		return
	}
	// Recursively generate combinations by appending each character from the character set
	for _, c := range charset {
		generateCombinations(charset, length-1, prefix+string(c), found)
	}
}