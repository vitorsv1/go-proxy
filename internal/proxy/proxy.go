package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// HandleRequest processes incoming HTTP requests and forwards them
func HandleRequest(res http.ResponseWriter, req *http.Request) {
	log.Printf("Received request: %s %s", req.Method, req.URL)

	// Parse the target URL
	targetURL, err := url.Parse(req.URL.String())
	if err != nil {
		http.Error(res, "Bad Request", http.StatusBadRequest)
		return
	}

	// Create a reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	proxy.ServeHTTP(res, req)
}
