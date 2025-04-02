package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnectDB() {
	DB_URI := os.Getenv("DB_URI")

	conn, err := pgx.Connect(context.Background(), DB_URI)
	if err != nil {
		log.Fatalf("ocorreu um erro ao carregar o .env: %v", err)
	}
	defer conn.Close(context.Background())
}
