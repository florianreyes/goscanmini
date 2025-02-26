package pkg

import (
	"flag"
	"fmt"
)

func PrintUsage() {
	fmt.Println("Usage: goscan -host 10.10.10.10 [-p 80,443,3306 | -ports 80,443,3306 | --all] [-max-threads 50]")
	fmt.Println("Flags:")
	flag.PrintDefaults()
}
