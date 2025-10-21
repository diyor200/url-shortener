package rest

func (h *Handler) RegisterRoutes() {
	h.Mux.HandleFunc("/", h.home)
}
