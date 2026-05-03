package router

import (
	"net"
	"strings"

	"github.com/devlopersabbir/terrorserver/internal/config"
)

type Table struct {
	routes map[string]config.Route
}

type RequestContext struct {
	Host         string
	ListenAddr   string
	ListenerAddr string
}

func NewTable(routes map[string]config.Route) *Table {
	if routes == nil {
		routes = map[string]config.Route{}
	}
	return &Table{routes: routes}
}

func (t *Table) Len() int {
	if t == nil {
		return 0
	}
	return len(t.routes)
}

func Lookup(t *Table, req RequestContext) (config.Route, bool) {
	if t == nil {
		return config.Route{}, false
	}

	host := strings.ToLower(req.Host)
	hostOnly := host
	if h, _, err := net.SplitHostPort(host); err == nil {
		hostOnly = h
	}

	if route, ok := t.routes[host]; ok {
		return route, true
	}
	if route, ok := t.routes[hostOnly]; ok {
		return route, true
	}
	if port := RequestPort(req); port != "" {
		if route, ok := t.routes[":"+port]; ok {
			return route, true
		}
	}
	return config.Route{}, false
}

func RequestPort(req RequestContext) string {
	if _, port, err := net.SplitHostPort(req.Host); err == nil {
		return port
	}
	if port := PortFromAddr(req.ListenAddr); port != "" {
		return port
	}
	return PortFromAddr(req.ListenerAddr)
}

func IsPort80(listenAddr, listenerAddr, host string) bool {
	if PortFromAddr(listenAddr) == "80" {
		return true
	}
	if PortFromAddr(listenerAddr) == "80" {
		return true
	}
	if _, port, err := net.SplitHostPort(host); err == nil {
		return port == "80"
	}
	return false
}

func PortFromAddr(addr string) string {
	if addr == "" {
		return ""
	}
	if strings.HasPrefix(addr, ":") {
		return strings.TrimPrefix(addr, ":")
	}
	if _, port, err := net.SplitHostPort(addr); err == nil {
		return port
	}
	return ""
}
