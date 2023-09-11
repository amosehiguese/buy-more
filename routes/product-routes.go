package routes

import (
	"github.com/amosehiguese/buy-more/handlers"
	"github.com/gorilla/mux"
)

func ProductRoutes(mux *mux.Router) {
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/products", handlers.AllProducts)
	mux.HandleFunc("/product/:id/:slug", handlers.RetrieveProduct)
}
