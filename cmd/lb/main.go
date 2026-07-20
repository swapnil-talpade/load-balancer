package main

import (
	"log"
	"net/http"
	"net/url"

	"github.com/swapnil-talpade/load-balancer/internal/backend"
	"github.com/swapnil-talpade/load-balancer/internal/balancer"
	"github.com/swapnil-talpade/load-balancer/internal/proxy"
)

func main() {
	backendURLs := []string{
		"http://localhost:9001",
		"http://localhost:9002",
	}

	var backends []*backend.Backend

	for _, rawURLs := range backendURLs {
		backendURL, err := url.Parse(rawURLs)
		if err != nil {
			log.Fatal(err)
		}

		reverseProxy := proxy.NewReverseProxy(backendURL)

		backends = append(backends, &backend.Backend{
			URL:   backendURL,
			Proxy: reverseProxy,
		})
	}

	loadBalancer := balancer.NewLoadBalancer(backends)

	log.Println("Load Balancer listening on :8080")

	log.Fatal(http.ListenAndServe(":8080", loadBalancer))

}
