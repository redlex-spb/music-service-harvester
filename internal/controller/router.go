package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handlers struct {
	Sources http.HandlerFunc
	Songs   http.HandlerFunc
	Image   http.HandlerFunc
}

func NewRouter(h Handlers) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/v1/harvest/sources", h.Sources)
	r.Post("/v1/harvest/songs", h.Songs)
	r.Post("/v1/harvest/image", h.Image)
	return r
}
