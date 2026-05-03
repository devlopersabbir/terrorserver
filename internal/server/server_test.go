package server

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
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
	if err := os.WriteFile(cfg, []byte(`
:80 {
    root `+root+`
    file_server
}
`), 0o644); err != nil {
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
