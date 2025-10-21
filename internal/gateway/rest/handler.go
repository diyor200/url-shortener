package rest

import (
	"html/template"
	"net/http"
)

type Handler struct {
	Mux  *http.ServeMux
	tmpl *template.Template
	//shortenUC shortenUC
}

func NewHandler() *Handler {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	tmpl := template.Must(template.ParseGlob("./template/index.html"))

	return &Handler{Mux: mux, tmpl: tmpl}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Mux.ServeHTTP(w, r)
}
