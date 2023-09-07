package routes

import (
	"github.com/amosehiguese/buy-more/handlers"
	"github.com/go-chi/chi/v5"
)

func ProductRoutes(mux *chi.Mux) {
	mux.HandleFunc("/", handlers.Index)
}
