package routes

import (
	"github.com/amosehiguese/buy-more/handlers"
	"github.com/gorilla/mux"
)

func UserRoutes(mux *mux.Router) {
	mux.HandleFunc("/login", handlers.Login)
}
