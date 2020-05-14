package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

// This handler does something a little more sophisticated
// by reading all the HTTP request headers and echoing
// them into the response body.
func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func metrics(w http.ResponseWriter, req *http.Request) {
	for key, value := range u {
		// Prometheus format:
		// <metric name>{<label name>=<label value>, ...}
		// http_requests_total{service="service", server="pod50", env="production"}
		fmt.Fprintf(w, "url_checker{ %v: http_status_code=%v, request_duration=%v}\n", key, value.responseCode, value.durationInMs)
	}
}

func version(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Version: 0.1 ")
}

func startServer(port int) {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/metrics", metrics)
	http.HandleFunc("/version", version)
	http.ListenAndServe(":8090", nil)
}
