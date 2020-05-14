package main

import (
	"fmt"
	"net/http"
)

func metrics(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "# HELP Version: 0.1 Alpha - https://github.com/arturgut/urlchecker\n")
	fmt.Fprintf(w, "# HELP url_checker. Label: HTTP Response code. Value: Request duration in Ms\n")
	for key, value := range u {
		// Prometheus format: <metric name>{<label name>=<label value>, ...}
		fmt.Fprintf(w, "url_checker{ url='%v', http_status_code=%v } %v\n", key, value.responseCode, value.durationInMs)
	}
}

func startServer(port int) {
	fmt.Println("INFO: Starting HTTP server on port 8090.\n You should be able to access the URL at http://localhost:8090")
	http.HandleFunc("/metrics", metrics)
	http.ListenAndServe(":8090", nil)
}
