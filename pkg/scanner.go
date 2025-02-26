package pkg

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

// Top 100 common ports
var Top100Ports = []int{
	7, 20, 21, 22, 23, 25, 26, 37, 53, 79, 80, 81, 88, 110, 111, 113, 119, 123, 135, 139,
	143, 161, 162, 179, 199, 389, 443, 445, 465, 500, 514, 515, 520, 523, 524, 530, 543,
	548, 554, 587, 626, 631, 636, 646, 648, 666, 669, 683, 687, 691, 700, 705, 711, 721,
	749, 750, 751, 765, 777, 783, 808, 843, 873, 902, 981, 987, 990, 992, 993, 995, 1000,
	1010, 1024, 1025, 1080, 1088, 1099, 1194, 1214, 1241, 1311, 1337, 1433, 1434, 1512,
	1524, 1723, 1755, 1801, 1900, 2049, 2100, 2222, 2301, 2381, 3128, 3306, 3389, 3689,
	4444, 4899, 5000, 5051, 5060, 5190, 5353, 5432, 5631, 5900, 6000, 6667, 8000, 8080,
}

// Scan a port using a 1-second timeout
func ScanPort(host string, port int, wg *sync.WaitGroup, openPorts *[]int, mu *sync.Mutex) {
	defer wg.Done()
	timeout := 1 * time.Second
	address := net.JoinHostPort(host, strconv.Itoa(port)) // Supports IPv6

	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return // Closed port
	}
	defer conn.Close()

	// Add open port to list
	mu.Lock()
	*openPorts = append(*openPorts, port)
	mu.Unlock()

	fmt.Printf("[+] Port %d is open\n", port)
}
