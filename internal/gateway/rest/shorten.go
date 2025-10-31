package rest

import (
	"encoding/json"
	"github.com/diyor200/url-shortener/internal/gateway/rest/scheme"
	"io"
	"net/http"
)

func (h *Handler) shorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method now allowed", http.StatusMethodNotAllowed)
		return
	}

	var req scheme.URLRequest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	res, err := h.shortenUC.Shorten(r.Context(), req.URL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	url := scheme.URL{
		ID:        res.ID,
		ShortURL:  "http://" + r.Host + "/r/" + res.Short,
		LongURL:   res.Long,
		CreatedAt: scheme.ConvertTimeToString(res.CreatedAt),
	}

	resp, err := json.Marshal(url)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	w.Write(resp)
	return
}
