package handlers

import (
	"net/http"

	"github.com/amosehiguese/buy-more/store"
	"github.com/amosehiguese/buy-more/util/tmpl"
)

func Home(w http.ResponseWriter, r *http.Request) {
	q := store.GetQuery()
	products, _ := q.ProductQueries.GetProducts()
	tmpl.GenerateHTML(w, products, "layout","home")
}

func AllProducts(w http.ResponseWriter, r *http.Request) {}

func RetrieveProduct(w http.ResponseWriter, r *http.Request) {}

