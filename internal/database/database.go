package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnectDB() {
	DB_URI := os.Getenv("DB_URI")

	conn, err := pgx.Connect(context.Background(), DB_URI)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	fmt.Println("Database connection successful.")
}
