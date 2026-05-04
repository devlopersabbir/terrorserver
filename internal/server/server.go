package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"os"
	"sort"
	"strings"
	"sync/atomic"
	"time"

	"github.com/devlopersabbir/terrorserver/internal/config"
	"github.com/devlopersabbir/terrorserver/internal/logger"
	"github.com/devlopersabbir/terrorserver/internal/proxy"
	"github.com/devlopersabbir/terrorserver/internal/server/handler"
	"github.com/devlopersabbir/terrorserver/internal/server/response"
	"github.com/devlopersabbir/terrorserver/internal/server/router"
	"golang.org/x/crypto/acme/autocert"
)

// Server is the core terrorserver HTTP engine.
type Server struct {
	configPath  string
	listenAddr  string
	tablePtr    atomic.Pointer[router.Table]
	proxyPool   *proxy.Pool
	httpServers []*http.Server
	listeners   []net.Listener

	// status fields
	startedAt   time.Time
	configLoads int64
}

// New creates a Server.
func New(cfgPath string) *Server {
	s := &Server{
		configPath: cfgPath,
		proxyPool:  proxy.NewPool(),
	}
	return s
}

// LoadConfig parses the config file and atomically swaps the routing table.
// Returns an error if parsing fails; on error the OLD table remains in effect.
func (s *Server) LoadConfig() error {
	cfg, err := config.Parse(s.configPath)
	if err != nil {
		return err
	}
	tbl := router.NewTable(cfg.RouteMap())
	s.tablePtr.Store(tbl)
	atomic.AddInt64(&s.configLoads, 1)

	// Flush proxy cache so old upstreams aren't reused after a target change
	s.proxyPool.Flush()

	logger.Info("loaded %d route(s)", tbl.Len())
	for _, r := range cfg.Routes {
		switch r.Type {
		case config.RouteProxy:
			logger.Info("  %s → proxy %s", r.Host, r.Target)
		case config.RouteStatic:
			logger.Info("  %s → static %s", r.Host, r.Root)
		}
	}
	return nil
}

// Start binds the server and begins serving.
func (s *Server) Start(addr string) error {
	s.startedAt = time.Now()
	s.listenAddr = addr

	cfg, err := config.Parse(s.configPath)
	if err != nil {
		return fmt.Errorf("cannot read listener config: %w", err)
	}

	domains := domainRoutes(cfg.Routes)
	httpHandler := http.Handler(s)
	if len(domains) > 0 && autoTLSDisabled() {
		logger.Warn("automatic SSL disabled by TERROR_AUTO_TLS=false")
	} else if len(domains) > 0 {
		manager := &autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			Cache:      autocert.DirCache(certCacheDir()),
			HostPolicy: autocert.HostWhitelist(domains...),
		}
		if httpsRedirectEnabled() {
			httpHandler = manager.HTTPHandler(s.httpsRedirectHandler())
		} else {
			httpHandler = manager.HTTPHandler(s)
		}
		if err := s.startListener(":443", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			s.ServeHTTP(w, r)
		}), &tls.Config{GetCertificate: manager.GetCertificate, MinVersion: tls.VersionTLS12}); err != nil {
			return err
		}
		logger.Info("automatic SSL enabled for: %s", strings.Join(domains, ", "))
		if httpsRedirectEnabled() {
			logger.Info("domain HTTP to HTTPS redirects enabled")
		}
	}

	for _, listen := range listenAddrs(addr, cfg.Routes, len(domains) > 0 && !autoTLSDisabled()) {
		if err := s.startListener(listen, httpHandler, nil); err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) httpsRedirectHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if s.shouldRedirectToHTTPS(r) {
			target := "https://" + hostOnly(r.Host) + r.URL.RequestURI()
			http.Redirect(w, r, target, http.StatusMovedPermanently)
			return
		}
		s.ServeHTTP(w, r)
	})
}

func (s *Server) shouldRedirectToHTTPS(r *http.Request) bool {
	if r.TLS != nil || r.Method != http.MethodGet && r.Method != http.MethodHead {
		return false
	}
	host := hostOnly(r.Host)
	if host == "" || net.ParseIP(host) != nil || strings.HasPrefix(host, ":") {
		return false
	}
	route, ok := router.Lookup(s.table(), router.RequestContext{
		Host:         r.Host,
		ListenAddr:   s.listenAddr,
		ListenerAddr: listenerAddr(r),
	})
	return ok && route.Host != "" && !strings.HasPrefix(route.Host, ":")
}

// Shutdown gracefully stops the HTTP server.
func (s *Server) Shutdown(ctx context.Context) error {
	var err error
	for _, srv := range s.httpServers {
		if shutdownErr := srv.Shutdown(ctx); shutdownErr != nil && err == nil {
			err = shutdownErr
		}
	}
	return err
}

// ServeHTTP is the main request dispatcher.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	rw := response.NewRecorder(w)

	route, ok := router.Lookup(s.table(), router.RequestContext{
		Host:         r.Host,
		ListenAddr:   s.listenAddr,
		ListenerAddr: listenerAddr(r),
	})
	if !ok {
		if router.IsPort80(s.listenAddr, listenerAddr(r), r.Host) {
			handler.Welcome(rw)
			logger.Request(r.Method, r.Host, r.URL.Path, rw.Code, time.Since(start))
			return
		}
		http.Error(rw, "Not Found", http.StatusNotFound)
		logger.Request(r.Method, r.Host, r.URL.Path, rw.Code, time.Since(start))
		return
	}

	switch route.Type {
	case config.RouteProxy:
		handler.Proxy(rw, r, s.proxyPool, route)
	case config.RouteStatic:
		handler.Static(rw, r, route)
	default:
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
	}

	logger.Request(r.Method, r.Host, r.URL.Path, rw.Code, time.Since(start))
}

func (s *Server) table() *router.Table {
	if tbl := s.tablePtr.Load(); tbl != nil {
		return tbl
	}
	return router.NewTable(nil)
}

func (s *Server) startListener(addr string, h http.Handler, tlsConfig *tls.Config) error {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("cannot bind %s: %w", addr, err)
	}

	srv := &http.Server{
		Handler:      h,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  120 * time.Second,
		TLSConfig:    tlsConfig,
	}

	s.listeners = append(s.listeners, ln)
	s.httpServers = append(s.httpServers, srv)

	logger.Info("listening on %s", addr)
	go func() {
		var serveErr error
		if tlsConfig != nil {
			serveErr = srv.ServeTLS(ln, "", "")
		} else {
			serveErr = srv.Serve(ln)
		}
		if serveErr != nil && serveErr != http.ErrServerClosed {
			logger.Error("listener %s stopped: %v", addr, serveErr)
		}
	}()
	return nil
}

func listenerAddr(r *http.Request) string {
	if addr, ok := r.Context().Value(http.LocalAddrContextKey).(net.Addr); ok && addr != nil {
		return addr.String()
	}
	return ""
}

func listenAddrs(defaultAddr string, routes []config.Route, needsACMEHTTP bool) []string {
	seen := map[string]bool{}
	addrs := []string{}
	add := func(addr string) {
		if addr == "" || seen[addr] {
			return
		}
		seen[addr] = true
		addrs = append(addrs, addr)
	}

	add(defaultAddr)
	if needsACMEHTTP {
		add(":80")
	}
	for _, route := range routes {
		if strings.HasPrefix(route.Host, ":") {
			add(route.Host)
		}
	}
	sort.Strings(addrs)
	return addrs
}

func domainRoutes(routes []config.Route) []string {
	seen := map[string]bool{}
	var domains []string
	for _, route := range routes {
		host := strings.ToLower(strings.TrimSpace(route.Host))
		if host == "" || strings.HasPrefix(host, ":") {
			continue
		}
		if h, _, err := net.SplitHostPort(host); err == nil {
			host = h
		}
		if net.ParseIP(host) != nil || seen[host] {
			continue
		}
		seen[host] = true
		domains = append(domains, host)
	}
	sort.Strings(domains)
	return domains
}

func hostOnly(host string) string {
	host = strings.ToLower(strings.TrimSpace(host))
	if strings.HasPrefix(host, ":") {
		return host
	}
	if h, _, err := net.SplitHostPort(host); err == nil {
		return h
	}
	return host
}

func certCacheDir() string {
	if dir := os.Getenv("TERROR_CERT_CACHE"); dir != "" {
		return dir
	}
	return "/var/lib/terror/certs"
}

func autoTLSDisabled() bool {
	v := strings.ToLower(strings.TrimSpace(os.Getenv("TERROR_AUTO_TLS")))
	return v == "0" || v == "false" || v == "no"
}

func httpsRedirectEnabled() bool {
	v := strings.ToLower(strings.TrimSpace(os.Getenv("TERROR_HTTPS_REDIRECT")))
	return v == "1" || v == "true" || v == "yes"
}
