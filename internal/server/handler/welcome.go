package handler

import (
	"net/http"

	"github.com/devlopersabbir/terrorserver/templates"
)

func Welcome(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(templates.WelcomePageHTML))
}
