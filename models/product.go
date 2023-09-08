package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Category struct {
	Uuid 			uuid.UUID			`db:"uuid" validate:"required,uuid"`
	Name 			string 				`db:"name" validate:"required,lte=200"`
	Slug			string				`db:"slug" validate:"required,lte=200"`
	Products		[]Product			`db:"products"`
	CreatedAt		time.Time			`db:"created_at"`
	UpdatedAt		time.Time			`db:"updated_at"`
}

type Product struct{
	Uuid  			uuid.UUID			`db:"uuid" validate:"required,uuid"`
	Category		Category			`db:"category" validate:"required"`
	Name			string 				`db:"name"  validate:"required,lte200"`
	Slug			string				`db:"slug" validate:"required,lte=200"`
	Images			[]Image				`db:"images"`
	Description		string				`db:"description"`
	Price			decimal.Decimal		`db:"price" validate:"gte=0"`
	Available		bool				`db:"available"`
	CreatedAt		time.Time			`db:"created_at"`
	UpdatedAt		time.Time			`db:"updated_at"`
}

type Image struct {
	Uuid			uuid.UUID		`db:"uuid" validate:"required,uuid"`
	ProductID		uuid.UUID 		`db:"product_id" validate:"require,uuid"`
	Path			string			`db:"path"`
	CreatedAt		time.Time		`db:"created_at"`
	UpdatedAt		time.Time		`db:"updated_at"`

}
