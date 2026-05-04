package main

import "testing"

func TestNormalizeDialTarget(t *testing.T) {
	tests := map[string]string{
		"localhost:3000":        "localhost:3000",
		"http://localhost:3000": "localhost:3000",
		"example.com":           "example.com:80",
	}

	for input, want := range tests {
		if got := normalizeDialTarget(input); got != want {
			t.Fatalf("normalizeDialTarget(%q) = %q, want %q", input, got, want)
		}
	}
}

func TestIsDomainRoute(t *testing.T) {
	if !isDomainRoute("example.com") {
		t.Fatal("expected example.com to be a domain route")
	}
	if isDomainRoute(":80") {
		t.Fatal("expected :80 not to be a domain route")
	}
	if isDomainRoute("127.0.0.1") {
		t.Fatal("expected IP not to be a domain route")
	}
}
