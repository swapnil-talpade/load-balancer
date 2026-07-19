package main

import (
	"log"
	"net/http"
	"net/url"

	"github.com/swapnil-talpade/load-balancer/internal/proxy"
)

func main() {
	// the backend we want to forward requests to
	backendURL, err := url.Parse("http://localhost:9001")
	if err != nil {
		log.Fatal(err)
	}

	// build the reverse proxy for that backend
	reverseProxy := proxy.NewReverseProxy(backendURL)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", reverseProxy))
}
