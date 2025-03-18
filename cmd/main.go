package main

import (
	"log"
	"net/http"

	//Import proxy configuration
	"github.com/vitorsv1/go-proxy/internal/proxy"
)

func main() {
	//Defining the HTTP port
	port := ":8080"
	log.Printf("Starting proxy server on %s...", port)

	http.HandleFunc("/", proxy.HandleRequest)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Could not start server: %s", err)
	}
}
