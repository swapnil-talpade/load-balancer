package proxy

import (
	"io"
	"net/http"
	"net/url"
)

type ReverseProxy struct {
	BackendURL *url.URL
	Client     *http.Client
}

func NewReverseProxy(backendURL *url.URL) *ReverseProxy {
	return &ReverseProxy{
		BackendURL: backendURL,
		Client:     http.DefaultClient,
	}
}

func (p *ReverseProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	targetURL := p.BackendURL.ResolveReference(r.URL)

	req, err := http.NewRequest(r.Method, targetURL.String(), r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// forward request headers
	for key, values := range r.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	req.Header.Set("X-Forwarded-By", "go-load-balancer")

	// send request to the backend
	resp, err := p.Client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// forward response headers
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	// forward status code, then the body
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)

}
