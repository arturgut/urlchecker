package main

import (
	"fmt"
)

// ScanResult - Create scan result struct
type ScanResult struct {
	responseCode int   // also known as status code
	durationInMs int64 // measured duration in ms
}

var u = make(map[string]ScanResult)

func main() {

	printVersion()                   // Print version
	loadConfiguration("config.yaml") // load configuration

	// Main loop. For each item scan URL.
	for _, url := range config.Urls {
		fmt.Println("\nDEBUG: main() => loop through url", url)
		// Scan URL's
		urlScan(url)
	}

	// Start HTTP server
	startServer(config.Server.Port)

}
