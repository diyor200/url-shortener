package rest

func (h *Handler) RegisterRoutes() {
	h.Mux.HandleFunc("/", h.home)
	h.Mux.HandleFunc("/shorten", h.shorten)
}
