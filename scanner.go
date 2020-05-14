// This is URL scanner file
package main

import (
	"fmt"
	"net/http"
	"time"
)

func scannerLoop() {

	fmt.Println("INFO: Starting URL scanner")

	c := make(chan string)

	for _, url := range config.Urls { // Main loop. For each item initiate go routine
		go urlScan(url, c)
	}

	for l := range c { // Infinite for loop to continusly scan URL listed in the channel 'c'
		go func(url string) {
			time.Sleep(config.Client.Period * time.Second)
			urlScan(url, c)
		}(l)
	}
}

func urlScan(url string, c chan string) { // Main urlScan function
	fmt.Println("INFO: Starting URL scan", url)
	start := time.Now() // Measure time taken to establish http connection

	client := &http.Client{
		Timeout: config.Client.Timeout * time.Second,
	}

	resp, err := client.Get(url)
	t := time.Now()                            // Measure time taken to establish http connection
	elapsed := int(t.Sub(start)) / 1000 / 1000 // Convert to from date.tinme to miliseconds
	if err != nil {
		fmt.Println("Error:", err)
		c <- url // return result of scan the channel
		return
	}
	// fmt.Println("DEBUG: urlScan() -> resp result:", resp, "DEBUG:Duration in ms", elapsed, "Duration as string() is:", elapsed)
	// Add URL to global 'u' map
	u[url] = ScanResult{resp.StatusCode, elapsed}
	c <- url // return result of scan the channel
}
