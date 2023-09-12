package models

import (
	"log"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ProductQueries struct {
	*sqlx.DB
}

func (q *ProductQueries) CreatedAtDate() string {
	p := Product{}
	return p.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

func (q *ProductQueries) GetProducts() (*[]Product, error){
	products := []Product{}

	query := `SELECT * FROM products`
	if err := q.Select(&products, query); err != nil {
		return nil, err
	}

	log.Println("from q ->", products)

	return &products, nil
}

func (q *ProductQueries) RetrieveProduct(id uuid.UUID, slug string) (*Product, error) {
	var product Product
	query := `SELECT * FROM products WHERE id=$1 AND slug=$2`
	if err := q.Get(&product, query, id, slug); err != nil {
		return nil, err
	}

	return &product, nil
}

func (q *ProductQueries) GetProductsByCategory(cat_slug string) (*[]Product, error) {
	var products []Product
	query := `SELECT * FROM products WHERE category_id=$1'`
	if err := q.Select(&products, query, cat_slug); err != nil {
		return nil, err
	}

	return &products, nil
}


