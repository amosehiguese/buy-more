package store

import (
	"log"

	"github.com/amosehiguese/buy-more/models"
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
	err error
)

func SetUpDB() {
	db, err = postgresConn()
	if err != nil {
		log.Fatalln("unable to connect to db ->", err)
	}
	postgresMigration()
}

type queries struct {
	*models.ProductQueries
}

func GetQuery() *queries {
	return &queries{
		ProductQueries: &models.ProductQueries{DB: db},
	}
}
