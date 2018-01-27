package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const (
	port    = 8080
	version = 2
	timeout = 2 * time.Second
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong\n")
}

func hostHandler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "%s\n", hostname)
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%d\n", version)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", pingHandler)
	mux.HandleFunc("/host", hostHandler)
	mux.HandleFunc("/version", versionHandler)

	s := http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		Handler:           mux,
		ReadTimeout:       timeout,
		WriteTimeout:      timeout,
		IdleTimeout:       timeout,
		ReadHeaderTimeout: timeout,
	}
	s.ListenAndServe()
}
