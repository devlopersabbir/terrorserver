package main

import (
	"testing"

	"github.com/devlopersabbir/terrorserver/internal/config"
)

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

func TestExpectedListenAddrs(t *testing.T) {
	addrs := expectedListenAddrs(":80", []config.Route{
		{Host: ":9090"},
		{Host: "example.com"},
	})

	want := map[string]bool{":80": true, ":443": true, ":9090": true}
	if len(addrs) != len(want) {
		t.Fatalf("expected %d addrs, got %v", len(want), addrs)
	}
	for _, addr := range addrs {
		if !want[addr] {
			t.Fatalf("unexpected addr %q in %v", addr, addrs)
		}
	}
}
