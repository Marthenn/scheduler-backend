package database

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// GetDBConnection returns a connection to the database on cockroachlab
func getDBConnection() *pgx.Conn {
	// load the env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := os.Getenv("DATABASE_URL")

	con, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	return con
}
