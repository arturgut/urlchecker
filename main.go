package main

import (
	"fmt"
	"time"
)

// ScanResult - Create scan result struct
type ScanResult struct {
	responseCode int           // also known as status code
	durationInMs time.Duration // measured duration in ms
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

	// fmt.Println("DEBUG: Final map content:")
	// for key, value := range u {
	// 	// fmt.Println("URL:", u[key], "HTTP Code:", u[value].responseCode, "Duration:", u[value].durationInMs)
	// 	fmt.Println(key, value)
	// }

	// Start HTTP server
	startServer(config.Server.Port)

}
