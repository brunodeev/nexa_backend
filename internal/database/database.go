package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnectDB() *pgx.Conn {
	DB_URI := os.Getenv("DB_URI")

	conn, err := pgx.Connect(context.Background(), DB_URI)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	fmt.Println("Database connection successful.")

	return conn
}
