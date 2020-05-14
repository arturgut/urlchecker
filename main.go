package main

import "fmt"

// ScanResult - Create scan result struct
type ScanResult struct {
	responseCode int   // also known as HTTP status code
	durationInMs int64 // measured duration in ms
}

var u = make(map[string]ScanResult)

func main() {

	printVersion()                   // Print version
	loadConfiguration("config.yaml") // load configuration

	fmt.Println("INFO: Starting URL scanner")
	for _, url := range config.Urls { // Main loop. For each item scan URL.
		urlScan(url)
	}

	// Start HTTP server
	startServer(config.Server.Port)

}
