package main

// ScanResult - Create scan result struct
type ScanResult struct {
	responseCode int // also known as HTTP status code
	durationInMs int // measured duration in ms
}

var u = make(map[string]ScanResult) // Create a map of string => ScanResult
var c = make(chan string)           // Initiate go routine channel. This needs to be globally accessible

func main() {

	initLogging()
	printVersion()                   // Print version
	loadConfiguration("config.yaml") // load configuration
	readEnv(&config)                 // Read environment attributes
	go scannerLoop()                 // Start URL scanner routines
	startServer()                    // Start HTTP server
}
