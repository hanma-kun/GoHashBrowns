package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/hanma-kun/GoHashBrowns/cryptoutils/encoding"
	"github.com/hanma-kun/GoHashBrowns/cryptoutils/hashing"
	"github.com/hanma-kun/GoHashBrowns/cryptoutils/dnsres"
	"github.com/hanma-kun/GoHashBrowns/cryptoutils/password"
	"github.com/hanma-kun/GoHashBrowns/cryptoutils/random"
)

func main() {
	fmt.Println("Select an option:")
	fmt.Println("1. Generate SHA-256 hash")
	fmt.Println("2. Perform Base64 encoding/decoding")
	fmt.Println("3. Generate secure random number")
	fmt.Println("4. Brute force password cracker")
	fmt.Println("5. Resolve domain to IP addresses")

	reader := bufio.NewReader(os.Stdin)
	option, _ := reader.ReadString('\n')
	option = strings.TrimSpace(option)

	switch option {
	case "1":
		fmt.Print("Enter the data to hash: ")
		data, _ := reader.ReadString('\n')
		data = strings.TrimSpace(data)
		hash := hashing.HashSHA256(data)
		fmt.Println("SHA-256 Hash:", hash)

	case "2":
		fmt.Print("Enter the data to encode/decode: ")
		data, _ := reader.ReadString('\n')
		data = strings.TrimSpace(data)

		fmt.Println("Select an option:")
		fmt.Println("1. Encode to Base64")
		fmt.Println("2. Decode from Base64")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			encoded := encoding.Base64Encode([]byte(data))
			fmt.Println("Encoded data:", encoded)

		case "2":
			decoded, err := encoding.Base64Decode(data)
			if err != nil {
				fmt.Println("Error decoding data:", err)
				return
			}
			fmt.Println("Decoded data:", string(decoded))

		default:
			fmt.Println("Invalid option.")
		}

	case "3":
		randomNumber, err := random.GenerateSecureRandomNumber()
		if err != nil {
			fmt.Println("Error generating secure random number:", err)
			return
		}
		fmt.Println("Secure random number:", randomNumber)

	case "4":
		fmt.Print("Enter the password to crack: ")
		userPassword, _ := reader.ReadString('\n')
		userPassword = strings.TrimSpace(userPassword)
		maxLength := 6 // You can adjust the maximum length as needed

		found := password.BruteForce(userPassword, maxLength)
		if found == "" {
			fmt.Println("Password not found!")
		} else {
			fmt.Println("Password found:", found)
		}

	case "5":
		fmt.Print("Enter the domain to resolve: ")
		domain, _ := reader.ReadString('\n')
		domain = strings.TrimSpace(domain)

		ips, err := dnsres.ResolveDomain(domain)
		if err != nil {
			fmt.Println("Error resolving domain:", err)
			return
		}

		fmt.Println("IP addresses associated with the domain:")
		for _, ip := range ips {
			fmt.Println(ip)
		}

	default:
		fmt.Println("Invalid option.")
	}
}