package backend

import (
	"net/url"

	"github.com/swapnil-talpade/load-balancer/internal/proxy"
)

type Backend struct {
	URL         *url.URL
	Proxy       *proxy.ReverseProxy
	Alive       bool
	Connections int
}
