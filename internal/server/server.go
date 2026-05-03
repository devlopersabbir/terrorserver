package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/devlopersabbir/terrorserver/internal/config"
	"github.com/devlopersabbir/terrorserver/internal/logger"
	"github.com/devlopersabbir/terrorserver/internal/proxy"
	"github.com/devlopersabbir/terrorserver/internal/server/handler"
	"github.com/devlopersabbir/terrorserver/internal/server/response"
	"github.com/devlopersabbir/terrorserver/internal/server/router"
)

// Server is the core terrorserver HTTP engine.
type Server struct {
	configPath string
	listenAddr string
	tablePtr   atomic.Pointer[router.Table]
	proxyPool  *proxy.Pool
	httpServer *http.Server
	listener   net.Listener

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
	rw := response.NewRecorder(w)

	route, ok := router.Lookup(s.table(), router.RequestContext{
		Host:         r.Host,
		ListenAddr:   s.listenAddr,
		ListenerAddr: s.listenerAddr(),
	})
	if !ok {
		if router.IsPort80(s.listenAddr, s.listenerAddr(), r.Host) {
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

func (s *Server) listenerAddr() string {
	if s.listener == nil {
		return ""
	}
	return s.listener.Addr().String()
}
