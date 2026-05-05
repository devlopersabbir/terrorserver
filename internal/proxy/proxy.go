package proxy

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"sync"
)

// Pool caches ReverseProxy instances keyed by target address.
type Pool struct {
	mu    sync.RWMutex
	cache map[string]*httputil.ReverseProxy
}

func NewPool() *Pool {
	return &Pool{cache: make(map[string]*httputil.ReverseProxy)}
}

// Get returns a cached or newly created ReverseProxy for the given target.
// target is a bare host:port (e.g. "localhost:5000").
func (p *Pool) Get(target string) (*httputil.ReverseProxy, error) {
	p.mu.RLock()
	rp, ok := p.cache[target]
	p.mu.RUnlock()
	if ok {
		return rp, nil
	}

	addr := target
	if !strings.Contains(addr, "://") {
		addr = "http://" + addr
	}
	u, err := url.Parse(addr)
	if err != nil {
		return nil, fmt.Errorf("invalid proxy target %q: %w", target, err)
	}

	rp = httputil.NewSingleHostReverseProxy(u)

	// Custom error handler to return clean 502
	rp.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		http.Error(w, "Bad Gateway", http.StatusBadGateway)
	}

	// Preserve original Host header by default (pass-through mode)
	origDirector := rp.Director
	rp.Director = func(req *http.Request) {
		originalHost := req.Host
		originalRemoteAddr := req.RemoteAddr
		origDirector(req)
		req.Host = originalHost
		req.Header.Set("X-Forwarded-Host", originalHost)
		req.Header.Set("X-Forwarded-Proto", forwardedProto(req))
		req.Header.Set("X-Forwarded-Port", forwardedPort(req, originalHost))
		req.Header.Set("X-Real-IP", realIP(req, originalRemoteAddr))
		if _, ok := req.Header["User-Agent"]; !ok {
			req.Header.Set("User-Agent", "")
		}
	}

	p.mu.Lock()
	p.cache[target] = rp
	p.mu.Unlock()

	return rp, nil
}

// Flush clears all cached proxies (called on config reload).
func (p *Pool) Flush() {
	p.mu.Lock()
	p.cache = make(map[string]*httputil.ReverseProxy)
	p.mu.Unlock()
}

func forwardedProto(r *http.Request) string {
	if r.TLS != nil {
		return "https"
	}
	return "http"
}

func forwardedPort(r *http.Request, host string) string {
	if _, port, err := net.SplitHostPort(host); err == nil {
		return port
	}
	if r.TLS != nil {
		return "443"
	}
	return "80"
}

func realIP(r *http.Request, remoteAddr string) string {
	if ip := r.Header.Get("X-Forwarded-For"); ip != "" {
		return strings.Split(ip, ",")[0]
	}
	if ip := r.Header.Get("X-Real-IP"); ip != "" {
		return ip
	}
	return clientIP(remoteAddr)
}

func clientIP(remoteAddr string) string {
	host, _, err := net.SplitHostPort(remoteAddr)
	if err == nil {
		return host
	}
	return remoteAddr
}
