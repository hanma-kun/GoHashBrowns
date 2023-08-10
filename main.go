package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/hanma-kun/GoHashBrowns/cryptoutils/dnsres"
	"github.com/hanma-kun/GoHashBrowns/cryptoutils/encoding"
	"github.com/hanma-kun/GoHashBrowns/cryptoutils/hashing"
	"github.com/hanma-kun/GoHashBrowns/cryptoutils/random"
	"github.com/hanma-kun/GoHashBrowns/forensics/mimetype"
	"github.com/hanma-kun/GoHashBrowns/forensics/portscanner"
)

func main() {
	menu()
	fmt.Println("Select an option:")
	fmt.Println("1. Generate SHA-256 hash")
	fmt.Println("2. Perform Base64 encoding/decoding")
	fmt.Println("3. Generate secure random number")
	fmt.Println("4. Resolve domain to IP addresses")
	fmt.Println("5. Find MIME Type")
	fmt.Println("6. Port Scan")

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

	case "5":
		fmt.Print("Enter the file path: ")
		filePath, _ := reader.ReadString('\n')
		filePath = strings.TrimSpace(filePath)

		// Open the file
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		// Read the file contents
		fileContents := make([]byte, 512) // Read the first 512 bytes for magic bytes detection
		_, err = file.Read(fileContents)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		mimeType, err := mimetype.GuessMIMEType(fileContents)
		if err != nil {
			fmt.Println("Error guessing MIME type:", err)
			return
		}

		fmt.Println("Guessed MIME type:", mimeType)

	case "6":
		portscanner.PortScan()

	default:
		fmt.Println("Invalid option.")

	}
}

func menu() {
	mnu := `
	 ####     ####             ##  ##     ##      ####    ##  ##            #####    #####     ####    ##   ##  ##  ##    ####   
	##  ##   ##  ##            ##  ##    ####    ##  ##   ##  ##            ##  ##   ##  ##   ##  ##   ##   ##  ### ##   ##  ##  
	##       ##  ##            ##  ##   ##  ##   ##       ##  ##            ##  ##   ##  ##   ##  ##   ##   ##  ######   ##      
	## ###   ##  ##            ######   ######    ####    ######            #####    #####    ##  ##   ## # ##  ######    ####   
	##  ##   ##  ##            ##  ##   ##  ##       ##   ##  ##            ##  ##   ####     ##  ##   #######  ## ###       ##  
	##  ##   ##  ##            ##  ##   ##  ##   ##  ##   ##  ##            ##  ##   ## ##    ##  ##   ### ###  ##  ##   ##  ##  
	 ####     ####             ##  ##   ##  ##    ####    ##  ##            #####    ##  ##    ####    ##   ##  ##  ##    ####   
			                                                                                       
 													             [Cybersecurity Utility Tool, hanma-kun] 
	`
	fmt.Println(mnu)
}
