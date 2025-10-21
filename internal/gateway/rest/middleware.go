package rest

import (
	"fmt"
	"net/http"
	"time"
)

func (h *Handler) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
		fmt.Printf("Completed %s %s in %v from %s\n", r.Method, r.URL.Path, time.Since(start), r.RemoteAddr)
	})
}
