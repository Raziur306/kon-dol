package router

import (
	"net/http"

	"github.com/Raziur306/kon-dol/internal/handler"
	"github.com/go-chi/chi/v5"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status": "ok", "message": "Server is running"}`))

	})

	r.Get("/incidents", handler.GetIncidents)

	return r
}
