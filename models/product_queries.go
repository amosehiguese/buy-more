package models

import "github.com/jmoiron/sqlx"

type ProductQueries struct {
	*sqlx.DB
}
