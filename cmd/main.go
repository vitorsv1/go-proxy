package main

import (
	"log"
	"net"
	"net/http"

	//Import proxy configuration
	"github.com/vitorsv1/go-proxy/internal/proxy"
)

func GetLocalIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddress := conn.LocalAddr().(*net.UDPAddr)

	return localAddress.IP
}

func main() {
	//Defining the HTTP port
	port := ":8080"
	log.Printf("Starting proxy server on %s%s...", GetLocalIP(), port)

	http.HandleFunc("/", proxy.HandleRequest)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Could not start server: %s", err)
	}
}
