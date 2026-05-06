package handler

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/devlopersabbir/terrorserver/internal/config"
)

func Static(w http.ResponseWriter, r *http.Request, route config.Route) {
	fs := http.Dir(route.Root)
	fileServer := http.FileServer(fs)

	// Clean path to avoid traversal and ensure it starts with /
	path := r.URL.Path
	if path == "" {
		path = "/"
	}

	f, err := fs.Open(path)
	if err != nil {
		// Handle fallback
		fallbackPath := route.Fallback
		if fallbackPath == "" {
			fallbackPath = "/" // Default behavior
		}

		// Special case for index.html to avoid FileServer's 301 redirect to /
		if fallbackPath == "/index.html" {
			fallbackPath = "/"
		}

		// Try to see if fallback exists
		if _, err := fs.Open(fallbackPath); err == nil {
			r2 := *r
			r2.URL.Path = fallbackPath
			fileServer.ServeHTTP(w, &r2)
			return
		}

		// If fallback also doesn't exist, let FileServer handle 404
		fileServer.ServeHTTP(w, r)
		return
	}
	f.Close()

	// Issue 1: Force download for executable/installer files
	ext := strings.ToLower(filepath.Ext(path))
	downloadExts := map[string]bool{
		".sh":   true,
		".bash": true,
		".exe":  true,
		".msi":  true,
		".deb":  true,
		".rpm":  true,
		".dmg":  true,
		".bin":  true,
		".iso":  true,
		".zip":  true,
	}

	if downloadExts[ext] {
		filename := filepath.Base(path)
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	}

	fileServer.ServeHTTP(w, r)
}
