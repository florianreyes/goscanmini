package main

import (
	"flag"
	"fmt"
	pkg "goscanmini/pkg"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/atotto/clipboard"
)

func main() {
	// Define flags
	hostFlag := flag.String("host", "", "The host to scan (e.g. 10.10.10.17)")
	portFlag := flag.String("p", "", "Comma-Separated values of ports to scan (e.g. 80,443,3306)")
	flag.StringVar(portFlag, "ports", "", "Comma-Separated values of ports to scan (e.g. 80,443,3306)")
	maxThreadFlag := flag.Int("max-threads", 10, "Max amount of threads to spawn")
	allPortsFlag := flag.Bool("all", false, "Scan all 65535 ports")

	flag.Parse()

	// Validate host
	if *hostFlag == "" {
		fmt.Println("[!] Error: -host flag required")
		pkg.PrintUsage() // Corrected: Use pkg.PrintUsage()
		os.Exit(1)
	}

	// Ensure max threads is positive
	if *maxThreadFlag <= 0 {
		fmt.Println("[!] Error: max-threads must be a positive number")
		os.Exit(1)
	}

	// Determine ports to scan
	var validPorts []int

	if *allPortsFlag {
		// Scan all 65535 ports
		fmt.Println("[*] Scanning all 65535 ports...")
		for port := 1; port <= 65535; port++ {
			validPorts = append(validPorts, port)
		}
	} else if *portFlag == "" {
		// Default to top 100 ports
		fmt.Println("[*] No ports supplied, defaulting to top 100 ports...")
		validPorts = pkg.Top100Ports // Corrected: Use pkg.Top100Ports
	} else {
		// Parse the provided ports
		portList := strings.Split(*portFlag, ",")
		for _, port := range portList {
			port = strings.TrimSpace(port)
			portNumber, err := strconv.Atoi(port)
			if err != nil {
				fmt.Printf("[!] Error: %s is not a valid number.\n", port)
				os.Exit(1)
			}

			if portNumber < 1 || portNumber > 65535 {
				fmt.Printf("[!] Error: %d is not a valid port [1-65535].\n", portNumber)
				os.Exit(1)
			}

			validPorts = append(validPorts, portNumber)
		}
	}

	// Concurrency
	var wg sync.WaitGroup
	var mu sync.Mutex
	semaphore := make(chan struct{}, *maxThreadFlag)
	var openPorts []int

	for _, port := range validPorts {
		wg.Add(1)
		semaphore <- struct{}{}

		go func(port int) {
			defer func() { <-semaphore }()
			pkg.ScanPort(*hostFlag, port, &wg, &openPorts, &mu) // Corrected: Use pkg.ScanPort()
		}(port)
	}

	wg.Wait()

	// Copy open ports to clipboard
	if len(openPorts) > 0 {
		portListStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(openPorts)), ","), "[]")
		clipboard.WriteAll(portListStr)
		fmt.Printf("[*] Open ports copied to clipboard: %s\n", portListStr)
	} else {
		fmt.Println("[!] No open ports found.")
	}
}
