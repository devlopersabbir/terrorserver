package proxy

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReverseProxySetsProductionForwardedHeaders(t *testing.T) {
	var got http.Header
	var gotHost string

	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got = r.Header.Clone()
		gotHost = r.Host
		w.WriteHeader(http.StatusNoContent)
	}))
	defer upstream.Close()

	pool := NewPool()
	rp, err := pool.Get(upstream.URL)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodGet, "http://jenkins.example.com/login", nil)
	req.RemoteAddr = "203.0.113.7:55123"
	rr := httptest.NewRecorder()

	rp.ServeHTTP(rr, req)

	if rr.Code != http.StatusNoContent {
		t.Fatalf("expected upstream status, got %d", rr.Code)
	}
	if gotHost != "jenkins.example.com" {
		t.Fatalf("expected original Host to be preserved, got %q", gotHost)
	}
	if got.Get("X-Forwarded-Host") != "jenkins.example.com" {
		t.Fatalf("expected X-Forwarded-Host, got %q", got.Get("X-Forwarded-Host"))
	}
	if got.Get("X-Forwarded-Proto") != "http" {
		t.Fatalf("expected X-Forwarded-Proto=http, got %q", got.Get("X-Forwarded-Proto"))
	}
	if got.Get("X-Forwarded-Port") != "80" {
		t.Fatalf("expected X-Forwarded-Port=80, got %q", got.Get("X-Forwarded-Port"))
	}
	if got.Get("X-Forwarded-For") != "203.0.113.7" {
		t.Fatalf("expected X-Forwarded-For client IP, got %q", got.Get("X-Forwarded-For"))
	}
	if got.Get("X-Real-IP") != "203.0.113.7" {
		t.Fatalf("expected X-Real-IP client IP, got %q", got.Get("X-Real-IP"))
	}
}

func TestForwardedPortUsesHostPortAndTLS(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "https://jenkins.example.com/", nil)
	req.TLS = &tls.ConnectionState{}

	if got := forwardedPort(req, "jenkins.example.com:8443"); got != "8443" {
		t.Fatalf("expected explicit host port, got %q", got)
	}
	if got := forwardedPort(req, "jenkins.example.com"); got != "443" {
		t.Fatalf("expected TLS default port, got %q", got)
	}
}
