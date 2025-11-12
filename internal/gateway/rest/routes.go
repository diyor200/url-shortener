package rest

import (
	"net/http/pprof"
)

func (h *Handler) RegisterRoutes() {
	h.Mux.HandleFunc("/", h.home)
	h.Mux.HandleFunc("/shorten", h.shorten)
	h.Mux.HandleFunc("/r/", h.redirect)
	// debug
	h.Mux.HandleFunc("/debug/pprof", pprof.Index)
	h.Mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	h.Mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	h.Mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	h.Mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
}
