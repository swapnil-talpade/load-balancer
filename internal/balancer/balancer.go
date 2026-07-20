package balancer

import (
	"net/http"

	"github.com/swapnil-talpade/load-balancer/internal/backend"
)

type LoadBalancer struct {
	Backends []*backend.Backend
	Current  int
}

func (lb *LoadBalancer) nextBackend() *backend.Backend {
	backend := lb.Backends[lb.Current]

	lb.Current++

	if lb.Current >= len(lb.Backends) {
		lb.Current = 0
	}

	return backend

}

func (lb *LoadBalancer) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	backend := lb.nextBackend()

	backend.Proxy.ServeHTTP(w, r)

}

func NewLoadBalancer(backends []*backend.Backend) *LoadBalancer {
	return &LoadBalancer{
		Backends: backends,
	}
}
