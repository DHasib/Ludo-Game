package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

// Connect initializes the global DB pool.
func Connect(url string) {
	var err error
	DB, err = pgxpool.New(context.Background(), url)
	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
}
