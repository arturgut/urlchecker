package main

// ScanResult - Create scan result struct
type ScanResult struct {
	responseCode int // also known as HTTP status code
	durationInMs int // measured duration in ms
}

var u = make(map[string]ScanResult) // Create a map of string => ScanResult

func main() {

	printVersion()                   // Print version
	loadConfiguration("config.yaml") // load configuration
	go scannerLoop()                 // Start URL scanner routines
	startServer()                    // Start HTTP server
}
