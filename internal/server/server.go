package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/devlopersabbir/terrorserver/internal/config"
	"github.com/devlopersabbir/terrorserver/internal/logger"
	"github.com/devlopersabbir/terrorserver/internal/proxy"
	"github.com/devlopersabbir/terrorserver/templates"
)

// routingTable is the live lookup map, swapped atomically on reload.
type routingTable struct {
	routes map[string]config.Route
}

// Server is the core terrorserver HTTP engine.
type Server struct {
	configPath string
	listenAddr string
	tablePtr   unsafe.Pointer // *routingTable
	proxyPool  *proxy.Pool
	httpServer *http.Server
	listener   net.Listener
	mu         sync.Mutex

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
	tbl := &routingTable{routes: cfg.RouteMap()}
	atomic.StorePointer(&s.tablePtr, unsafe.Pointer(tbl))
	atomic.AddInt64(&s.configLoads, 1)

	// Flush proxy cache so old upstreams aren't reused after a target change
	s.proxyPool.Flush()

	logger.Info("loaded %d route(s)", len(tbl.routes))
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

// table safely dereferences the atomic pointer.
func (s *Server) table() *routingTable {
	p := atomic.LoadPointer(&s.tablePtr)
	if p == nil {
		return &routingTable{routes: map[string]config.Route{}}
	}
	return (*routingTable)(p)
}

// lookup returns the Route matching the request Host header.
// It tries exact host match first, then port-only match (:port).
func (s *Server) lookup(r *http.Request) (config.Route, bool) {
	tbl := s.table()

	host := strings.ToLower(r.Host)
	// Strip port from host header for named-host matching
	hostOnly := host
	if h, _, err := net.SplitHostPort(host); err == nil {
		hostOnly = h
	}

	// 1. Exact full match (e.g. "api.example.com" or "api.example.com:80")
	if route, ok := tbl.routes[host]; ok {
		return route, true
	}
	// 2. Host-only match (strip port from incoming host)
	if route, ok := tbl.routes[hostOnly]; ok {
		return route, true
	}
	// 3. Port-only match (e.g. config key ":4000" matches any host on :4000)
	if _, port, err := net.SplitHostPort(r.Host); err == nil {
		key := ":" + port
		if route, ok := tbl.routes[key]; ok {
			return route, true
		}
	}
	return config.Route{}, false
}

// Start binds the server and begins serving.
func (s *Server) Start(addr string) error {
	s.startedAt = time.Now()
	s.listenAddr = addr

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("cannot bind %s: %w", addr, err)
	}
	s.listener = ln

	mux := http.NewServeMux()
	mux.Handle("/", s)

	s.httpServer = &http.Server{
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	logger.Info("listening on %s", addr)
	go s.httpServer.Serve(ln)
	return nil
}

// Shutdown gracefully stops the HTTP server.
func (s *Server) Shutdown(ctx context.Context) error {
	if s.httpServer == nil {
		return nil
	}
	return s.httpServer.Shutdown(ctx)
}

// ServeHTTP is the main request dispatcher.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	rw := &responseWriter{ResponseWriter: w, code: http.StatusOK}

	route, ok := s.lookup(r)
	if !ok {
		if s.shouldServeWelcome(r) {
			s.handleWelcome(rw)
			logger.Request(r.Method, r.Host, r.URL.Path, rw.code, time.Since(start))
			return
		}
		http.Error(rw, "Not Found", http.StatusNotFound)
		logger.Request(r.Method, r.Host, r.URL.Path, rw.code, time.Since(start))
		return
	}

	switch route.Type {
	case config.RouteProxy:
		s.handleProxy(rw, r, route)
	case config.RouteStatic:
		s.handleStatic(rw, r, route)
	default:
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
	}

	logger.Request(r.Method, r.Host, r.URL.Path, rw.code, time.Since(start))
}

func (s *Server) shouldServeWelcome(r *http.Request) bool {
	if portFromAddr(s.listenAddr) == "80" {
		return true
	}
	if s.listener != nil && portFromAddr(s.listener.Addr().String()) == "80" {
		return true
	}
	if _, port, err := net.SplitHostPort(r.Host); err == nil {
		return port == "80"
	}
	return false
}

func portFromAddr(addr string) string {
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

func (s *Server) handleWelcome(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(templates.WelcomePageHTML))
}

func (s *Server) handleProxy(w http.ResponseWriter, r *http.Request, route config.Route) {
	rp, err := s.proxyPool.Get(route.Target)
	if err != nil {
		logger.Error("proxy pool error for %s: %v", route.Target, err)
		http.Error(w, "Bad Gateway", http.StatusBadGateway)
		return
	}
	rp.ServeHTTP(w, r)
}

func (s *Server) handleStatic(w http.ResponseWriter, r *http.Request, route config.Route) {
	// SPA fallback: serve index.html for paths without an extension
	fs := http.Dir(route.Root)
	fileServer := http.FileServer(fs)

	// Try opening the exact path; fall back to index.html for SPA routing
	if _, err := fs.Open(r.URL.Path); err != nil {
		r2 := *r
		r2.URL.Path = "/"
		fileServer.ServeHTTP(w, &r2)
		return
	}
	fileServer.ServeHTTP(w, r)
}

// Status returns live runtime statistics.
type Status struct {
	Running     bool
	Uptime      time.Duration
	RouteCount  int
	ConfigLoads int64
	ConfigPath  string
}

func (s *Server) Status() Status {
	tbl := s.table()
	return Status{
		Running:     s.httpServer != nil,
		Uptime:      time.Since(s.startedAt),
		RouteCount:  len(tbl.routes),
		ConfigLoads: atomic.LoadInt64(&s.configLoads),
		ConfigPath:  s.configPath,
	}
}

// responseWriter wraps http.ResponseWriter to capture the status code.
type responseWriter struct {
	http.ResponseWriter
	code    int
	written bool
}

func (rw *responseWriter) WriteHeader(code int) {
	if !rw.written {
		rw.code = code
		rw.written = true
		rw.ResponseWriter.WriteHeader(code)
	}
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	if !rw.written {
		rw.written = true
	}
	return rw.ResponseWriter.Write(b)
}
