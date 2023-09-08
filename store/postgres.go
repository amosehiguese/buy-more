package store

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)



func postgresConn() (*sqlx.DB, error) {
	postgresURI := os.Getenv("POSTGRES_URL")
	dbMaxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	dbMaxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	dbMaxLifeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

	dbX, errX := sqlx.ConnectContext(context.Background(), "postgres", postgresURI)
	if err != nil {
		return nil, errX
	}

	dbX.SetConnMaxLifetime(time.Duration(dbMaxLifeConn))
	dbX.SetMaxIdleConns(dbMaxIdleConn)
	dbX.SetMaxOpenConns(dbMaxConn)

	if err := dbX.Ping(); err != nil {
		defer dbX.Close()
		return nil, err
	}

	log.Println("successfully connected to postgres database")

	return dbX, nil
}

func postgresMigration() error {
	postgresURI := os.Getenv("POSTGRES_URL")
	migSourceURL := "file://db/migrations"

	m, err := migrate.New(migSourceURL, postgresURI)
	if err != nil {
		log.Println("failed to generate migration instance ->", err)
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Println("Migration up failed ->", err)
		return err
	} else if (err == migrate.ErrNoChange) {
		log.Println("DB is up-to-date. No migrations done ->",err)
	} else{
		log.Println("Migration successful")
	}

	return nil
}
