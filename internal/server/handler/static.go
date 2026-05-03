package handler

import (
	"net/http"

	"github.com/devlopersabbir/terrorserver/internal/config"
)

func Static(w http.ResponseWriter, r *http.Request, route config.Route) {
	fs := http.Dir(route.Root)
	fileServer := http.FileServer(fs)

	if _, err := fs.Open(r.URL.Path); err != nil {
		r2 := *r
		r2.URL.Path = "/"
		fileServer.ServeHTTP(w, &r2)
		return
	}
	fileServer.ServeHTTP(w, r)
}
