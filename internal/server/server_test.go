package server

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/devlopersabbir/terrorserver/internal/config"
)

func TestWelcomePageOnPort80(t *testing.T) {
	s := New("")
	s.listenAddr = ":80"

	req := httptest.NewRequest(http.MethodGet, "http://example.com/", nil)
	rr := httptest.NewRecorder()

	s.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
	if ct := rr.Header().Get("Content-Type"); ct != "text/html; charset=utf-8" {
		t.Fatalf("expected HTML content type, got %q", ct)
	}
	if !strings.Contains(rr.Body.String(), "Terror Server is running") {
		t.Fatal("expected welcome page body")
	}
}

func TestUnmatchedRouteStill404AwayFromPort80(t *testing.T) {
	s := New("")
	s.listenAddr = ":8080"

	req := httptest.NewRequest(http.MethodGet, "http://example.com/", nil)
	rr := httptest.NewRecorder()

	s.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", rr.Code)
	}
}

func TestPortOnlyRouteMatchesHostWithoutPort(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "index.html"), []byte("static welcome"), 0o644); err != nil {
		t.Fatal(err)
	}

	cfg := filepath.Join(t.TempDir(), "Runtime")
	runtime := `
:80 {
    root ` + root + `
    file_server
}
`
	if err := os.WriteFile(cfg, []byte(runtime), 0o644); err != nil {
		t.Fatal(err)
	}

	s := New(cfg)
	s.listenAddr = ":80"
	if err := s.LoadConfig(); err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodGet, "http://203.0.113.10/", nil)
	rr := httptest.NewRecorder()

	s.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %q", rr.Code, rr.Body.String())
	}
	if strings.TrimSpace(rr.Body.String()) != "static welcome" {
		t.Fatalf("expected static welcome page, got %q", rr.Body.String())
	}
}

func TestDomainHTTPRedirectsToHTTPS(t *testing.T) {
	t.Setenv("TERROR_HTTPS_REDIRECT", "true")

	cfg := filepath.Join(t.TempDir(), "Runtime")
	runtime := `
example.com {
    proxy localhost:4000
}
`
	if err := os.WriteFile(cfg, []byte(runtime), 0o644); err != nil {
		t.Fatal(err)
	}

	s := New(cfg)
	s.listenAddr = ":80"
	if err := s.LoadConfig(); err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodGet, "http://example.com/dashboard", nil)
	rr := httptest.NewRecorder()

	s.httpsRedirectHandler().ServeHTTP(rr, req)

	if rr.Code != http.StatusMovedPermanently {
		t.Fatalf("expected HTTPS redirect, got %d", rr.Code)
	}
	if got := rr.Header().Get("Location"); got != "https://example.com/dashboard" {
		t.Fatalf("expected HTTPS location, got %q", got)
	}
}

func TestPortHTTPDoesNotRedirectToHTTPS(t *testing.T) {
	t.Setenv("TERROR_HTTPS_REDIRECT", "true")

	cfg := filepath.Join(t.TempDir(), "Runtime")
	runtime := `
:9090 {
    proxy localhost:4000
}
`
	if err := os.WriteFile(cfg, []byte(runtime), 0o644); err != nil {
		t.Fatal(err)
	}

	s := New(cfg)
	s.listenAddr = ":9090"
	if err := s.LoadConfig(); err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodGet, "http://127.0.0.1:9090/", nil)
	rr := httptest.NewRecorder()

	s.httpsRedirectHandler().ServeHTTP(rr, req)

	if rr.Code == http.StatusMovedPermanently {
		t.Fatal("did not expect port route to redirect to HTTPS")
	}
}

func TestHTTPSRedirectDisabledByDefault(t *testing.T) {
	t.Setenv("TERROR_HTTPS_REDIRECT", "")

	if httpsRedirectEnabled() {
		t.Fatal("expected HTTPS redirects to be disabled by default")
	}
}

func TestListenAddrsIncludesPortRoutes(t *testing.T) {
	got := listenAddrs(":80", []config.Route{
		{Host: ":9090", Type: config.RouteProxy, Target: "localhost:4000"},
		{Host: "example.com", Type: config.RouteProxy, Target: "localhost:4000"},
	}, false)

	want := map[string]bool{":80": true, ":9090": true}
	if len(got) != len(want) {
		t.Fatalf("expected %d listen addrs, got %v", len(want), got)
	}
	for _, addr := range got {
		if !want[addr] {
			t.Fatalf("unexpected listen addr %q in %v", addr, got)
		}
	}
}

func TestListenAddrsIncludesPort80ForACME(t *testing.T) {
	got := listenAddrs(":8080", nil, true)
	want := map[string]bool{":80": true, ":8080": true}

	if len(got) != len(want) {
		t.Fatalf("expected %d listen addrs, got %v", len(want), got)
	}
	for _, addr := range got {
		if !want[addr] {
			t.Fatalf("unexpected listen addr %q in %v", addr, got)
		}
	}
}

func TestDomainRoutesFiltersPortAndIPRoutes(t *testing.T) {
	got := domainRoutes([]config.Route{
		{Host: ":9090"},
		{Host: "127.0.0.1"},
		{Host: "example.com"},
		{Host: "example.com:80"},
	})

	if len(got) != 1 || got[0] != "example.com" {
		t.Fatalf("expected only example.com, got %v", got)
	}
}
