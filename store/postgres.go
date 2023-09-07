package store

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func postgresConn() (*sqlx.DB, error) {
	postgresURI := os.Getenv("POSTGRES_URL")
	dbMaxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	dbMaxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	dbMaxLifeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

	db, err := sqlx.ConnectContext(context.Background(), "postgres", postgresURI)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Duration(dbMaxLifeConn))
	db.SetMaxIdleConns(dbMaxIdleConn)
	db.SetMaxOpenConns(dbMaxConn)

	migSourceURL := "file://migrations"
	err = postgresMigration(migSourceURL, postgresURI)
	if err != nil {
		log.Println("migration failed ->", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		defer db.Close()
		return nil, err
	}

	log.Println("successfully connected to postgres database")

	return db, nil
}

func postgresMigration(sourceURL, dbURL string) error {
	m, err := migrate.New(sourceURL, dbURL)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil {
		return err
	}

	return nil
}
