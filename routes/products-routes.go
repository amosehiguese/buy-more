package routes

import (
	"github.com/amosehiguese/buy-more/handlers"
	"github.com/gorilla/mux"
)

func ProductRoutes(mux *mux.Router) {
	mux.HandleFunc("/", handlers.Index)
}
