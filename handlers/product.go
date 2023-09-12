package handlers

import (
	"log"
	"net/http"

	"github.com/amosehiguese/buy-more/store"
	"github.com/amosehiguese/buy-more/util/httperror"
	"github.com/amosehiguese/buy-more/util/tmpl"
)

func Home(w http.ResponseWriter, r *http.Request) {
	q := store.GetQuery()
	products, err := q.ProductQueries.GetProducts()
	if err != nil{
		log.Println(err)
		httperror.Message(w, r, "No products available")
	}
	err = tmpl.GenerateHTML(w, products, "layout","public.navbar","public.sidebar","home")
	if err != nil {
		log.Println(err)
	}

}

func AllProducts(w http.ResponseWriter, r *http.Request) {}

func RetrieveProduct(w http.ResponseWriter, r *http.Request) {}

