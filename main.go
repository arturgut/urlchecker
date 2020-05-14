package main

import (
	log "github.com/sirupsen/logrus"
)

// ScanResult - Create scan result struct
type ScanResult struct {
	responseCode int // also known as HTTP status code
	durationInMs int // measured duration in ms
}

var u = make(map[string]ScanResult) // Create a map of string => ScanResult

func main() {

	initLogging()

	log.Info("Starting URL checker. Version: 0.1")

	printVersion()                   // Print version
	loadConfiguration("config.yaml") // load configuration
	readEnv(&config)                 // Read environment attributes
	go scannerLoop()                 // Start URL scanner routines
	startServer()                    // Start HTTP server
}
