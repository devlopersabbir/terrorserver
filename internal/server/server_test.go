package server

import (
	"net/http"
	"net/http/httptest"
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
