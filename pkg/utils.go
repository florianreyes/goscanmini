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

func ShowBanner() {
	banner := `
   ___     ___                  __  __ _      _ 
  / __|___/ __| __ __ _ _ _ ___|  \/  (_)_ _ (_)
 | (_ / _ \__ \/ _/ _' | ' \___| |\/| | | ' \| |
  \___\___/___/\__\__,_|_||_|  |_|  |_|_|_||_|_|

	Author: CharlieGauss
	`
	// Print the banner
	fmt.Println(banner)
}
