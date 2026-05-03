package handler

import (
	"net/http"

	"github.com/devlopersabbir/terrorserver/internal/config"
	"github.com/devlopersabbir/terrorserver/internal/logger"
	"github.com/devlopersabbir/terrorserver/internal/proxy"
)

func Proxy(w http.ResponseWriter, r *http.Request, pool *proxy.Pool, route config.Route) {
	rp, err := pool.Get(route.Target)
	if err != nil {
		logger.Error("proxy pool error for %s: %v", route.Target, err)
		http.Error(w, "Bad Gateway", http.StatusBadGateway)
		return
	}
	rp.ServeHTTP(w, r)
}
