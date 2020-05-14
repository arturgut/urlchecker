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
	elapsed := int64(t.Sub(start)) / 1000 / 1000 // Convert to miliseconds
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("DEBUG: urlScan() -> resp result:", resp)
	fmt.Println("DEBUG:Duration in ms", elapsed, "Duration as string() is:", elapsed)
	// Set map value
	u[url] = ScanResult{resp.StatusCode, elapsed}
	fmt.Println("DEBUG: elapsed ", elapsed, "which is typeOf", reflect.TypeOf(elapsed))
}
