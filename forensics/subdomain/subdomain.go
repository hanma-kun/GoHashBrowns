package subdomain

import (
	"fmt"
	"net"
	"strings"
	"os"
	"bufio"
)

func EnumerateSubdomains(targetDomain, wordlistPath string) {
	wordlist, err := readWordlist(wordlistPath)
	if err != nil {
		fmt.Println("Error reading wordlist:", err)
		return
	}

	for _, subdomain := range wordlist {
		fullDomain := subdomain + "." + targetDomain
		ips, err := net.LookupHost(fullDomain)
		if err == nil {
			fmt.Printf("Subdomain found: %s (%v)\n", fullDomain, ips)
		}
	}
}

func readWordlist(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var wordlist []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word != "" {
			wordlist = append(wordlist, word)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return wordlist, nil
}