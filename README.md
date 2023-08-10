# GoHashBrowns

GoHashBrowns is a command-line tool that offers a variety of cryptographic and networking functionalities using the Go programming language. It provides a user-friendly interface to perform tasks such as generating hash values, performing Base64 encoding and decoding, generating secure random numbers, cracking passwords using brute force, and resolving domain names to IP addresses.

## Features

- Generate SHA-256 hash values for input data.
- Perform Base64 encoding and decoding of data.
- Generate cryptographically secure random numbers.
- Brute force password cracker.
- Resolve domain names to IP addresses.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/hanma-kun/GoHashBrowns.git
   ```

2. Navigate to the project directory:

   ```bash
   cd GoHashBrowns
   ```

3. Run the program:

   ```bash
   go run main.go
   ```

## Usage

The tool presents a menu of options to choose from:

1. Generate SHA-256 Hash
   - Input data and get its SHA-256 hash value.

2. Perform Base64 Encoding/Decoding
   - Encode data to Base64 or decode Base64-encoded data.

3. Generate Secure Random Number
   - Generate a cryptographically secure random number.

4. Brute Force Password Cracker
   - Attempt to crack a given password using brute force.

5. Resolve Domain to IP Addresses
   - Enter a domain name to resolve and display associated IP addresses.

## Folder Structure

- `cryptoutils`
  - `encoding`: Base64 encoding and decoding functions.
  - `hashing`: SHA-256 hash generation function.
  - `dnsres`: Domain name resolution to IP addresses.
  - `password`: Brute force password cracker function.
  - `random`: Secure random number generation function.

## Dependencies

This project uses Go modules for dependency management. The required dependencies are automatically managed by the Go module system.

## Contribution

Contributions are welcome! Feel free to open issues or submit pull requests for bug fixes, improvements, or new features.

## License

This project is licensed under the [MIT License](LICENSE).
---
