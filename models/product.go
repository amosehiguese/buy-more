package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Category struct {
	ID				uuid.UUID			`db:"id" validate:"required,uuid"`
	Name 			string 				`db:"name" validate:"required,lte=200"`
	Slug			string				`db:"slug" validate:"lte=200"`
	CreatedAt		time.Time			`db:"created_at"`
	UpdatedAt		time.Time			`db:"updated_at"`
}

type Product struct{
	ID  			uuid.UUID			`db:"id" validate:"required,uuid"`
	CategoryID		uuid.UUID			`db:"category_id" validate:"required"`
	Name			string 				`db:"name"  validate:"required,lte200"`
	Slug			string				`db:"slug" validate:"lte=200"`
	ImagePath		string				`db:"path"`
	Description		string				`db:"description"`
	Price			decimal.Decimal		`db:"price" validate:"gte=0"`
	Available		bool				`db:"available" validate:"required"`
	CreatedAt		time.Time			`db:"created_at"`
	UpdatedAt		time.Time			`db:"updated_at"`
}

type Image struct {
	ID				uuid.UUID		`db:"id" validate:"required,uuid"`
	ProductID		uuid.UUID 		`db:"product_id" validate:"required,uuid"`
	Path			string			`db:"path" validate:"required"`
	CreatedAt		time.Time		`db:"created_at"`
	UpdatedAt		time.Time		`db:"updated_at"`
}
