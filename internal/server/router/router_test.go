package router

import (
	"testing"

	"github.com/devlopersabbir/terrorserver/internal/config"
)

func TestPortOnlyRouteMatchesHostWithoutPort(t *testing.T) {
	table := NewTable(map[string]config.Route{
		":80": {Host: ":80", Type: config.RouteStatic, Root: "/srv/www"},
	})

	route, ok := Lookup(table, RequestContext{
		Host:       "203.0.113.10",
		ListenAddr: ":80",
	})
	if !ok {
		t.Fatal("expected port-only route to match host without explicit port")
	}
	if route.Root != "/srv/www" {
		t.Fatalf("wrong route matched: %#v", route)
	}
}

func TestDomainRouteMatchesHostWithPort(t *testing.T) {
	table := NewTable(map[string]config.Route{
		"app.example.com": {Host: "app.example.com", Type: config.RouteProxy, Target: "localhost:3000"},
	})

	route, ok := Lookup(table, RequestContext{Host: "app.example.com:80"})
	if !ok {
		t.Fatal("expected domain route to match host with port")
	}
	if route.Target != "localhost:3000" {
		t.Fatalf("wrong route matched: %#v", route)
	}
}
