package portscanner

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type ScanOptions struct {
	Host       string
	StartPort  int
	EndPort    int
	NumWorkers int
}

type ScanResult struct {
	Port  int
	Open  bool
	Error error
}

func ScanPorts(options ScanOptions) []ScanResult {
	var results []ScanResult
	var wg sync.WaitGroup

	ports := make([]int, 0, options.EndPort-options.StartPort+1)
	for port := options.StartPort; port <= options.EndPort; port++ {
		ports = append(ports, port)
	}

	portCh := make(chan int, len(ports))
	for _, port := range ports {
		portCh <- port
	}
	close(portCh)

	resultsCh := make(chan ScanResult, len(ports))

	for i := 0; i < options.NumWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for port := range portCh {
				result := scanPort(options.Host, port)
				resultsCh <- result
			}
		}()
	}

	go func() {
		wg.Wait()
		close(resultsCh)
	}()

	for result := range resultsCh {
		results = append(results, result)
	}

	return results
}

func scanPort(host string, port int) ScanResult {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, 1*time.Second)
	if err != nil {
		return ScanResult{Port: port, Open: false, Error: err}
	}
	conn.Close()
	return ScanResult{Port: port, Open: true, Error: nil}
}

func parsePort(input string) int {
	port, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid port number.")
		os.Exit(1)
	}
	return port
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
	return input
}

func getScanOptions() ScanOptions {
	var options ScanOptions

	fmt.Print("Enter the target host (IP address or domain name): ")
	options.Host = readInput()

	fmt.Print("Enter the starting port: ")
	options.StartPort = parsePort(readInput())

	fmt.Print("Enter the ending port: ")
	options.EndPort = parsePort(readInput())

	fmt.Print("Enter the number of workers: ")
	options.NumWorkers = parsePort(readInput())

	return options
}

func PortScan() {
	fmt.Println("\nPort Scanning:")
	options := getScanOptions()
	results := ScanPorts(options)

	fmt.Println("\nPort Scanning Results:")
	for _, result := range results {
		openStr := "Closed"
		if result.Open {
			openStr = "Open"
		}
		fmt.Printf("Port %d: %s\n", result.Port, openStr)
	}
}