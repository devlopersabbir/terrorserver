package server

import (
	"sync/atomic"
	"time"
)

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
		Running:     len(s.httpServers) > 0,
		Uptime:      time.Since(s.startedAt),
		RouteCount:  tbl.Len(),
		ConfigLoads: atomic.LoadInt64(&s.configLoads),
		ConfigPath:  s.configPath,
	}
}
