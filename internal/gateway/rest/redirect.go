package rest

import (
	"errors"
	"github.com/diyor200/url-shortener/internal/errs"
	"net/http"
	"strings"
)

func (h *Handler) redirect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	key := strings.TrimPrefix(r.URL.Path, "/r/")
	if key == "" {
		http.NotFound(w, r)
		return
	}

	// add check from caching later
	// get from db
	data, err := h.shortenUC.Get(r.Context(), key)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			http.NotFound(w, r)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, data.Long, http.StatusMovedPermanently)
}
