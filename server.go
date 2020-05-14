package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func metrics(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "# HELP Version: 0.1 Alpha - https://github.com/arturgut/urlchecker\n")
	fmt.Fprintf(w, "# HELP url_checker. Label: HTTP Response code. Value: Request duration in Ms\n")
	for key, value := range u {
		// Prometheus format: <metric name>{<label name>=<label value>, ...}
		fmt.Fprintf(w, "url_checker{ url='%v', http_status_code=%v } %v\n", key, value.responseCode, value.durationInMs)
	}
}

func startServer() {
	fmt.Println("INFO: Starting HTTP server on port:", config.Server.Port, ".\n You should be able to access the URL at http://localhost:", config.Server.Port)
	http.HandleFunc("/metrics", metrics)
	serverPort := ":" + strconv.FormatInt(int64(config.Server.Port), 10)
	fmt.Println("Server port:", serverPort)
	http.ListenAndServe(serverPort, nil)
}
