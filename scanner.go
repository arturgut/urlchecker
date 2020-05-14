// This is URL scanner file
package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"time"
)

func urlScan(url string) {
	start := time.Now()
	resp, err := http.Get(url)
	t := time.Now()
	elapsed := t.Sub(start)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("DEBUG: urlScan() -> resp result:", resp)
	fmt.Println("DEBUG:Duration in ms", elapsed, "Duration as string() is:", string(elapsed))
	// Set map value
	u[url] = ScanResult{resp.StatusCode, elapsed}
	fmt.Println("DEBUG: HTTP RESPONSE CODE:", u[url].responseCode, u[url].durationInMs)
	fmt.Println("DEBUG: elapsed is type of:", reflect.TypeOf(elapsed))
}
