package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/devlopersabbir/terrorserver/internal/config"
)

func printStatus(cfgPath, addr string) {
	fmt.Println()
	fmt.Println("  terrorserver status")
	fmt.Println("  -------------------------------------")

	if _, err := os.Stat(cfgPath); err != nil {
		fmt.Printf("  x config missing: %s\n", cfgPath)
		fmt.Println("  issue: install or restore Runtime config")
		fmt.Println()
		return
	}
	fmt.Printf("  ok config: %s\n", cfgPath)

	cfg, err := config.Parse(cfgPath)
	if err != nil {
		fmt.Printf("  x config invalid: %v\n", err)
		fmt.Println("  issue: fix Runtime and run 'terror validate'")
		fmt.Println()
		return
	}

	fmt.Printf("  ok listen: %s\n", addr)
	fmt.Printf("  ok routes: %d configured\n", len(cfg.Routes))
	printServiceStatus()
	printRuntimeWatcherStatus()
	printListenerStatus(expectedListenAddrs(addr, cfg.Routes))
	printTLSStatus(cfg.Routes)

	if len(cfg.Routes) == 0 {
		fmt.Println()
		fmt.Println("  issue: no routes found in Runtime")
		fmt.Println()
		return
	}

	fmt.Println()
	fmt.Println("  routes")
	for _, route := range cfg.Routes {
		printRouteStatus(route)
	}
	fmt.Println()
}

func printServiceStatus() {
	if _, err := exec.LookPath("systemctl"); err != nil {
		fmt.Println("  warn service: systemctl unavailable")
		return
	}
	if err := exec.Command("systemctl", "is-active", "--quiet", "terror").Run(); err != nil {
		fmt.Println("  warn service: terror is not active")
		return
	}
	fmt.Println("  ok service: terror is active")
}

func printRuntimeWatcherStatus() {
	if _, err := exec.LookPath("systemctl"); err != nil {
		return
	}
	if err := exec.Command("systemctl", "is-active", "--quiet", "terror.path").Run(); err != nil {
		fmt.Println("  warn watcher: terror.path is not active; Runtime changes may need manual restart")
		return
	}
	fmt.Println("  ok watcher: terror.path is active")
}

func printTLSStatus(routes []config.Route) {
	for _, route := range routes {
		if isDomainRoute(route.Host) {
			host := hostOnly(route.Host)
			if autoTLSDisabled() {
				fmt.Println("  warn ssl: automatic SSL disabled by TERROR_AUTO_TLS")
			} else if canDialLocalPort("443") {
				fmt.Println("  ok ssl: automatic Let's Encrypt SSL enabled")
				if httpsRedirectEnabled() {
					fmt.Println("  ok ssl: domain HTTP redirects to HTTPS")
				} else {
					fmt.Println("  ok ssl: domain HTTP stays available; set TERROR_HTTPS_REDIRECT=true to force HTTPS")
				}
				printHTTPSProbe(host)
			} else {
				fmt.Println("  warn ssl: expected :443 listener is not reachable locally")
			}
			return
		}
	}
}

func printHTTPSProbe(host string) {
	dialer := &net.Dialer{Timeout: 2 * time.Second}
	transport := &http.Transport{
		DialContext: dialer.DialContext,
		TLSClientConfig: &tls.Config{
			ServerName: host,
			MinVersion: tls.VersionTLS12,
		},
	}
	defer transport.CloseIdleConnections()

	client := &http.Client{
		Timeout:   3 * time.Second,
		Transport: transport,
	}
	resp, err := client.Head("https://" + host + "/")
	if err != nil {
		fmt.Printf("  warn ssl: HTTPS probe failed for %s (%v)\n", host, err)
		return
	}
	_ = resp.Body.Close()
	fmt.Printf("  ok ssl: HTTPS probe returned %s\n", resp.Status)
}

func printListenerStatus(addrs []string) {
	for _, addr := range addrs {
		port := portFromListenAddr(addr)
		if port == "" {
			fmt.Printf("  warn listener: cannot inspect %s\n", addr)
			continue
		}
		if canDialLocalPort(port) {
			fmt.Printf("  ok listener: %s is accepting local connections\n", addr)
		} else {
			fmt.Printf("  x listener: %s is not accepting local connections\n", addr)
		}
	}
}

func printRouteStatus(route config.Route) {
	switch route.Type {
	case config.RouteStatic:
		printStaticRouteStatus(route)
	case config.RouteProxy:
		printProxyRouteStatus(route)
	default:
		fmt.Printf("  x %s -> unknown route type\n", route.Host)
	}
}

func printStaticRouteStatus(route config.Route) {
	status := "ok"
	message := "serving static files"

	if _, err := os.Stat(route.Root); err != nil {
		status = "x"
		message = "root missing"
	}

	fmt.Printf("  %s %s -> static %s (%s)\n", status, route.Host, route.Root, message)
	printDomainHint(route.Host)
}

func printProxyRouteStatus(route config.Route) {
	status := "ok"
	message := "upstream reachable"

	target := normalizeDialTarget(route.Target)
	conn, err := net.DialTimeout("tcp", target, 800*time.Millisecond)
	if err != nil {
		status = "warn"
		message = "upstream unreachable"
	} else {
		_ = conn.Close()
	}

	fmt.Printf("  %s %s -> proxy %s (%s)\n", status, route.Host, route.Target, message)
	printDomainHint(route.Host)
}

func printDomainHint(host string) {
	if !isDomainRoute(host) {
		return
	}
	ips, err := net.LookupHost(hostOnly(host))
	if err != nil {
		fmt.Printf("    warn dns: %s does not resolve here yet\n", host)
		return
	}
	fmt.Printf("    ok dns: %s -> %s\n", host, strings.Join(ips, ", "))
}

func normalizeDialTarget(target string) string {
	if strings.HasPrefix(target, "http://") {
		target = strings.TrimPrefix(target, "http://")
	}
	if strings.HasPrefix(target, "https://") {
		target = strings.TrimPrefix(target, "https://")
	}
	target = strings.TrimRight(target, "/")
	if _, _, err := net.SplitHostPort(target); err == nil {
		return target
	}
	return net.JoinHostPort(target, "80")
}

func isDomainRoute(host string) bool {
	h := hostOnly(host)
	return h != "" && !strings.HasPrefix(h, ":") && net.ParseIP(h) == nil
}

func hostOnly(host string) string {
	host = strings.TrimSpace(strings.ToLower(host))
	if strings.HasPrefix(host, ":") {
		return host
	}
	if h, _, err := net.SplitHostPort(host); err == nil {
		return h
	}
	return host
}

func autoTLSDisabled() bool {
	v := strings.ToLower(strings.TrimSpace(os.Getenv("TERROR_AUTO_TLS")))
	return v == "0" || v == "false" || v == "no"
}

func httpsRedirectEnabled() bool {
	v := strings.ToLower(strings.TrimSpace(os.Getenv("TERROR_HTTPS_REDIRECT")))
	return v == "1" || v == "true" || v == "yes"
}

func expectedListenAddrs(defaultAddr string, routes []config.Route) []string {
	seen := map[string]bool{}
	var addrs []string
	add := func(addr string) {
		if addr == "" || seen[addr] {
			return
		}
		seen[addr] = true
		addrs = append(addrs, addr)
	}

	add(defaultAddr)
	if hasDomainRoute(routes) && !autoTLSDisabled() {
		add(":80")
		add(":443")
	}
	for _, route := range routes {
		if strings.HasPrefix(route.Host, ":") {
			add(route.Host)
		}
	}
	return addrs
}

func hasDomainRoute(routes []config.Route) bool {
	for _, route := range routes {
		if isDomainRoute(route.Host) {
			return true
		}
	}
	return false
}

func canDialLocalPort(port string) bool {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort("127.0.0.1", port), 500*time.Millisecond)
	if err != nil {
		return false
	}
	_ = conn.Close()
	return true
}

func portFromListenAddr(addr string) string {
	addr = strings.TrimSpace(addr)
	if strings.HasPrefix(addr, ":") {
		return strings.TrimPrefix(addr, ":")
	}
	if _, port, err := net.SplitHostPort(addr); err == nil {
		return port
	}
	return ""
}
