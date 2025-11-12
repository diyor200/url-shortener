package rest

import (
	"github.com/rs/zerolog/log"
	"net/http"
)

func (h *Handler) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info().
			Str("IP", r.RemoteAddr).
			Str("Method", r.Method).
			Str("URL", r.URL.String()).
			Msg("Request received")

		next.ServeHTTP(w, r)
	})
}
